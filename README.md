# OrderSystem
A very simple practice Order system



go mod download

go mod tidy



##### beego项目


beego中文文档：

https://beego.gocn.vip/beego/zh/developing/#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B



##### 遇到的问题



1、beego框架出现的问题-----panic: ./ippanichandle.exe flag redefined: graceful

https://blog.csdn.net/qwerty1372431588/article/details/119244413

问题出在，项目本身同时使用了astaxie/beego 和 beego/beego



2、beego 2.0升级的巨坑，graceful错误

https://blog.csdn.net/hotqin888/article/details/122270738



Go 升级,通过该方式用户可以升级 beego 框架，强烈推荐该方式
go get -u github.com/beego/beego/v2
