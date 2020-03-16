package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanSATA(t *testing.T) {
	a := assert.New(t)

	storage, err := ScanSATA(make(map[string]StorageDevice))

	a.NoError(err)

	for _, dev := range storage {
		t.Log(dev.Device())
		t.Log(dev.Type())
	}
}
