package controller_test

import (
	"testing"

	"github.com/imroc/req/v3"
)

func init() {
	req.DevMode()
	// req.DevMode

}

var baseurl = "http://127.0.0.1:10080/api"

func TestXxx(t *testing.T) {
	b := ``
	req.SetBodyJsonString(b).Post(baseurl)
}

func TestTextGET(t *testing.T) {
	b := `/text?d=test`
	req.Get(baseurl + b)
}

// func TestTextPost(t *testing.T) {
// 	d := `{"Text":"test post "}`
// 	req.SetBodyJsonString(d).Post(baseurl + "/text")
// }

func TestPost(t *testing.T) {
	d := `{"d":"test post "}`
	req.SetBodyJsonString(d).Post(baseurl)
}

func TestGET(t *testing.T) {

	// b := req.AddCommonQueryParams().
	// QueryParams.Encode()
	req.Get(baseurl)
}
