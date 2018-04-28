package shared

import (
	"fmt"
	"net/http"
	"net/url"
)

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response         *http.Response
	ErrorType        string `json:"error"`
	ErrorDescription string `json:"errorDescription"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, SanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.ErrorType, r.ErrorDescription)
}

func SanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("typetalkToken")) > 0 {
		params.Set("typetalkToken", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
