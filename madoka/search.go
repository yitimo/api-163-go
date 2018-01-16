package madoka

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Search 执行搜索 params: 关键词 类型 页码 数量 return: 字符串形式的请求结果
func Search(words string, stype string, page int, limit int) string {
	// 创建客户端
	client := &http.Client{}
	// 格式化参数
	_o, _l := formatParams(page, limit)
	// 设置body
	form := url.Values{}
	form.Set("s", words)
	form.Set("type", stype)
	form.Set("limit", _l)
	form.Set("offset", _o)
	body := strings.NewReader(form.Encode())
	// 创建请求
	request, _ := http.NewRequest("POST", "http://music.163.com/api/search/get/", body)
	//设置头部
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "appver=2.0.2")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Content-Length", (string)(body.Len()))
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		fmt.Println("Fatal error ", reqErr.Error())
		return `{"data": null, "state": false, "msg": "请求失败"}`
	}
	defer response.Body.Close()
	resBody, _ := ioutil.ReadAll(response.Body)
	return string(resBody)
}