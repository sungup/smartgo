package internal

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
	"os"
	"testing"
)

func TestIoCtl(t *testing.T) {
	a := assert.New(t)

	// 1. open device file
	testDevice := "/dev/disk0" // default MacBook device file
	//_, err := os.Open(testDevice)
	_, err := unix.Open(testDevice, os.O_RDWR, 0640)
	a.NoError(err)

}
