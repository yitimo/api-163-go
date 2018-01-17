package homura

import (
	"encoding/json"
	"net/http"

	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// InfoGroupInit 初始化搜索路由组
func InfoGroupInit(m *martini.ClassicMartini) {
	m.Group("/song", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			reqRs, reqErr := madoka.SongInfo("[" + p["id"] + "]")
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["songs"]})
			}
		})
		router.Post("/", func(req *http.Request, r render.Render) {
			req.ParseForm()
			for data := range req.Form {
				var pParse map[string]interface{}
				if err := json.Unmarshal([]byte(data), &pParse); err == nil {
					reqRs, reqErr := madoka.SongInfo(formatIds(pParse["ids"].([]interface{})))
					if reqErr != nil {
						r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
						return
					}
					var originParse map[string]interface{}
					if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
						r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
					} else {
						r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["songs"]})
					}
				} else {
					r.JSON(200, map[string]interface{}{"state": false, "msg": "参数错误", "data": nil})
				}
			}
		})
	})
}
