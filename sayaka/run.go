package sayaka

import (
	"net/http"

	"../homura"
	"../kyouko"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// Run 启动martini
func Run(host string) {
	m := martini.Classic()
	m.Use(render.Renderer())
	kyouko.ConnectInit(m)
	kyouko.AuthInit(m)
	m.Get("/", func() string {
		return "Hello this is saber Sayaka !"
	})
	homura.SearchGroupInit(m)
	homura.DownloadGroupInit(m)
	homura.InfoGroupInit(m)
	homura.LyricGroupInit(m)
	m.Use(func(res http.ResponseWriter) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
	})
	m.NotFound(func() string {
		return "Sorry, Sayaka can not understand what you asked :("
	})
	m.RunOnAddr(host)
}
