package controllers

import (
	"aahframe.work"

	"aahcfgtest/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application's home page.
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Web Application",
		},
	}

	c.Reply().Ok().HTML(data)
}
