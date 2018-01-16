package madoka

import (
	"strconv"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func get(_url string) (string, error) {
	res, resErr := http.Get(_url)
	// 错误处理
	if resErr != nil {
		return "", resErr
	}
	defer res.Body.Close()
	rs, rsErr := ioutil.ReadAll(res.Body)
	return string(rs), rsErr
}

func post(_url, params, encSecKey string) (string, error) {
	client := &http.Client{}
	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	body := strings.NewReader(form.Encode())
	request, _ := http.NewRequest("POST", _url, body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Content-Length", (string)(body.Len()))
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		return "", reqErr
	}
	defer response.Body.Close()
	resBody, resErr := ioutil.ReadAll(response.Body)
	if resErr != nil {
		return "", resErr
	}
	return string(resBody), nil
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
