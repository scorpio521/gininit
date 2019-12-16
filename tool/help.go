package tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"time"
)

const (
	regular = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
)

func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func FromatData(code interface{}, data interface{}, msg interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	res["status"] = code
	res["data"] = data
	res["msg"] = msg

	return
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToMapViaJson(obj interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)
	return
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Alihost() (host string) {
	host = "http://wmfrontdata.oss-cn-beijing.aliyuncs.com/"
	return
}
