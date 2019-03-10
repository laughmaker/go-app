package app

import (
	"net/http"

	"app/src/pkg/conf"
	"app/src/pkg/db"
	"app/src/pkg/log"
	"app/src/pkg/mongo"
	"app/src/pkg/redis"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	C *gin.Context
}

// 加载配置组件
func Config() {
	conf.Setup()
	log.Setup()
	db.Setup()
	redis.Setup()
	mongo.Setup()
}

func (resp *Resp) Send(httpStatus, code int, data interface{}, page interface{}) {
	message := GetMessage(code)
	if httpStatus != http.StatusOK {
		message = http.StatusText(httpStatus)
	}

	resp.C.JSON(httpStatus, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
		"page":    page,
	})
}

func (resp *Resp) List(list interface{}, page Page) {
	resp.Send(http.StatusOK, SUCCESS, list, page)
}

func (resp *Resp) Model(model interface{}) {
	resp.Send(http.StatusOK, SUCCESS, model, nil)
}

func (resp *Resp) Failed(code int) {
	resp.Send(http.StatusOK, code, nil, nil)
}

func (resp *Resp) Error(httpStatus int) {
	resp.Send(httpStatus, FAILED, nil, nil)
}
