package models

import (
	"log"
	"time"
	"strconv"
	"github.com/jyu/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Account       string `gorm:"unique"`
	Password      string
	Salt          string
	Email         string `valid:"email"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableNanme() string {
	return "user_basic"
}

func GetUserList() []UserBasic {
	data := make([]UserBasic, 10)
	utils.DB_MySQL.Find(&data)

	for _, v := range data {
		log.Println(v)
	}
	return data
}

func FindUserByAccount(account string) (UserBasic, error) {
	var data UserBasic
	err := utils.DB_MySQL.Where("account = ?", account).First(&data).Error
	return data, err
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Where("account = ?", user.Account).Delete(&UserBasic{})
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Where("account = ?", user.Account).Model(user).Updates(UserBasic{Email: user.Email, Phone: user.Phone})
}

func UpdatePassword(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Where("account = ?", user.Account).Model(user).Updates(UserBasic{Password: user.Password})
}

func UpdateEmail(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Where("account = ?", user.Account).Model(user).Updates(UserBasic{Email: user.Email})
}

func UpdatePhone(user UserBasic) *gorm.DB {
	return utils.DB_MySQL.Where("account = ?", user.Account).Model(user).Updates(UserBasic{Phone: user.Phone})
}

func UpdateIdentity(user UserBasic) *gorm.DB {
	tmp := utils.MD5EnCode(strconv.FormatInt(time.Now().Unix(), 10))
	return utils.DB_MySQL.Where("account = ?", user.Account).Model(&user).Update("identity", tmp)
}