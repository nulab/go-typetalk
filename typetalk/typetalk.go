package typetalk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	libraryVersion = "1.2.0"
	apiVersion     = "v1"
	defaultBaseURL = "https://typetalk.com/api/"
	userAgent      = "go-typetalk/" + libraryVersion

	defaultMediaType = "application/octet-stream"
)

type service struct {
	client *Client
}

type Client struct {
	client *http.Client

	BaseURL       *url.URL
	UserAgent     string
	typetalkToken string

	common service

	Accounts      *AccountsService
	Files         *FilesService
	Mentions      *MentionsService
	Messages      *MessagesService
	Notifications *NotificationsService
	Organizations *OrganizationsService
	Talks         *TalksService
	Topics        *TopicsService
	Likes         *LikesService
}

func (c *Client) SetTypetalkToken(token string) *Client {
	c.typetalkToken = token
	return c
}

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		if values, err := structToValues(body); err != nil {
			return nil, err
		} else {
			buf = strings.NewReader(values.Encode())
		}

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
	if c.typetalkToken != "" {
		req.Header.Set("X-Typetalk-Token", c.typetalkToken)
	}
	return req, nil
}

func (c *Client) newUploadRequest(urlStr string, reader io.Reader, size int64, mediaType string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("POST", url.String(), reader)
	if err != nil {
		return nil, err
	}
	req.ContentLength = size

	if mediaType == "" {
		mediaType = defaultMediaType
	}
	req.Header.Set("Content-Type", mediaType)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.typetalkToken != "" {
		req.Header.Set("X-Typetalk-Token", c.typetalkToken)
	}
	return req, nil

}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		// refs: https://github.com/google/go-github/pull/317
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := &Response{Response: resp}

	err = checkResponse(resp)
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

func (c *Client) call(ctx context.Context, method string, url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, req, &v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) post(ctx context.Context, url string, body interface{}, v interface{}) (*Response, error) {
	return c.call(ctx, "POST", url, body, v)
}

func (c *Client) put(ctx context.Context, url string, body interface{}, v interface{}) (*Response, error) {
	return c.call(ctx, "PUT", url, body, v)
}

func (c *Client) delete(ctx context.Context, url string, v interface{}) (*Response, error) {
	return c.call(ctx, "DELETE", url, nil, v)
}

func (c *Client) get(ctx context.Context, url string, v interface{}) (*Response, error) {
	return c.call(ctx, "GET", url, nil, v)
}

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
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.ErrorType, r.ErrorDescription)
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion + "/")

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.Accounts = (*AccountsService)(&c.common)
	c.Files = (*FilesService)(&c.common)
	c.Mentions = (*MentionsService)(&c.common)
	c.Messages = (*MessagesService)(&c.common)
	c.Notifications = (*NotificationsService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Talks = (*TalksService)(&c.common)
	c.Topics = (*TopicsService)(&c.common)
	c.Likes = (*LikesService)(&c.common)

	return c
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
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

func structToValues(data interface{}) (url.Values, error) {
	result := make(map[string]interface{})
	b, _ := json.Marshal(data)
	d := json.NewDecoder(strings.NewReader(string(b)))
	d.UseNumber()
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	values := url.Values{}
	for k, v := range result {
		if as, ok := v.([]interface{}); ok {
			for i, v := range as {
				values.Add(fmt.Sprintf(k, i), fmt.Sprintf("%v", v))
			}
		} else {
			values.Add(k, fmt.Sprintf("%v", v))
		}
	}
	return values, nil
}

func addQueries(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	if values, err := structToValues(opt); err != nil {
		return s, err
	} else {
		u.RawQuery = values.Encode()
		return u.String(), nil
	}
}

func sanitizeURL(uri *url.URL) *url.URL {
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
