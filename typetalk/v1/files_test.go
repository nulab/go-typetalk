package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/nulab/go-typetalk/typetalk/internal"
)

func Test_FilesService_UploadAttachmentFile_should_upload_an_attachment_file(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "upload-attachment-file.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%v/attachments", topicId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, string(b))
	})

	f, _ := os.Open(fixturesPath + "sample.jpg")
	defer f.Close()
	result, _, err := client.Files.UploadAttachmentFile(context.Background(), topicId, f)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &AttachmentFile{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_FilesService_DownloadAttachmentFile_should_download_an_attachment_file(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	attachmentId := 1
	filename := "sample.jpg"
	b, _ := ioutil.ReadFile(fixturesPath + filename)
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d/attachments/%d/%s", topicId, postId, attachmentId, filename), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", internal.DefaultMediaType)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		fmt.Fprint(w, string(b))
	})

	reader, err := client.Files.DownloadAttachmentFile(context.Background(), topicId, postId, attachmentId, filename)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Returned bad reader: %v", err)
	}
	if !bytes.Equal(b, content) {
		t.Errorf("Returned %+v, want %+v", content, b)
	}
}
