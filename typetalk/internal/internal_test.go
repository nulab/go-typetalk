package internal

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/shared"
)

func Test_CheckResponse_should_return_invalid_request_error(t *testing.T) {

	resp := &http.Response{}
	resp.StatusCode = 400
	resp.Header = make(map[string][]string)
	resp.Header.Add("WWW-Authenticate", `Bearer error="invalid_request", error_description="Access token is not found"`)

	err := CheckResponse(resp)
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

	err := CheckResponse(resp)
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

	err := CheckResponse(resp)
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

	err := CheckResponse(resp)
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

		if got := SanitizeURL(inURL); !reflect.DeepEqual(got, want) {
			t.Errorf("sanitizeURL(%v) returned %v, want %v", tt.in, got, want)
		}
	}
}

func Test_Client_structToValues_should_convert_struct_to_url_values(t *testing.T) {
	type User struct {
		ID         int        `json:"id"`
		Name       string     `json:"name"`
		Groups     []string   `json:"groups[%d]"`
		CreateDate *time.Time `json:"create_date,omitempty"`
	}
	createDate := time.Date(2018, time.May, 1, 0, 0, 0, 0, time.Local)
	user := User{ID: 9184675, Name: "nu-man", Groups: []string{"red", "blue", "yellow"}, CreateDate: &createDate}
	if values, err := StructToValues(user); err != nil {
		t.Errorf("structToValues failed to convert: %v", err)
	} else {
		if got := values.Get("id"); !reflect.DeepEqual(got, strconv.Itoa(user.ID)) {
			t.Errorf("structToValues returned id %v, want %v", got, user.ID)
		}
		if got := values.Get("name"); !reflect.DeepEqual(got, user.Name) {
			t.Errorf("structToValues returned name %v, want %v", got, user.Name)
		}
		if got := values.Get("groups[0]"); !reflect.DeepEqual(got, user.Groups[0]) {
			t.Errorf("structToValues returned groups[0] %v, want %v", got, user.Groups[0])
		}
		if got := values.Get("groups[1]"); !reflect.DeepEqual(got, user.Groups[1]) {
			t.Errorf("structToValues returned groups[1] %v, want %v", got, user.Groups[1])
		}
		if got := values.Get("groups[2]"); !reflect.DeepEqual(got, user.Groups[2]) {
			t.Errorf("structToValues returned groups[2] %v, want %v", got, user.Groups[2])
		}
		if got := values.Get("create_date"); !reflect.DeepEqual(got, user.CreateDate.Format(time.RFC3339)) {
			t.Errorf("structToValues returned create_date %v, want %v", got, user.CreateDate.Format(time.RFC3339))
		}
	}
}

func Test_Client_addQueries_should_add_queries_to_url(t *testing.T) {
	type Option struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	opt := Option{9184675, "nu-man"}
	if got, err := AddQueries("http://localhost:80/example", opt); err != nil {
		t.Errorf("addQueries failed: %v", err)
	} else {
		want := "http://localhost:80/example?id=9184675&name=nu-man"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("addQueries returned got %v, want %v", got, want)
		}
	}
}
