package shared

import (
	"net/http"
	"testing"
)

func Test_ErrorResponse_Error(t *testing.T) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{ErrorType: "error type", Response: res}
	if err.Error() == "" {
		t.Error("Expected non-empty ErrorResponse.Error()")
	}
}
