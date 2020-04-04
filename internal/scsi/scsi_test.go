package scsi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDevice_GetIdentify(t *testing.T) {
	a := assert.New(t)

	dev, err := New("/dev/sda")

	a.NoError(err)

	tested, err := dev.getIdentify()
	a.NoError(err)

	t.Log(tested.Model.String())
}
