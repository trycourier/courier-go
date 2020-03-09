package courier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SendResponse struct {
	MessageId string
}

func (s *Client) Send(message []byte) (*SendResponse, error) {
	url := fmt.Sprintf(s.BaseUrl + "/send")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data SendResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
