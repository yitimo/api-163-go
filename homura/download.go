package homura

import (
	"fmt"
	"bytes"
	"github.com/go-martini/martini"
	"../madoka"
	"encoding/json"
	"github.com/martini-contrib/render"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"math"
	// "strconv"
	"math/big"
)

/**
 * 初始化搜索路由组
 */
 func DownloadGroupInit(m *martini.ClassicMartini) {
	m.Group("/download", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "320000", r)
		})
		router.Get("/low/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "160000", r)
		})
		router.Get("/middle/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "320000", r)
		})
		router.Get("/high/:id", func(p martini.Params, r render.Render) {
			getDownloadUrl(p["id"], "640000", r)
		})
	})
}
/**
 * 执行搜索并使用render返回json数据
 */
func getDownloadUrl(id string, rate string, r render.Render) {
	initStr := `{"ids": "[` + id + `]", "br": "128000", "csrf_token": ""}`
	params, key, err := getParams(initStr)
	if err != nil {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	}
	// 发送POST请求得到最后包含url的结果
	reqRs := madoka.Download(params, key)
	// 应该可以解析到第一层json
	var originParse map[string] interface{}
	if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
		r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
	} else {
		r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse["data"]})
	}
}

func getParams(text string) (string, string, error) {
	modulus := "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	nonce := "0CoJUm6Qyw8W8jud"
	pubKey := "010001"
	secKey := createSecretKey(16);
	if aes1, err1 := aesEncrypt(text, nonce); err1 != nil {
		return "", "", err1
	} else {
		if aes2, err2 := aesEncrypt(aes1, secKey); err2 != nil {
			return "", "", err2
		} else {
			return aes2, rsaEncrypt(secKey, pubKey, modulus), nil
		}
	}
}

func createSecretKey(size int) string {
	keys := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/";
	rs := "";
	for i := 0; i < size; i++ {
		pos := rand.Intn(len(keys))
		rs += keys[pos:pos+1]
	}
	return rs
}

func aesEncrypt(sSrc string, sKey string) (string, error) {
	iv := []byte("0102030405060708")
	block, err := aes.NewCipher([]byte(sKey))
	if err != nil {
		return "", err
	}
	padding := block.BlockSize() - len([]byte(sSrc)) % block.BlockSize()
	src := append([]byte(sSrc), bytes.Repeat([]byte{byte(padding)}, padding)...)
	model := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(src))
	model.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText) , nil
}

func rsaEncrypt(text string, pubKey string, modulus string) string {
	// 倒序
	rText := ""
	for i := len(text) - 1; i >= 0; i-- {
		rText += text[i:i+1]
	}
	// 打印python下的三个pow参数
	// 对比golang下的三个pow参数
	// pow结果如何转16进制字符串然后填充0
	// 字符串转ascii
	textParse := ""
	for _, char := range []rune(rText) {
		textParse += fmt.Sprintf("%x", int(char))
	}
	i1, _ := big.NewInt(0).SetString(textParse, 16)
	i2, _ := big.NewInt(0).SetString(pubKey, 16)
	i3, _ := big.NewInt(0).SetString(modulus, 16)
	rs := i1.Exp(i1, i2, i3)
	strRs := fmt.Sprintf("%x", rs)
	return addPadding(strRs, modulus)
}

func addPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--;
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}

func pow(x float64, n float64, lim float64) int64 {
    return int64(math.Pow(x, n)) % int64(lim)
}
