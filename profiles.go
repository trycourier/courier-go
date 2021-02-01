package courier

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Profile struct {
	Profile interface{} `json:"profile"`
}

type PatchOp struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type ProfilePatchRequest struct {
	Patch []PatchOp `json:"patch"`
}

type ProfileSubscriptionResponse struct {
	Paging  *PagingResponse
	Results []*ListResponse
}

func (s *Client) GetProfile(recipientID string) (*Profile, error) {
	if recipientID == "" {
		return nil, errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", recipientID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Profile
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) MergeProfile(recipientID string, profile []byte) error {
	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", recipientID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) UpdateProfile(recipientID string, profile []byte) error {
	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", recipientID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) PatchProfile(recipientID string, patch []PatchOp) error {
	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	if patch == nil {
		return errors.New("Patch cannot be nil")
	}

	payload := ProfilePatchRequest{
		Patch: patch,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", recipientID)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) GetProfileLists(recipientID string, cursor string) (*ProfileSubscriptionResponse, error) {
	if recipientID == "" {
		return nil, errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s/lists", recipientID)
	if cursor != "" {
		url = url + "?cursor=" + cursor
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ProfileSubscriptionResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
