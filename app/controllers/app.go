package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Social() revel.Result {
	return c.Render()
}

func (c App) Posts() revel.Result {
	return c.Render()
}

func (c App) Post(id uint) revel.Result {
	// Look Up Id and Render Specific Posts
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Resume() revel.Result {
	return c.Render()
}

func (c App) Links() revel.Result {
	return c.Render()
}