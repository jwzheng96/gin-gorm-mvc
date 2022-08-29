# gin-gorm-mvc
 Restful API in Golang following best practices, built with gin, gorm(sqlite), and MVC architecture.

 ## 目录说明

     .
    ├── Config      配置目录
    ├── Controllers 控制器
    ├── Models      model目录
    ├── Route       router
    └── main.go     主文件


```
go mod init gin-gorm-mvc

go get github.com/go-sql-driver/mysql

go get github.com/gin-gonic/gin

go get github.com/jinzhu/gorm

go run main.go
```