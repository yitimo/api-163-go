package kyouko

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

// ConnectInit 连接守卫
func ConnectInit(m *martini.ClassicMartini, whitelist []interface{}) {
	list := make([]string, 0)
	for l := range whitelist {
		list = append(list, (string)(l))
	}
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     list,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
