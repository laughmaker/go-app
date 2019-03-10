package model

import (
	"app/src/pkg/app"
	"app/src/pkg/db"
)

type Topic struct {
	Model
	ID       int    `json:"topicId"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	ImageUrl string `json:"imageUrl"`
}

func (t *Topic) Topics(page *app.Page) (list []Topic, err error) {
	err = db.DB.Find(&list).Count(&page.TotalRows).Error
	page.Update()

	return
}

func (t *Topic) Insert() (id int, err error) {
	result := db.DB.Create(&t)
	id = t.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
