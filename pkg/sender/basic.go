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

import "fmt"

const (
	Aliyun       = "Aliyun SMS"
	TencentCloud = "Tencent Cloud SMS"
	VolcEngine   = "Volc Engine SMS"
	Huyi         = "Huyi SMS"
)

type SmsClient interface {
	SendMessage(param map[string]string, targetPhoneNumber ...string) error
}

func NewSmsClient(provider string, accessId string, accessKey string, sign string, template string, other ...string) (SmsClient, error) {
	switch provider {
	case Aliyun:
		return GetAliyunClient(accessId, accessKey, sign, template)
	case TencentCloud:
		return GetTencentClient(accessId, accessKey, sign, template, other)
	case VolcEngine:
		return GetVolcClient(accessId, accessKey, sign, template, other)
	case Huyi:
		return GetHuyiClient(accessId, accessKey, template)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
