package handler

import (
	"net/http"
	"net/http/httptest"
)

func TestHandler(t *testing.t) {
	srv := httptest.NewServer(http.HandlerFunc(handler))
	resp, err := http.Get(srv.URL)
	if err=nil{
		t.Log(err)
		t.Fail()

	}
	textBytes,err:=ioutil.ReadAll(resp.Body)
	if err=nil{
		t.Log(err)
		t.Fail()

	}
	defer resp.Body.Close()
	text:=string(textBytes)
	if text!="message from test handler"{
		t.Log(err)
		t.Fail()
	}

}
//go test -v -count=1 ./handler/...

