package homura

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// DownloadGroupInit 初始化下载API组
func DownloadGroupInit(m *martini.ClassicMartini) {
	m.Group("/download", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			getDownloadURL("["+p["id"]+"]", "320000", r)
		})
		router.Get("/low/:id", func(p martini.Params, r render.Render) {
			getDownloadURL("["+p["id"]+"]", "160000", r)
		})
		router.Get("/middle/:id", func(p martini.Params, r render.Render) {
			getDownloadURL("["+p["id"]+"]", "320000", r)
		})
		router.Get("/high/:id", func(p martini.Params, r render.Render) {
			getDownloadURL("["+p["id"]+"]", "640000", r)
		})
		router.Post("/multi", func(req *http.Request, r render.Render) {
			req.ParseForm()
			for data, _ := range req.Form {
				var pParse map[string]interface{}
				if err := json.Unmarshal([]byte(data), &pParse); err == nil {
					getDownloadURL(formatIds(pParse["ids"].([]interface{})), "320000", r)
				} else {
					r.JSON(200, map[string]interface{}{"state": false, "msg": "参数错误", "data": nil})
				}
			}
		})
	})
}

func formatIds(ids []interface{}) string {
	rs := "["
	for _, id := range ids {
		rs += fmt.Sprintf("%.0f", id) + ","
	}
	rs = rs[0 : len(rs)-1]
	rs += "]"
	return rs
}

/**
 * 执行搜索并使用render返回json数据
 */
func getDownloadURL(ids string, rate string, r render.Render) {
	// 发送POST请求得到最后包含url的结果
	reqRs, reqErr := madoka.Download(ids, rate)
	if reqErr != nil {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
		return
	}
	// 应该可以解析到第一层json
	var originParse map[string]interface{}
	if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	} else {
		r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["data"]})
	}
}
