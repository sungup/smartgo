package scsi

import (
	"github.com/stretchr/testify/assert"
	"github.com/sungup/smartgo/internal/ata"
	"testing"
)

const (
	defaultProtocol = DMA
	defaultMulti    = uint8(8)
	defaultTLen     = uint8(3)

	testHardOffline        = uint8(6)
	testSoftOffline        = uint8(2)
	expectedHardOfflineVal = uint8(2)
	expectedSoftOfflineVal = uint8(1)

	expectedControl = byte(0xa)
)

func makeTestScsiATAPassThroughCDB(a *assert.Assertions) AtaCDB {
	cdb := makeAtaCDB()

	cdb.setProtocol(defaultProtocol)
	cdb.setExtend()
	cdb.setMultiple(defaultMulti)
	cdb.setHostDir()
	cdb.setBlockSize(defaultTLen)

	// baseline test
	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())

	return cdb
}

func TestAtaCDB_Protocol(t *testing.T) {
	a := assert.New(t)

	cdb := makeTestScsiATAPassThroughCDB(a)

	cdb.setProtocol(PIODataIn)

	a.Equal(PIODataIn, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())
}

func TestAtaCDB_Extend(t *testing.T) {
	a := assert.New(t)

	cdb := makeTestScsiATAPassThroughCDB(a)

	cdb.unsetExtend()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.False(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())

	cdb.setExtend()
	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())
}

func TestAtaCDB_Multiple(t *testing.T) {
	a := assert.New(t)

	cdb := makeTestScsiATAPassThroughCDB(a)

	cdb.resetMultiple()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(uint8(1), cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())

	cdb.setMultiple(128)

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(uint8(128), cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())
}

func TestAtaCDB_TDir(t *testing.T) {
	a := assert.New(t)

	cdb := makeTestScsiATAPassThroughCDB(a)

	cdb.resetDir()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.False(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())

	cdb.setHostDir()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())

	cdb.setDevDir()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.False(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(defaultTLen, cdb.getTLen())
}

func TestAtaCDB_TLen(t *testing.T) {
	a := assert.New(t)

	cdb := makeTestScsiATAPassThroughCDB(a)

	cdb.clearSize()

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.False(cdb.isBlockCmd())
	a.Equal(uint8(0), cdb.getTLen())

	cdb.setBlockSize(1)

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.True(cdb.isBlockCmd())
	a.Equal(uint8(1), cdb.getTLen())

	cdb.setByteSize(2)

	a.Equal(defaultProtocol, cdb.getProtocol())
	a.True(cdb.isExtend())
	a.Equal(defaultMulti, cdb.getMultiple())
	a.True(cdb.isHostDir())
	a.False(cdb.isBlockCmd())
	a.Equal(uint8(2), cdb.getTLen())
}

func TestAtaCDB_Reset(t *testing.T) {
	a := assert.New(t)

	hardReset := makeHardReset(testHardOffline)
	a.Equal(ataProtocol(HardReset), hardReset.getProtocol())
	a.Equal(expectedHardOfflineVal, hardReset[2]>>OfflinePos)

	softReset := makeSoftReset(testSoftOffline)
	a.Equal(ataProtocol(SRST), softReset.getProtocol())
	a.Equal(expectedSoftOfflineVal, softReset[2]>>OfflinePos)
}

func TestAtaCDB_ATACmd48bit(t *testing.T) {
	a := assert.New(t)

	testCDB := makeTestScsiATAPassThroughCDB(a)
	testCDB.setControl(expectedControl)

	cleanCDB := testCDB

	expectedCmd := ata.Cmd48bit{}
	cleanCmd := ata.Cmd48bit{}
	for i := 0; i < len(expectedCmd); i++ {
		expectedCmd[i] = byte(i)
	}
	a.NotEqual(expectedCmd, cleanCmd)

	// Set/Get test
	testCDB.setATACmd48bit(expectedCmd)
	a.Equal(expectedCmd, testCDB.getATACmd48bit())
	a.NotEqual(testCDB, cleanCDB)
	a.Equal(expectedControl, testCDB.getControl())

	// CDB layout test
	testCDB.setATACmd48bit(cleanCmd)
	a.Equal(cleanCmd, testCDB.getATACmd48bit())
	a.Equal(testCDB, cleanCDB)
	a.Equal(expectedControl, testCDB.getControl())
}

func TestAtaCDB_MakeIdentify(t *testing.T) {
	a := assert.New(t)

	cdb := makeIdentify()

	a.Equal(byte(0x08), cdb[1])
	a.Equal(byte(0x0e), cdb[2])
	a.Equal(byte(ata.IdentifyCmd), cdb[14])
}
