package scsi

import (
	"github.com/sungup/smartgo/internal/ata"
	"math"
)

type scsiOPCode byte
type ataProtocol byte

const (
	// SCSI command (https://en.wikipedia.org/wiki/SCSI_command)
	scsiAtaPassThrough16 = scsiOPCode(0x85)

	// ATA command Pass-Through Revision 8 (p.5)
	ProtocolMASK  = byte(((0x01 << 4) - 1) << 1)
	ProtocolUMASK = ^ProtocolMASK

	HardReset   = byte(0x00 << 1) // exclude hard reset from protocol
	SRST        = byte(0x01 << 1) // exclude soft reset from protocol
	NonData     = ataProtocol(0x03 << 1)
	PIODataIn   = ataProtocol(0x04 << 1)
	PIODataOut  = ataProtocol(0x05 << 1)
	DMA         = ataProtocol(0x06 << 1)
	DMAQueued   = ataProtocol(0x07 << 1)
	DevDiag     = ataProtocol(0x08 << 1)
	DevReset    = ataProtocol(0x09 << 1)
	UDMADataIn  = ataProtocol(0x0a << 1)
	UDMADataOut = ataProtocol(0x0b << 1)
	FPDMA       = ataProtocol(0x0c << 1)
	ReturnResp  = ataProtocol(0x0f << 1)

	ExtendMASK  = byte(0x01)
	ExtendUMASK = ^ExtendMASK

	MultiplePos   = 5
	MultipleUMASK = ProtocolMASK | ExtendMASK

	TDirMASK  = byte(0x01 << 3)
	TDirUMASK = ^TDirMASK

	BytBlokMASK = byte(0x01 << 2)
	TLenMASK    = byte((0x01 << 2) - 1)
	TLenUMASK   = ^TLenMASK & ^BytBlokMASK

	OfflinePos = 6
)

// ATA Command Pass-Through Revision 8 (p.9)
// https://t10.org/ftp/t10/document.04/04-262r8.pdf
// ATA Extended Command CDB
// [0]:    OPERATION CODE(0x85)
// [1]:    MULTIPLE_COUNT(7..5) | PROTOCOL(4..1) | EXTEND(0)
// [2]:    OFF_LINE(7..6) | CK_COND(5) | T_DIR(3) | BYT_BLOCK (2) | T_LENGTH (1..0)
// [14:3]: ATA 48bit Command
// [15]:   CONTROL

type AtaCDB [16]byte

func (cdb *AtaCDB) setProtocol(protocol ataProtocol) {
	cdb[1] = (cdb[1] & ProtocolUMASK) | (byte(protocol) & ProtocolMASK)
}

func (cdb AtaCDB) getProtocol() ataProtocol {
	return ataProtocol(cdb[1] & ProtocolMASK)
}

func (cdb *AtaCDB) setExtend() {
	cdb[1] = cdb[1] | ExtendMASK
}

func (cdb *AtaCDB) unsetExtend() {
	cdb[1] = cdb[1] & ExtendUMASK
}

func (cdb AtaCDB) isExtend() bool {
	return (cdb[1] & ExtendMASK) == ExtendMASK
}

func (cdb *AtaCDB) setMultiple(sectors uint8) {
	Pof2 := uint8(math.Log2(float64(sectors))) << MultiplePos

	cdb[1] = (cdb[1] & MultipleUMASK) | Pof2
}

func (cdb *AtaCDB) resetMultiple() {
	cdb[1] = cdb[1] & MultipleUMASK
}

func (cdb AtaCDB) getMultiple() uint8 {
	return 0x01 << (cdb[1] >> MultiplePos)
}

func (cdb *AtaCDB) setHostDir() {
	cdb[2] = cdb[2] | TDirMASK
}

func (cdb *AtaCDB) setDevDir() {
	cdb[2] = cdb[2] & TDirUMASK
}

func (cdb *AtaCDB) resetDir() {
	cdb.setDevDir()
}

func (cdb AtaCDB) isDevDir() bool {
	return (cdb[2] & TDirMASK) == 0
}

func (cdb AtaCDB) isHostDir() bool {
	return (cdb[2] & TDirMASK) != 0
}

func (cdb *AtaCDB) setBlockSize(blocks uint8) {
	size := blocks&TLenMASK | BytBlokMASK

	cdb[2] = (cdb[2] & TLenUMASK) | size
}

func (cdb *AtaCDB) setByteSize(bytes uint8) {
	cdb[2] = (cdb[2] & TLenUMASK) | (bytes & TLenMASK)
}

func (cdb *AtaCDB) clearSize() {
	cdb[2] = cdb[2] & TLenUMASK
}

func (cdb AtaCDB) isBlockCmd() bool {
	return (cdb[2] & BytBlokMASK) == BytBlokMASK
}

func (cdb AtaCDB) getTLen() uint8 {
	return cdb[2] & TLenMASK
}

func (cdb *AtaCDB) setATACmd48bit(cmd ata.Cmd48bit) {
	copy(cdb[3:15], cmd[:])
}

func (cdb AtaCDB) getATACmd48bit() ata.Cmd48bit {
	cmd := ata.Cmd48bit{}

	copy(cmd[:], cdb[3:15])

	return cmd
}

func (cdb *AtaCDB) setControl(ctrl byte) {
	cdb[15] = ctrl
}

func (cdb AtaCDB) getControl() byte {
	return cdb[15]
}

func makeAtaCDB() AtaCDB {
	cdb := AtaCDB{byte(scsiAtaPassThrough16)}

	return cdb
}

func __makeReset(protocol, second uint8) AtaCDB {
	cdb := makeAtaCDB()

	offline := uint8(math.Log2(float64(second+2))) - 1

	cdb[1] = (cdb[1] & ProtocolUMASK) | (protocol & ProtocolMASK)
	cdb[2] = offline << OfflinePos

	return cdb
}

func makeHardReset(second uint8) AtaCDB {
	return __makeReset(HardReset, second)
}

func makeSoftReset(second uint8) AtaCDB {
	return __makeReset(SRST, second)
}

func makeIdentify() AtaCDB {
	cmd := ata.Cmd48bit{}
	cmd.SetCommand(ata.IdentifyCmd)

	cdb := makeAtaCDB()
	cdb.setProtocol(PIODataIn)
	cdb.setHostDir()
	cdb.setBlockSize(2)
	cdb.setATACmd48bit(cmd)

	return cdb
}
