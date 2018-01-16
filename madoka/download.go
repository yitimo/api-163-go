package madoka

// Download 根据传入id返回生成的mp3地址
func Download(ids, rate string) (string, error) {
	preParams := "{\"ids\": \"" + ids + "\", \"br\": \"" + rate + "\", \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	// 3. request, resolve, return
	res, resErr := post("http://music.163.com/weapi/song/enhance/player/url?csrf_token=", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}