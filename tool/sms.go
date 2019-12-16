package tool

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/pkg/errors"
)

func SendSms(phone string, templateCode string, code string) (flag bool, err error) {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "", "")
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["SignName"] = "阿朵生活"
	request.QueryParams["TemplateCode"] = templateCode
	templateParam := map[string]string{"code": code}
	mjson, _ := json.Marshal(templateParam)
	mString := string(mjson)
	request.QueryParams["TemplateParam"] = mString
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(response.GetHttpContentString())
	jsonMap := make(map[string]interface{})
	errs := json.Unmarshal([]byte(response.GetHttpContentString()), &jsonMap)
	if errs != nil {
		flag = false
		err = errors.New("发送失败！")
	}
	if err != nil {
		flag = false
		err = errors.New(jsonMap["Message"].(string))
	}
	if jsonMap["Code"] == "OK" {
		flag = true
		err = nil
	} else {
		flag = false
		err = errors.New(jsonMap["Message"].(string))
	}

	return
}
