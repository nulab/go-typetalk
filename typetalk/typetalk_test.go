package typetalk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]interface{}

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, fmt.Sprintf("%v", v))
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters:\n got  %v,\n want %v", got, want)
	}
}

func testQueryValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, fmt.Sprintf("%v", v))
	}

	u, _ := url.Parse(r.RequestURI)
	if got := u.Query(); !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters:\n got  %v,\n want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned\n %q, \n want %q", header, got, want)
	}
}

func Test_CheckResponse_should_return_invalid_request_error(t *testing.T) {

	resp := &http.Response{}
	resp.StatusCode = 400
	resp.Header = make(map[string][]string)
	resp.Header.Add("WWW-Authenticate", `Bearer error="invalid_request", error_description="Access token is not found"`)

	err := checkResponse(resp)
	if err == nil {
		t.Error("error is nil")
	}
	switch e := err.(type) {
	case *ErrorResponse:
		if e.ErrorType != "invalid_request" {
			t.Errorf("error type: want %s, got %s", "invalid_request", e.ErrorType)
		}
	default:
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CheckResponse_should_return_access_token_is_not_found_error(t *testing.T) {

	resp := &http.Response{}
	resp.StatusCode = 400
	resp.Header = make(map[string][]string)
	resp.Header.Add("WWW-Authenticate", `Bearer error="invalid_token", error_description="The access token is not found"`)

	err := checkResponse(resp)
	if err == nil {
		t.Error("error is nil")
	}
	switch e := err.(type) {
	case *ErrorResponse:
		if e.ErrorType != "invalid_token" || e.ErrorDescription != "The access token is not found" {
			t.Errorf("error type: want %s, got %s", "invalid_token", e.ErrorType)
		}
	default:
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CheckResponse_should_return_access_token_expired_error(t *testing.T) {

	resp := &http.Response{}
	resp.StatusCode = 400
	resp.Header = make(map[string][]string)
	resp.Header.Add("WWW-Authenticate", `Bearer error="invalid_token", error_description="The access token expired"`)

	err := checkResponse(resp)
	if err == nil {
		t.Error("error is nil")
	}
	switch e := err.(type) {
	case *ErrorResponse:
		if e.ErrorType != "invalid_token" || e.ErrorDescription != "The access token expired" {
			t.Errorf("error type: want %s, got %s", "invalid_token", e.ErrorType)
		}
	default:
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CheckResponse_should_return_invalid_scope_error(t *testing.T) {

	resp := &http.Response{}
	resp.StatusCode = 400
	resp.Header = make(map[string][]string)
	resp.Header.Add("WWW-Authenticate", `Bearer error="invalid_scope"`)

	err := checkResponse(resp)
	if err == nil {
		t.Error("error is nil")
	}
	switch e := err.(type) {
	case *ErrorResponse:
		if e.ErrorType != "invalid_scope" {
			t.Errorf("error type: want %s, got %s", "invalid_scope", e.ErrorType)
		}
	default:
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_sanitizeURL_should_sanitize_typetalk_token_value(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"/?a=b", "/?a=b"},
		{"/?a=b&typetalkToken=secret", "/?a=b&typetalkToken=REDACTED"},
		{"/?a=b&client_id=id&typetalkToken=secret", "/?a=b&client_id=id&typetalkToken=REDACTED"},
	}

	for _, tt := range tests {
		inURL, _ := url.Parse(tt.in)
		want, _ := url.Parse(tt.want)

		if got := sanitizeURL(inURL); !reflect.DeepEqual(got, want) {
			t.Errorf("sanitizeURL(%v) returned %v, want %v", tt.in, got, want)
		}
	}
}

func Test_ErrorResponse_Error(t *testing.T) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{ErrorType: "error type", Response: res}
	if err.Error() == "" {
		t.Error("Expected non-empty ErrorResponse.Error()")
	}
}

func Test_Client_newRequest_should_add_typetalk_token_to_header_if_use_SetTypetalkToken(t *testing.T) {
	req, _ := NewClient(nil).SetTypetalkToken("mytypetalktoken").newRequest("GET", "example", nil)
	if token := req.Header.Get("X-Typetalk-Token"); token != "mytypetalktoken" {
		t.Errorf("Invalid Typetalk Token: %s", token)
	}
}
