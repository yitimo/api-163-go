package madoka

// Search 执行搜索 params: 关键词 类型 页码 数量 return: 字符串形式的请求结果
func Search(words string, stype string, page int, limit int) (string, error) {
	_offset, _limit := formatParams(page, limit)
	preParams := "{\"s\": \"" + words + "\", \"type\": \"" + stype + "\", \"offset\": " + _offset + ", \"limit\": " + _limit + ", \"total\": true, \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/search/get", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
