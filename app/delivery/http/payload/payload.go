package payload

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	once sync.Once
	// globalPayloadConfig is a map which hold list of payloadConfig for validation errors
	globalPayloadConfig config
)

func LoadPayloadConfigFile(filepath string) {
	once.Do(func() {
		data, err := os.ReadFile(filepath)
		if err != nil {
			panic(fmt.Errorf("read payload field config file: %w", err))
		}

		if err = yaml.Unmarshal(data, &globalPayloadConfig); err != nil {
			panic(fmt.Errorf("unmarshal payload field config file: %w", err))
		}
	})
}
