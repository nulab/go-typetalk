package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"bytes"
	"mime/multipart"
	"os"

	"github.com/nulab/go-typetalk/typetalk/shared"
)

const (
	version        = "3.1.0"
	DefaultBaseURL = "https://typetalk.com/api/"
	UserAgent      = "go-typetalk/" + version

	DefaultMediaType = "application/octet-stream"
)

type ClientCore struct {
	Client *http.Client

	BaseURL       *url.URL
	UserAgent     string
	TypetalkToken string
}

func (c *ClientCore) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		values, err := StructToValues(body)
		if err != nil {
			return nil, err
		}
		buf = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.TypetalkToken != "" {
		req.Header.Set("X-Typetalk-Token", c.TypetalkToken)
	}
	return req, nil
}

func (c *ClientCore) NewMultipartRequest(urlStr string, values map[string]io.Reader) (*http.Request, error) {
	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)
	for key, reader := range values {
		if closable, ok := reader.(io.Closer); ok {
			defer closable.Close()
		}
		var fieldWriter io.Writer
		var err error
		if file, ok := reader.(*os.File); ok {
			if fieldWriter, err = multipartWriter.CreateFormFile(key, file.Name()); err != nil {
				return nil, err
			}
		} else {
			if fieldWriter, err = multipartWriter.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fieldWriter, reader); err != nil {
			return nil, err
		}
	}
	multipartWriter.Close()
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	resolvedURL := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodPost, resolvedURL.String(), &buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.TypetalkToken != "" {
		req.Header.Set("X-Typetalk-Token", c.TypetalkToken)
	}
	return req, nil
}

func (c *ClientCore) NewUploadRequest(urlStr string, reader io.Reader, size int64, mediaType string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	resolvedURL := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodPost, resolvedURL.String(), reader)
	if err != nil {
		return nil, err
	}
	req.ContentLength = size

	if mediaType == "" {
		mediaType = DefaultMediaType
	}
	req.Header.Set("Content-Type", mediaType)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.TypetalkToken != "" {
		req.Header.Set("X-Typetalk-Token", c.TypetalkToken)
	}
	return req, nil

}

func (c *ClientCore) Do(ctx context.Context, req *http.Request, v interface{}) (*shared.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if e, ok := err.(*url.Error); ok {
			if parsedURL, err := url.Parse(e.URL); err == nil {
				e.URL = shared.SanitizeURL(parsedURL).String()
				return nil, e
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	response := &shared.Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				// ignore EOF errors caused by empty response body
				err = nil
			}
		}
	}

	return response, err
}

func (c *ClientCore) Call(ctx context.Context, method string, url string, body interface{}, v interface{}) (*shared.Response, error) {
	req, err := c.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(ctx, req, &v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *ClientCore) Post(ctx context.Context, url string, body interface{}, v interface{}) (*shared.Response, error) {
	return c.Call(ctx, http.MethodPost, url, body, v)
}

func (c *ClientCore) Put(ctx context.Context, url string, body interface{}, v interface{}) (*shared.Response, error) {
	return c.Call(ctx, http.MethodPut, url, body, v)
}

func (c *ClientCore) Delete(ctx context.Context, url string, v interface{}) (*shared.Response, error) {
	return c.Call(ctx, http.MethodDelete, url, nil, v)
}

func (c *ClientCore) Get(ctx context.Context, url string, v interface{}) (*shared.Response, error) {
	return c.Call(ctx, http.MethodGet, url, nil, v)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &shared.ErrorResponse{Response: r}
	if c := r.StatusCode; c == 400 || c == 401 {
		errorStr := r.Header.Get("WWW-Authenticate")
		if errorStr == "" {
			return errorResponse
		}
		for _, v := range strings.Split(errorStr, ",") {
			errors := strings.Split(v, "=")
			if len(errors) != 2 {
				continue
			}
			k := errors[0]
			v := strings.Trim(errors[1], `"`)
			if strings.Contains(k, "error_description") {
				errorResponse.ErrorDescription = v
			} else if strings.Contains(k, "error") {
				errorResponse.ErrorType = v
			}
		}
	}
	return errorResponse
}

func StructToValues(data interface{}) (url.Values, error) {
	result := make(map[string]interface{})
	b, _ := json.Marshal(data)
	d := json.NewDecoder(strings.NewReader(string(b)))
	d.UseNumber()
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	values := url.Values{}
	for k, v := range result {
		addValues(values, k, v)
	}
	return values, nil
}

func addValues(values url.Values, k string, v interface{}) {
	if as, ok := v.([]interface{}); ok {
		for i, v := range as {
			addValues(values, fmt.Sprintf(k, i), v)
		}
	} else {
		values.Add(k, fmt.Sprintf("%v", v))
	}
}

func AddQueries(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	values, err := StructToValues(opt)
	if err != nil {
		return s, err
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
