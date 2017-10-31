package homura

import (
	"github.com/go-martini/martini"
	"../madoka"
	"encoding/json"
	"github.com/martini-contrib/render"
)


/**
 * 初始化搜索路由组
 */
func InfoGroupInit(m *martini.ClassicMartini) {
	m.Group("/info", func(router martini.Router) {
		router.Get("/music/:id", func(p martini.Params, r render.Render) {
			// 拿到字符串结果
			reqRs := madoka.SongInfo(p["id"])
			// 应该可以解析到第一层json
			var originParse map[string] interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["songs"]})
			}
		})
		router.Get("/artist/albums/:id", func(p martini.Params, r render.Render) {
			r.JSON(200, map[string]interface{}{"state": false, "msg": "正在完善中", "data": nil})
		})
		router.Get("/album/:id", func(p martini.Params, r render.Render) {
			r.JSON(200, map[string]interface{}{"state": false, "msg": "正在完善中", "data": nil})
		})
		router.Get("/playlist/:id", func(p martini.Params, r render.Render) {
			r.JSON(200, map[string]interface{}{"state": false, "msg": "正在完善中", "data": nil})
		})
		// router.Get("/:words/:page/:limit", func(p martini.Params, r render.Render) {
		// 	doSearch(p, r, "1")
		// })
		// router.Get("/album/:words/:page/:limit", func(p martini.Params, r render.Render) {
		// 	doSearch(p, r, "10")
		// })
		// router.Get("/artist/:words/:page/:limit", func(p martini.Params, r render.Render) {
		// 	doSearch(p, r, "100")
		// })
	})
}
