package taobaoopensdk

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
)

//参数排序返回排序的KEY 通过KEY循环取出
func SortParamters(params url.Values) []string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func Md5(data string) string {
	md5obj := md5.New()
	if _, err := md5obj.Write([]byte(data)); err != nil {
		return ""
	}
	md5Data := md5obj.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}
