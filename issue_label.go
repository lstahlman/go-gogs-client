// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)

type IssueLabelsOption struct {
	Labels []int64	`json:"labels"`
}

func (c *Client) GetIssueLabels(owner, repo string, index int64) ([]*Label, error) {
	labels := make([]*Label, 0)
	return labels, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/issues/%d/labels", owner, repo, index), nil, nil, &labels)
}

func (c *Client) AddIssueLabels(owner, repo string, index int64, opt IssueLabelsOption) ([]*Label, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	labels := make([]*Label, 0)
	return labels, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/issues/%d/labels", owner, repo, index),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), &labels)
}

func (c *Client) ReplaceIssueLabels(owner, repo string, index int64, opt IssueLabelsOption) ([]*Label, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	labels := make([]*Label, 0)
	return labels, c.getParsedResponse("PUT", fmt.Sprintf("/repos/%s/%s/issues/%d/labels", owner, repo, index),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), &labels)
}

func (c *Client) DeleteIssueLabel(owner, repo string, index int64, label int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/issues/%d/labels/%d", owner, repo, index, label), nil, nil)
	return err
}