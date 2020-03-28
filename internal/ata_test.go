package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanSATA(t *testing.T) {
	/*
		a := assert.New(t)

		storage, err := ScanSATA(make(map[string]StorageDevice))

		a.NoError(err)

		for _, dev := range storage {
			t.Log(dev.Device())
			t.Log(dev.Type())
		}
	*/
}

func TestAtaCDB(t *testing.T) {
	a := assert.New(t)

	cdb := makeAtaCDB()

	// Protocol Test
	cdb.setExtendBit()
	cdb.setMultiple(16)
	cdb.setProtocol(PIODataIn)
	a.Equal(PIODataIn, cdb.getProtocol())
	cdb.setProtocol(PIODataOut)
	a.Equal(PIODataOut, cdb.getProtocol())
	cdb.unsetExtendBit()
	cdb.resetMultiple()

	// ExtendBit Test
	a.False(cdb.isExtendSet())
	cdb.setExtendBit()
	a.True(cdb.isExtendSet())
	cdb.unsetExtendBit()
	a.False(cdb.isExtendSet())

	// Multiple field
	a.Equal(uint8(1), cdb.getMultiple())
	cdb.setMultiple(16)
	a.Equal(uint8(4), cdb[1]>>MultiplePos)
	a.Equal(uint8(16), cdb.getMultiple())
	cdb.resetMultiple()
	a.Equal(uint8(1), cdb.getMultiple())

	// Device Direction
	a.True(cdb.isDevDir())
	cdb.devToHostDir()
	a.False(cdb.isDevDir())
	cdb.hostToDevDir()
	a.True(cdb.isDevDir())
}

func TestAtaCDBReset(t *testing.T) {
	a := assert.New(t)

	hardReset := makeHardReset(6)
	a.Equal(ataProtocol(HardReset), hardReset.getProtocol())
	a.Equal(uint8(2), hardReset[2]>>OfflinePos)

	softReset := makeSoftReset(2)
	a.Equal(ataProtocol(SRST), softReset.getProtocol())
	a.Equal(uint8(1), softReset[2]>>OfflinePos)
}
