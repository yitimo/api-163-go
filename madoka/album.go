package madoka

func Album(id string) (string, error) {
	preParams := "{\"id\": " + id + ", \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/v1/album/"+id, params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
