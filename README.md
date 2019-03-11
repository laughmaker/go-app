### 启动步骤

1. 相关的依赖库
```
go get github.com/gomodule/redigo/redis
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/jinzhu/gorm
go get github.com/go-ini/ini
go get gopkg.in/gomail.v2
go get github.com/gin-gonic/gin
go get github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/gin-swagger/swaggerFiles
go get github.com/mongodb/mongo-go-driver@v1.0.0-rc1
go get github.com/laughmaker/go-pkg
```

2. 生成文档，生成文档需要关闭mod，目前库不支持mod模式
```
swag init
```

3. 启动
* 热加载启动
    * 安装更新库：`go get github.com/oxequa/realize`
    * 启动项目: `realize start --run server`
* 普通启动项目: `go run main.go`

___

### 注意事项
* 编译时采用,避免自动修改mod.go
```
go build -mod=readonly
```

* 开启mod，在根目录下运行
```
export GO111MODULE=on 
```

* 设置代理，需要开启mod模式，设置代理后，代理库中的代码不会及时更新，有缓存，如某个库没有最新版，可关闭再下载
    * 开启：`export GOPROXY=https://goproxy.io`
    * 关闭：`export GOPROXY=`


### 效率工具

* 根据接口自动生成struct
  * 安装：https://github.com/galeone/rts
  * 示例：
  ```rts -server http://localhost:8000/topic/topics/sdf topic```

* 根据表结构自动生成结构
  * 安装：github.com/gohouse/converter
  * 示例：
  ```./table2struct.bin -file ../model/cat_diary.go -dsn root:123456@tcp(localhost:3306)/yirimao?charset=utf8mb4 -table cat_diary -tagKey gorm```
