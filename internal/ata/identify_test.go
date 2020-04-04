package ata

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestSizeOfDevIdentify(t *testing.T) {
	a := assert.New(t)

	size := unsafe.Sizeof(DevIdentify{})

	a.Equal(512, int(size))

}
