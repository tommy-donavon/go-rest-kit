package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	consul_register "github.com/yhung-mea7/go-rest-kit/register"
)

type RequestOptions struct {
	ServiceName string
	MethodType  string
	Endpoint    string
	Register    *consul_register.ConsulClient
	Body        []byte
	Headers     map[string]string
}

//send new http request
func SendNewRequest(reqOptions *RequestOptions) (*http.Response, error) {
	if reqOptions.Register == nil {
		return nil, fmt.Errorf("can not look up service with nil consul client")
	}
	ser, err := reqOptions.Register.LookUpService(reqOptions.ServiceName)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(reqOptions.MethodType, ser.GetHTTP()+reqOptions.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	if reqOptions.Body != nil {

		req.Body = ioutil.NopCloser(strings.NewReader(string(reqOptions.Body)))
	}

	for key, value := range reqOptions.Headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	return client.Do(req)

}
