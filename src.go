package bitbucket

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// NewFile holds basic new file info
type NewFile struct {
	GitPath   string
	LocalPath string
}

// AddCommit creates a new commit; adding, modifying, and deleting as appropriate
func (c *Client) AddCommit(
	addFiles []NewFile,
	deleteFiles []string,
	message string,
	author string,
	repositoryPath string,
) (interface{}, error) {

	urlStr := GetApiBaseURL() + "/repositories/" + repositoryPath + "/src/"

	// Data
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.WriteField("author", author)
	w.WriteField("message", message)
	for _, newFile := range addFiles {
		var fw io.Writer
		var err error
		r := mustOpen(newFile.LocalPath)

		if fw, err = w.CreateFormFile(newFile.GitPath, newFile.LocalPath); err != nil {
			panic(err)
		}
		if _, err = io.Copy(fw, r); err != nil {
			panic(err)
		}
	}
	for _, delFile := range deleteFiles {
		w.CreateFormField("files", delFile)
	}
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", urlStr, &b)
	if err != nil {
		return nil, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	c.authenticateRequest(req)
	return c.doRequest(req, true)
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
