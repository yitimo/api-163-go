package homura

import (
	"encoding/json"
	"strconv"
	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// ArtistGroupInit 初始化歌手路由组
func ArtistGroupInit(m *martini.ClassicMartini) {
	m.Group("/artist", func(router martini.Router) {
		router.Get("/top/:page/:limit", func(p martini.Params, r render.Render) {
			page, perr := strconv.ParseInt(p["page"], 0, 32)
			limit, lerr := strconv.ParseInt(p["limit"], 0, 32)
			if perr != nil || lerr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "参数错误", "data": nil})
			}
			// 发送POST请求得到最后包含url的结果
			reqRs, reqErr := madoka.ArtistTop((int)(page), (int)(limit))
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			// 应该可以解析到第一层json
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["artists"]})
			}
		})
	})
}
