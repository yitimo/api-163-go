package madoka

// TopArtist - get top list of artists
func TopArtist(page, limit int) (string, error) {
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

// Artist - get artist by id with hot songs
func Artist(id string, page, limit int) (string, error) {
	_offset, _limit := formatParams(page, limit)
	preParams := "{\"offset\": "+ _offset +", \"limit\": "+_limit +", \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/v1/artist/"+id, params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}

// ArtistAlbum - album of artist
func ArtistAlbum(id string, page, limit int) (string, error) {
	_offset, _limit := formatParams(page, limit)
	preParams := "{\"offset\": "+ _offset +", \"limit\": "+_limit +", \"total\": true, \"csrf_token\": \"\"}"
	params, encSecKey, encErr := EncParams(preParams)
	if encErr != nil {
		return "", encErr
	}
	res, resErr := post("http://music.163.com/weapi/artist/albums/"+id, params, encSecKey)
	if resErr != nil {
		return "", resErr
	}
	return res, nil
}
