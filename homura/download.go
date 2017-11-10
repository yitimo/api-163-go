package homura

import (
	"encoding/json"

	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// DownloadGroupInit 初始化下载API组
func DownloadGroupInit(m *martini.ClassicMartini) {
	m.Group("/download", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			getDownloadURL(p["id"], "320000", r)
		})
		router.Get("/low/:id", func(p martini.Params, r render.Render) {
			getDownloadURL(p["id"], "160000", r)
		})
		router.Get("/middle/:id", func(p martini.Params, r render.Render) {
			getDownloadURL(p["id"], "320000", r)
		})
		router.Get("/high/:id", func(p martini.Params, r render.Render) {
			getDownloadURL(p["id"], "640000", r)
		})
	})
}

/**
 * 执行搜索并使用render返回json数据
 */
func getDownloadURL(id string, rate string, r render.Render) {
	initStr := `{"ids": "[` + id + `]", "br": "` + rate + `", "csrf_token": ""}`
	params, key, err := madoka.EncParams(initStr)
	if err != nil {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	}
	// 发送POST请求得到最后包含url的结果
	reqRs := madoka.Download(params, key)
	// 应该可以解析到第一层json
	var originParse map[string]interface{}
	if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	} else {
		r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["data"]})
	}
}
