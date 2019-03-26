// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type mockTransport struct {
	gotReq  *http.Request
	gotBody []byte
	results []transportResult
}

type transportResult struct {
	res *http.Response
	err error
}

func (t *mockTransport) addResult(res *http.Response, err error) {
	t.results = append(t.results, transportResult{res, err})
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.gotReq = req
	t.gotBody = nil
	if req.Body != nil {
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		t.gotBody = bytes
	}
	if len(t.results) == 0 {
		return nil, fmt.Errorf("error handling request: t.results is empty")
	}
	result := t.results[0]
	// log.Printf("RountTrip.gotBody=%s\n", string(t.gotBody))
	// log.Printf("RountTrip.result=%v\n", result.res.Body)
	// myBytes, err := ioutil.ReadAll(result.res.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// log.Printf("RountTrip.result=%v\n", string(myBytes))
	// result.res.Body = ioutil.NopCloser(bytes.NewBuffer(myBytes))
	t.results = t.results[1:]
	return result.res, result.err
}

func (t *mockTransport) gotJSONBody() map[string]interface{} {
	m := map[string]interface{}{}
	if err := json.Unmarshal(t.gotBody, &m); err != nil {
		panic(err)
	}
	return m
}

func mockClient(t *testing.T, m *mockTransport) *storage.Client {
	client, err := storage.NewClient(context.Background(), option.WithHTTPClient(&http.Client{Transport: m}))
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func bodyReader(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}