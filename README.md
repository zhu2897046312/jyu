# jyu 教育信息管理平台
```

```
# 经验之谈      -- 因为这个丢了30块还没明白，白瞎30
```
1. go mod 导包路径 module-name/path  例如: jyu/services/mysql       
2. 包名随便写，不重复就行

3. gorm 奇怪错误  -->  在打开数据库时一定要加这些?charset=utf8&parseTime=True&loc=Local 否则就会遇到奇怪错误
```