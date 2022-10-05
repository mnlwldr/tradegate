package tradegate

import (
	"testing"
)

func TestGet(t *testing.T) {
	_, err := Get("DE000A1YCMM2") // SolarWorld AG
	if err != nil {
		t.Error(err)
	}
}
