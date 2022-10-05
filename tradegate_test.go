package tradegate

import (
	"testing"
)

func TestGet(t *testing.T) {
	got, err := Get("DE000A1YCMM2") // SolarWorld AG
	if err != nil {
		t.Error(err)
	}
}
