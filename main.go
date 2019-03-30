package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "yq_shop/models"
	_ "yq_shop/routers"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:feng629@tcp(118.25.35.217:3306)/yq_shop?charset=utf8")

	//注册静态文件(图片等)上传下载目录
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
