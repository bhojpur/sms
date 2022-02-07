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
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

type VolcClient struct {
	core       *sms.SMS
	sign       string
	template   string
	smsAccount string
}

func GetVolcClient(accessId, accessKey, sign, templateId string, smsAccount []string) (*VolcClient, error) {
	if len(smsAccount) < 1 {
		return nil, fmt.Errorf("missing parameter: smsAccount")
	}

	client := sms.NewInstance()
	client.Client.SetAccessKey(accessId)
	client.Client.SetSecretKey(accessKey)

	volcClient := &VolcClient{
		core:       client,
		sign:       sign,
		template:   templateId,
		smsAccount: smsAccount[0],
	}

	return volcClient, nil
}

func (c *VolcClient) SendMessage(param map[string]string, targetPhoneNumber ...string) error {
	requestParam, err := json.Marshal(param)
	if err != nil {
		return err
	}

	if len(targetPhoneNumber) < 1 {
		return fmt.Errorf("missing parameter: targetPhoneNumber")
	}

	phoneNumbers := bytes.Buffer{}
	phoneNumbers.WriteString(targetPhoneNumber[0])
	for _, s := range targetPhoneNumber[1:] {
		phoneNumbers.WriteString(",")
		phoneNumbers.WriteString(s)
	}

	req := &sms.SmsRequest{
		SmsAccount:    c.smsAccount,
		Sign:          c.sign,
		TemplateID:    c.template,
		TemplateParam: string(requestParam),
		PhoneNumbers:  phoneNumbers.String(),
	}
	_, _, err = c.core.Send(req)
	return err
}
