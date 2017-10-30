package sayaka

import (
	"github.com/go-martini/martini"
	"../madoka"
	"github.com/martini-contrib/render"
	"encoding/json"
)

/**
 * 初始化搜索路由组
 */
 func downloadGroupInit(m *martini.ClassicMartini) {
	m.Group("/download", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "320000", r)
		})
		router.Get("/low/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "160000", r)
		})
		router.Get("/middle/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "320000", r)
		})
		router.Get("/high/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "640000", r)
		})
	})
}
/**
 * 执行搜索并使用render返回json数据
 */
func getDownloadUrl(id string, rate string, r render.Render) {
	// 拿到字符串结果
	reqRs := madoka.Download(id, rate)
	// 应该可以解析到第一层json
	var originParse map[string] interface{}
	if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	} else {
		r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["data"]})
	}
}