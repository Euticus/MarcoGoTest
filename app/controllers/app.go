package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// I actually might be following what this page is doing here. Crazy it's just in a function like that though! 04012020
