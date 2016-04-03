package controllers

import(
	"github.com/revel/revel"
	"amiller.im/app/models"
	"github.com/davecgh/go-spew/spew"
"html/template"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//a := models.Post

	return c.Render()
}

func (c App) Social() revel.Result {
	return c.Render()
}

func (c App) Posts() revel.Result {
	posts := models.GetLastNPost(5)
	//spew.Dump(posts)
	return c.Render(posts)
}

func (c App) Post(id uint) revel.Result {
	// Look Up Id and Render Specific Posts
	post := models.GetPostById(id)
	//body := post.Body
	bodyTemplate := template.HTML(post.Body)
	spew.Dump(post)
	return c.Render(post, bodyTemplate)
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