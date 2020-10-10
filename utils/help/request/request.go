package request

import (
	"io"
	"io/ioutil"
	"net/http"
)

func DefaultRequest(method, uri string, reqBody io.Reader, reqFunc func(req *http.Request) error) (body []byte, err error) {
	req, err := http.NewRequest(method, uri, reqBody)
	if err != nil {
		return body, err
	}

	if reqFunc != nil {
		if err = reqFunc(req); err != nil {
			return body, err
		}
	} else {
		if err = setReqExtra(req); err != nil {
			return body, err
		}
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return body, err
	}

	defer rsp.Body.Close()

	body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func setReqExtra(req *http.Request) (err error) {
	return err
}
