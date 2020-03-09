package courier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Profile struct {
	Profile interface{} `json:"profile"`
}

func (s *Client) GetProfile(id string) (*Profile, error) {
	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", id)
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

func (s *Client) MergeProfile(id string, profile []byte) error {
	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", id)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) UpdateProfile(id string, profile []byte) error {
	url := fmt.Sprintf(s.BaseUrl+"/profiles/%s", id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}
