package kyouko

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

// ConnectInit 连接守卫
func ConnectInit(m *martini.ClassicMartini) {
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3001", "https://www.yitimo.com", "https://yitimo.com"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
