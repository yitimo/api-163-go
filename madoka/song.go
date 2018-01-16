package madoka

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// SongInfo 歌曲信息
func SongInfo(ids string) string {
	res, err := http.Get("http://music.163.com/api/song/detail/?ids=" + ids)
	// 错误处理
	if err != nil {
		fmt.Println("Fatal error ", err)
		return `{code: 0}`
	}
	defer res.Body.Close()
	rs, _ := ioutil.ReadAll(res.Body)
	return string(rs)
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
