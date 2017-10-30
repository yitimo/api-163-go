package sayaka

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run(host string) {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "Hello this is saber Sayaka !"
	})
	searchGroupInit(m)
	m.Use(func(res http.ResponseWriter) {
		// res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
	})
	m.NotFound(func() string {
		return "Sorry, Sayaka can not understand what you asked :("
	})
	m.RunOnAddr(host)
}
