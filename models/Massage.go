package models

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/fatih/set"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormID      string `json:"from_id"`   //发送者
	TargetID    string `json:"target_id"` //接收者
	ChatType    int    `json:"chat_type"` //消息类型 群聊 、 私聊 、 广播
	Media       int    `json:"media"`     //消息类型 文字、图片、音频
	Context     string `json:"context"`   //消息内容
	Pic         string `json:"pic"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var (
	// 映射
	clientMap map[string]*Node = make(map[string]*Node, 0)

	// 读写锁
	rwLocker sync.RWMutex

	udpsendChannel chan []byte = make(chan []byte, 1024)
)

func (table *Message) TableNanme() string {
	return "message"
}

func Chat(writer http.ResponseWriter, request *http.Request) {
	//获取URL参数
	query := request.URL.Query()
	account := query.Get("account")
	log.Println(" sendMsg >>> account: ", account)
	//accountID, _ := strconv.ParseInt(account, 10, 64)
	// msgType := query.Get("type")
	// targetId := query.Get("targetId")
	isvalida := true
	// context := query.Get("context")

	// 使用 websocket 升级器将 HTTP 连接升级为 WebSocket 连接
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		return // 如果升级失败，则退出函数
	}

	// 创建一个节点结构体，用于保存 WebSocket 连接、数据队列和组信息
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	// 加锁，向客户端映射表中添加客户端信息
	rwLocker.Lock()
	clientMap[account] = node
	rwLocker.Unlock()

	// 启动发送处理 goroutine
	go sendProc(node)

	// 启动接收处理 goroutine
	go recvProc(node)

	sendMsg(account, []byte("hello"))

}

// sendProc 函数用于处理向客户端发送消息的逻辑
func sendProc(node *Node) {
	for {
		select {
		// 从节点数据队列中接收待发送的数据
		case data := <-node.DataQueue:
			log.Println("[ws] this is sendProc >>> msg: ", data)
			// 使用 WebSocket 连接向客户端发送文本消息
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err)
				return
			}
		}

	}
}

// recvProc 函数用于处理从客户端接收消息的逻辑
func recvProc(node *Node) {
	for {
		// 从 WebSocket 连接中读取消息数据
		_, data, err := node.Conn.ReadMessage()
		log.Println("this is recvProc")
		if err != nil {
			log.Println(err)
			return
		}
		broadMsg(data) // 调用广播消息函数，将接收到的消息广播给所有客户端
		log.Println("[ws]  <<<<<  ", string(data))

	}
}

func broadMsg(data []byte) {
	udpsendChannel <- data
}

// init 函数用于在程序启动时初始化 UDP 发送和接收处理的 goroutine
func init() {
	go udpSendProc() // 启动 UDP 发送处理的 goroutine
	go udpRecvProc() // 启动 UDP 接收处理的 goroutine

	log.Println("init goroutine")
}

// udpSendProc 函数用于处理向指定 UDP 地址发送数据的逻辑
func udpSendProc() {
	// 连接到指定的 UDP 地址
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 171, 243), // 设置目标 IP 地址
		Port: 3000,                         // 设置目标端口
	})

	defer conn.Close() // 延迟关闭 UDP 连接
	if err != nil {
		log.Println(err) // 若连接出现错误，则打印错误信息
		return           // 退出发送处理函数
	}

	for {
		select {
		// 从 udpsendChannel 通道中接收待发送的数据
		case data := <-udpsendChannel:
			log.Println("this is udpSendProc data:", string(data)) // 打印日志，表示执行 UDP 发送处理函数
			// 向 UDP 地址发送数据
			_, err = conn.Write(data)
			if err != nil {
				log.Println(err) // 若发送数据出错，则打印错误信息
				return           // 退出发送处理函数
			}
		}

	}
}

// udpRecvProc 函数用于处理从指定 UDP 端口接收数据并进行分发的逻辑
func udpRecvProc() {
	// 监听指定的 UDP 地址
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero, // 监听所有网络接口上的 UDP 数据包
		Port: 3000,         // 监听的端口号
	})

	defer conn.Close() // 延迟关闭 UDP 连接
	if err != nil {
		log.Println(err) // 若出现错误，则打印错误信息
		return           // 退出接收处理函数
	}

	for {
		var buf [512]byte            // 创建一个大小为 512 字节的缓冲区
		n, err := conn.Read(buf[0:]) // 读取 UDP 数据包到缓冲区中
		if err != nil {
			log.Println(err) // 若读取数据出错，则打印错误信息
			return           // 退出接收处理函数
		}

		dispatch(buf[0:n]) // 调用 dispatch 函数对接收到的数据进行分发处理
	}
}

// dispatch 函数用于根据接收到的消息类型进行分发处理
/**
	models 绑定json 解决反序列化失败 ---bug 
*/
func dispatch(data []byte) {
	msg := Message{}                  // 创建一个空的 Message 结构体
	err := json.Unmarshal(data, &msg) // 解析接收到的 JSON 数据到 Message 结构体中
	if err != nil {
		log.Println(err) // 若解析数据出错，则打印错误信息
		return           // 退出分发处理函数
	}
	switch msg.ChatType {
	case 1: // 如果消息类型为 1

		sendMsg(msg.TargetID, data) // 调用 sendMsg 函数发送消息给指定账户
	}
}

// sendMsg 函数用于向指定账户发送消息
func sendMsg(account string, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[account] // 从 clientMap 中获取指定账户的节点信息
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg // 将消息写入该账户节点的数据队列中
	}
}
