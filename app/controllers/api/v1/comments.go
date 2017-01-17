// app/controllers/api/v1/comments.go
package controllers

import (
	"github.com/fk2000/golang_bbs/app/controllers"
	"github.com/fk2000/golang_bbs/app/models"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
)

type ApiV1Comments struct {
	ApiV1Controller
}

func (c ApiV1Comments) Index() revel.Result {
	comments := []models.Comment{}

	if err := controllers.DB.Order("id desc").Find(&comments).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	r := Response{comments}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Show(id int) revel.Result {
	comment := &models.Comment{}
	if err := controllers.DB.First(&comment, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{comment}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Create() revel.Result {
	comment := &models.Comment{}

	if err := c.BindParams(comment); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := validator.Validate(comment); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := controllers.DB.Create(comment).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	r := Response{comment}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Delete(id int) revel.Result {
	comment := models.Comment{}

	if err := controllers.DB.First(&comment, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := controllers.DB.Delete(&comment).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	r := Response{"success"}
	return c.RenderJson(r)
}
