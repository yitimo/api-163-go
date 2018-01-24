package homura

import (
	"encoding/json"
	"strconv"

	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// PlayListGroupInit 初始化歌单路由组
func PlayListGroupInit(m *martini.ClassicMartini) {
	m.Group("/playlist", func(router martini.Router) {
		router.Get("/:cat/:order/:page/:limit", func(p martini.Params, r render.Render) {
			page, perr := strconv.ParseInt(p["page"], 10, 64)
			if perr != nil {
				page = 1
			}
			limit, lerr := strconv.ParseInt(p["limit"], 10, 64)
			if lerr != nil {
				limit = 10
			}
			// 发送POST请求得到最后包含url的结果
			reqRs, reqErr := madoka.PlayList(p["cat"], p["order"], (int)(page), (int)(limit))
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			// 应该可以解析到第一层json
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse})
			}
		})
		router.Get("/catalogue", func(p martini.Params, r render.Render) {
			// 发送POST请求得到最后包含url的结果
			reqRs, reqErr := madoka.PlayListCatalogue()
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			// 应该可以解析到第一层json
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse})
			}
		})
		router.Get("/detail/:id", func(p martini.Params, r render.Render) {
			// 发送POST请求得到最后包含url的结果
			reqRs, reqErr := madoka.PlayListDetail(p["id"])
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			// 应该可以解析到第一层json
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse})
			}
		})
	})
}
