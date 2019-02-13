package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

//检验字段fieldName的值value 的长度
func CheckMaxLength(value string, maxLength int, fieldName string) {

	if len(value) > maxLength {
		panic(fmt.Sprintf("Invalid Arguments:the length of %s can not be larger than %d.", fieldName, maxLength))
	}
}

//检验字段fieldName的值value的最大列表长度
func CheckMaxListSize(value string, maxSize int, fieldName string) {

	if !CheckEmpty(value) && len(strings.Split(value, ",")) > maxSize {
		panic(fmt.Sprintf("Invalid Arguments:the listsize(the string split by \",\") of %s must be less than %d .", fieldName, maxSize))
	}
}

//检测最大值
func CheckMaxValue(value string, maxValue int, fieldName string) {
	if CheckEmpty(value) {
		return
	}
	if i, e := strconv.Atoi(value); e != nil {
		panic("fieldName cid Value Is Not Number")
	} else {
		if i > maxValue {
			panic(fmt.Sprintf("Invalid Arguments:the value of %s can not be larger than %d .", fieldName, maxValue))
		}
	}
}

//检测最小值
func CheckMinValue(value string, minVal int, fieldName string) {

	if CheckEmpty(value) {
		return
	}
	if i, e := strconv.Atoi(value); e != nil {
		panic("fieldName cid Value Is Not Number")
	} else {
		if i < minVal {
			panic(fmt.Sprintf("Invalid Arguments:the value of %s can not be less than %d .", fieldName, minVal))
		}
	}

}

//检测非空
func CheckNotNull(value, fieldName string) {
	if CheckEmpty(value) {
		panic(fmt.Sprintf("Missing Required Arguments: %s", fieldName))
	}
}

//检测是否为空 空则返回真
func CheckEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}
//检测是否为数字
func CheckNumber(value, fieldName string) {
	compile := regexp.MustCompile(`\d+`)
	if !compile.MatchString(value) {
		panic(fmt.Sprintf("Invalid Arguments:the value of %s is not number : %s.", fieldName, value))
	}
}
