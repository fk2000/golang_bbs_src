// top.go
package controllers

import "github.com/revel/revel"

type Top struct {
	*revel.Controller
}

func (h Top) Index() revel.Result {
	return h.Render()
}
