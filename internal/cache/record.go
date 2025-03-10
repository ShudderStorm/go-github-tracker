package cache

import "encoding/json"

type Access struct {
	ID     uint     `json:"id"`
	Token  string   `json:"token"`
	Type   string   `json:"type"`
	Scopes []string `json:"scopes"`
}

func SerializeAccess(access Access) ([]byte, error) {
	return json.Marshal(access)
}

func DeserializeAccess(data []byte) (Access, error) {
	var access Access
	err := json.Unmarshal(data, &access)
	return access, err
}

type State string

func SerializeState(state State) ([]byte, error) {
	return []byte(state), nil
}

func DeserializeState(data []byte) (State, error) {
	return State(data), nil
}
