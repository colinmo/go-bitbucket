package bitbucket

import (
	"encoding/json"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/mitchellh/mapstructure"
)

type NewFile struct {
	GitPath string
	LocalPath string
	Content []byte
}

// AddCommit creates a new commit; adding, modifying, and deleting as appropriate
func (c *Client) AddCommit(
	addFiles []NewFile,
	deleteFiles []string,
	message string,
	author string,
) (interface{}, error) {

	urlStr := GetApiBaseURL() + "/src/"

	// Data
	body := map[string]interface{}{}
	body["message"] = "Posting " + aPost.GetType() + " " + aPost.GetTitle()
	body["author"] = "Colin Morris <relapse@gmail.com>"

	for _, newFile in range addFiles {
		content, err := ioutil.ReadFile(newFile.LocalPath)
		if !err {
			body[sourcecontrol.PostsBase+newFile.GitPath] = content
		}
	}
	if len(deleteFiles) > 0 {
		body["files"] = deleteFiles
	}

	return c.execute("GET", urlStr, body)
}
