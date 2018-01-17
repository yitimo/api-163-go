package madoka

// SongInfo 歌曲信息
func SongInfo(ids string) (string, error) {
	preParams := "{\"ids\": \"" + ids + "\", \"c\": \""+formatIdc(ids)+"\", \"csrf_token\": \"\"}"
	params, encSecKey, err := EncParams(preParams)
	if err != nil {
		return "", err
	}
	res, resErr := post("http://music.163.com/weapi/v3/song/detail", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}

// SongLyric 歌词信息
func SongLyric(id string) (string, error) {
	preParams := "{\"id\": \"" + id + "\", \"lv\": -1, \"csrf_token\": \"\"}"
	params, encSecKey, err := EncParams(preParams)
	if err != nil {
		return "", err
	}
	// 3. request, resolve, return
	res, resErr := post("http://music.163.com/weapi/song/lyric?csrf_token=", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
