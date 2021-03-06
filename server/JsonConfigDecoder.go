package server

import (
	"encoding/json"
	"github.com/coveo/uabot-server/explorerlib"
	"io"
)

func DecodeConfig(jsonReader io.Reader) (*explorerlib.Config, error) {
	config := &explorerlib.Config{}
	err := json.NewDecoder(jsonReader).Decode(config)
	return config, err
}
