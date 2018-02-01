package madoka

// PlayList - get playlist by cat&order
func PlayList(cat, order string, page, limit int) (string, error) {
	_offset, _limit := formatParams(page, limit)
	preParams := "{\"cat\":\"" + cat + "\", \"order\":\"" + order + "\", \"offset\":\"" + _offset + "\", \"limit\":\"" + _limit + "\", \"total\": \"true\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/playlist/list", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}

// PlayListDetail - get playlist detail by id
func PlayListDetail(id string) (string, error) {
	preParams := "{\"id\": \"" + id + "\", \"n\": 100000, \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/api/playlist/detail?id="+id, params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}

// PlayListCatalogue - get cat list
func PlayListCatalogue() (string, error) {
	preParams := "{\"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/playlist/catalogue", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
