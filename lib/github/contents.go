package github

import (
	"context"
	"io/ioutil"
)

// DownloadContents downloads file contents from the given filepath
func (c *Client) DownloadContents(ctx context.Context, owner, repo, filepath string) ([]byte, error) {
	contents, err := c.restClient.Repositories.DownloadContents(ctx, owner, repo, filepath, nil)
	if err != nil {
		return nil, err
	}

	defer contents.Close()
	return ioutil.ReadAll(contents)
}
