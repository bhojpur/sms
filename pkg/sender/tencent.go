package sender

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type TencentClient struct {
	core     *sms.Client
	appId    string
	sign     string
	template string
}

func GetTencentClient(accessId string, accessKey string, sign string, templateId string, appId []string) (*TencentClient, error) {
	if len(appId) < 1 {
		return nil, fmt.Errorf("missing parameter: appId")
	}

	credential := common.NewCredential(accessId, accessKey)
	config := profile.NewClientProfile()
	config.HttpProfile.ReqMethod = "POST"

	region := "ap-guangzhou"
	client, err := sms.NewClient(credential, region, config)
	if err != nil {
		return nil, err
	}

	tencentClient := &TencentClient{
		core:     client,
		appId:    appId[0],
		sign:     sign,
		template: templateId,
	}

	return tencentClient, nil
}

func (c *TencentClient) SendMessage(param map[string]string, targetPhoneNumber ...string) error {
	var paramArray []string
	index := 0
	for {
		value := param[strconv.Itoa(index)]
		if len(value) == 0 {
			break
		}
		paramArray = append(paramArray, value)
		index++
	}

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(c.appId)
	request.Sign = common.StringPtr(c.sign)
	request.TemplateParamSet = common.StringPtrs(paramArray)
	request.TemplateID = common.StringPtr(c.template)
	request.PhoneNumberSet = common.StringPtrs(targetPhoneNumber)

	_, err := c.core.SendSms(request)
	return err
}
