package courier

import "encoding/json"

func toJSONMap(input interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	temp, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(temp, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
