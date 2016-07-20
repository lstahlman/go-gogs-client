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

type Label struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type LabelOption struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (c *Client) ListRepoLabels(owner, repo string) ([]*Label, error) {
	labels := make([]*Label, 0)
	return labels, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/labels", owner, repo), nil, nil, &labels)
}

func (c *Client) GetRepoLabel(owner, repo string, index int64) (*Label, error) {
	label := new(Label)
	return label, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/labels/%d", owner, repo, index), nil, nil, label)
}

func (c *Client) CreateLabel(owner, repo string, opt LabelOption) (*Label, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	label := new(Label)
	return label, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/labels", owner, repo),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), label)
}


func (c *Client) EditLabel(owner, repo string, index int64, opt LabelOption) (*Label, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	label := new(Label)
	return label, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/labels/%d", owner, repo, index),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), label)
}

func (c *Client) DeleteLabel(owner, repo string, index int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/labels/%d", owner, repo, index), nil, nil)
	return err
}