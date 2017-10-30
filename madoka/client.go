package madoka

import (
	"os"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"strconv"
)

/** 
 * 执行搜索
 * params: 	关键词 类型 页码 数量
 * return:	字符串形式的请求结果
 */
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
	if reqErr!= nil {
		fmt.Println("Fatal error ", reqErr.Error())
		os.Exit(0)
	}
	defer response.Body.Close()
	resBody, _ := ioutil.ReadAll(response.Body)
	return string(resBody)
}

/**
* 传入 搜索类型 页码 数量
* 返回 搜索类型 偏移 数量
*/
func formatParams(page int, limit int) (string, string) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 0
	}
	return strconv.Itoa((page - 1) * limit), strconv.Itoa(limit)
}
