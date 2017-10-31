package homura

import (
	"github.com/go-martini/martini"
	"../madoka"
	"strconv"
	"encoding/json"
	"github.com/martini-contrib/render"
)


/**
 * 初始化搜索路由组
 */
func SearchGroupInit(m *martini.ClassicMartini) {
	m.Group("/search", func(router martini.Router) {
		router.Get("/:words/:page/:limit", func(p martini.Params, r render.Render) {
			doSearch(p, r, "1")
		})
		router.Get("/album/:words/:page/:limit", func(p martini.Params, r render.Render) {
			doSearch(p, r, "10")
		})
		router.Get("/artist/:words/:page/:limit", func(p martini.Params, r render.Render) {
			doSearch(p, r, "100")
		})
	})
}
/**
 * 执行搜索并使用render返回json数据
 */
func doSearch(p martini.Params, r render.Render, t string) {
	page, perr := strconv.ParseInt(p["page"], 10, 64)
	if perr != nil {
		page = 1
	}
	limit, lerr := strconv.ParseInt(p["limit"], 10, 64)
	if lerr != nil {
		limit = 10
	}
	// 拿到字符串结果
	reqRs := madoka.Search((string)(p["words"]), t, (int)(page), (int)(limit))
	// 应该可以解析到第一层json
	var originParse map[string] interface{}
	if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	} else {
		r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["result"]})
	}
}
