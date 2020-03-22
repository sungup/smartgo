package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIoCtl(t *testing.T) {
	/*
	 */
	// Not working on MacOSX
	a := assert.New(t)

	// 1. open device file
	testDevice := "/dev/disk0" // default MacBook device file
	//_, err := os.Open(testDevice)
	//_, err := unix.Open(testDevice, os.O_RDWR, 0640)
	//a.NoError(err)

	stat, err := os.Stat(testDevice)
	a.NoError(err)
	t.Log(stat.Mode())
}
