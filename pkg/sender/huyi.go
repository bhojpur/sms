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
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type HuyiClient struct {
	appId    string
	appKey   string
	template string
}

func GetHuyiClient(appId string, appKey string, template string) (*HuyiClient, error) {
	return &HuyiClient{
		appId:    appId,
		appKey:   appKey,
		template: template,
	}, nil
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (hc *HuyiClient) SendMessage(param map[string]string, targetPhoneNumber ...string) error {
	code, ok := param["code"]
	if !ok {
		return fmt.Errorf("missing parameter: msg code")
	}

	if len(targetPhoneNumber) < 1 {
		return fmt.Errorf("missin parer: trgetPhoneNumber")
	}

	_now := strconv.FormatInt(time.Now().Unix(), 10)
	smsContent := fmt.Sprintf(hc.template, code)
	v := url.Values{}
	v.Set("account", hc.appId)
	v.Set("content", smsContent)
	v.Set("time", _now)
	passwordStr := hc.appId + hc.appKey + "%s" + smsContent + _now
	for _, mobile := range targetPhoneNumber {
		password := fmt.Sprintf(passwordStr, mobile)
		v.Set("password", GetMd5String(password))
		v.Set("mobile", mobile)

		body := strings.NewReader(v.Encode()) //encode form data
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		resp, err := client.Do(req) // request remote
		if err != nil {
			return err
		}
		defer resp.Body.Close() // ！ close ReadCloser
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
