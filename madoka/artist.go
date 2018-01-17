package madoka

// ArtistTop - get top list of artists
func ArtistTop(page, limit int) (string, error) {
	// 1. transfer page, limit into offset, limit
	_offset, _limit := formatParams(page, limit)
	// 2. encode request params
	preParams := "{\"offset\": "+ _offset +", \"limit\": "+_limit +", \"total\": true, \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	// 3. request, resolve, return
	res, resErr := post("http://music.163.com/weapi/artist/top?csrf_token=", params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}

// ArtistSong - get songs if artist by id
func ArtistSong(id string, page, limit int) (string, error) {
	// 1. transfer page, limit into offset, limit
	_offset, _limit := formatParams(page, limit)
	// 2. encode request params
	preParams := "{\"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	// 3. request, resolve, return
	res, resErr := post("http://music.163.com/weapi/v1/artist/"+id+"?offset="+_offset+"&limit="+_limit, params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
