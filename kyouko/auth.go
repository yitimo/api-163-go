package kyouko

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
)

// AuthInit 认证守卫
func AuthInit(m *martini.ClassicMartini) {
	m.Use(auth.BasicFunc(func(username, password string) bool {
		return auth.SecureCompare(username, "yitimo") && auth.SecureCompare(password, "iamyitimo")
	}))
}
