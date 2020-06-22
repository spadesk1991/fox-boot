package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func call(method, uri string, params map[string]string, body interface{}) (buff []byte, err error) {
	url := uri
	ps := make([]string, 0)
	if params != nil {
		for k, v := range params {
			ps = append(ps, fmt.Sprintf("&%s=%s", k, v))
		}
		url += strings.Join(ps, "&")
	}
	var in io.Reader
	if body != nil {
		var bodyBf []byte
		bodyBf, err = json.Marshal(body)
		if err != nil {
			return
		}
		in = bytes.NewBuffer(bodyBf)
	}
	request, err := http.NewRequest(method, url, in)
	if err != nil {
		return
	}
	client := http.DefaultClient
	r, err := client.Do(request)
	defer r.Body.Close()
	if err != nil {
		return
	}
	buff, err = ioutil.ReadAll(r.Body)
	return
}

func Post(uri string, params map[string]string, body interface{}) (result []byte, err error) {
	result, err = call("POST", uri, params, body)
	return
}

func Get(uri string, params map[string]string) (result []byte, err error) {
	result, err = call("GET", uri, params, nil)
	return
}

func Put(uri string, params map[string]string, body interface{}) (result []byte, err error) {
	result, err = call("PUT", uri, params, body)
	return
}

func Delete(uri string, params map[string]string) (result []byte, err error) {
	result, err = call("DELETE", uri, params, nil)
	return
}
