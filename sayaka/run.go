package sayaka

import (
	"net/http"

	"../homura"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

// Run 启动martini
func Run(host string) {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:3001", "https://www.yitimo.com", "https://yitimo.com"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	m.Use(auth.BasicFunc(func(username, password string) bool {
		return auth.SecureCompare(username, "yitimo") && auth.SecureCompare(password, "iamyitimo")
	}))
	m.Get("/", func() string {
		return "Hello this is saber Sayaka !"
	})
	homura.SearchGroupInit(m)
	homura.DownloadGroupInit(m)
	homura.InfoGroupInit(m)
	homura.LyricGroupInit(m)
	m.Use(func(res http.ResponseWriter) {
		// res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
	})
	m.NotFound(func() string {
		return "Sorry, Sayaka can not understand what you asked :("
	})
	m.RunOnAddr(host)
}
