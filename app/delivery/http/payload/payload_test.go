package payload

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	LoadPayloadConfigFile("../../../../assets/payload_field_config.yaml")

	os.Exit(m.Run())
}
