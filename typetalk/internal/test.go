package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type Values map[string]interface{}

func TestFormValues(t *testing.T, r *http.Request, values Values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, fmt.Sprintf("%v", v))
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters:\n got  %v,\n want %v", got, want)
	}
}

func TestQueryValues(t *testing.T, r *http.Request, values Values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, fmt.Sprintf("%v", v))
	}

	u, _ := url.Parse(r.RequestURI)
	if got := u.Query(); !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters:\n got  %v,\n want %v", got, want)
	}
}

func TestHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned\n %q, \n want %q", header, got, want)
	}
}
