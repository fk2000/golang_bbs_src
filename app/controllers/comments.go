package controllers

import (
	"github.com/fk2000/golang_bbs/app/models"
	"github.com/revel/revel"
	"net/http"
)

type Comments struct {
	*revel.Controller
}

func (c Comments) Index() revel.Result {
	comments := []models.Comment{}

	if err := DB.Order("id desc").Find(&comments).Error; err != nil {
		c.Response.Status = http.StatusNotFound
		// TODO: error rendering
	}

	c.RenderArgs["comments"] = comments
	return c.Render()
}
