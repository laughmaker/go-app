package api

import (
	"app/src/model"
	"app/src/pkg/app"

	"github.com/gin-gonic/gin"
)

// @Summary Get multiple topics
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /shop/v1/topic/topics/{id} [get]
func Topics(c *gin.Context) {
	page := app.NewPage(c)
	var topic model.Topic
	list, err := topic.Topics(&page)

	resp := app.Resp{C: c}
	if err != nil {
		resp.Failed(app.FAILED)
		return
	}

	resp.List(list, page)
}

func Save(c *gin.Context) {
	var topic model.Topic
	topic.Insert()
	resp := app.Resp{C: c}
	resp.Model(topic)
}
