package typetalk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
)

type FilesService service

type AttachmentFile struct {
	ContentType string `json:"contentType"`
	FileKey     string `json:"fileKey"`
	FileName    string `json:"fileName"`
	FileSize    int    `json:"fileSize"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/upload-attachment
func (s *FilesService) UploadAttachmentFile(ctx context.Context, topicId int, file *os.File) (*AttachmentFile, *Response, error) {
	u := fmt.Sprintf("topics/%v/attachments", topicId)
	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	if stat.IsDir() {
		return nil, nil, errors.New("to upload can't be a directory")
	}

	mediaType := mime.TypeByExtension(filepath.Ext(file.Name()))
	req, err := s.client.newUploadRequest(u, file, stat.Size(), mediaType)
	if err != nil {
		return nil, nil, err
	}

	attachmentFile := &AttachmentFile{}
	if resp, err := s.client.do(ctx, req, attachmentFile); err != nil {
		return nil, resp, err
	} else {
		return attachmentFile, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/download-attachment
func (s *FilesService) DownloadAttachmentFile(ctx context.Context, topicId, postId, attachmentId int, filename string) (io.ReadCloser, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/attachments/%d/%s", topicId, postId, attachmentId, filename)

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", defaultMediaType)

	resp, err := s.client.client.Do(req)
	if err := checkResponse(resp); err != nil {
		resp.Body.Close()
		return nil, err
	} else {
		return resp.Body, nil
	}
}
