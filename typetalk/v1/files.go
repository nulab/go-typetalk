package v1

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nulab/go-typetalk/v3/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
)

type FilesService service

type AttachmentFile struct {
	ContentType string `json:"contentType"`
	FileKey     string `json:"fileKey"`
	FileName    string `json:"fileName"`
	FileSize    int    `json:"fileSize"`
}

// UploadAttachmentFile uploads attachment file.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/upload-attachment
func (s *FilesService) UploadAttachmentFile(ctx context.Context, topicID int, file *os.File) (*AttachmentFile, *shared.Response, error) {
	u := fmt.Sprintf("topics/%v/attachments", topicID)
	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	if stat.IsDir() {
		return nil, nil, errors.New("to upload can't be a directory")
	}
	form := map[string]io.Reader{
		"file": file,
	}
	req, err := s.client.NewMultipartRequest(u, form)
	if err != nil {
		return nil, nil, err
	}

	attachmentFile := &AttachmentFile{}
	resp, err := s.client.Do(ctx, req, attachmentFile)
	if err != nil {
		return nil, resp, err
	}
	return attachmentFile, resp, nil
}

// DownloadAttachmentFile downloads attachment file.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/download-attachment
func (s *FilesService) DownloadAttachmentFile(ctx context.Context, topicID, postID, attachmentID int, filename string) (io.ReadCloser, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/attachments/%d/%s", topicID, postID, attachmentID, filename)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", internal.DefaultMediaType)

	resp, err := s.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := internal.CheckResponse(resp); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return resp.Body, nil
}
