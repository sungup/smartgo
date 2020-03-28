package internal

import "math"

type ataProtocol uint8

const (
	AtaIdentifyDev = 0xEC

	// SCSI command (https://en.wikipedia.org/wiki/SCSI_command)
	scsiAtaPassThrough16 = 0x85

	// ATA command Pass-Through Revision 8 (p.5)
	ProtocolMASK  = uint8(((0x01 << 4) - 1) << 1)
	ProtocolUMASK = ^ProtocolMASK

	HardReset   = uint8(0x00 << 1) // exclude hard reset from protocol
	SRST        = uint8(0x01 << 1) // exclude soft reset from protocol
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

	ExtendMASK  = uint8(0x01)
	ExtendUMASK = ^ExtendMASK

	MultiplePos   = 5
	MultipleUMASK = ProtocolMASK | ExtendMASK

	TDirMASK  = uint8(0x01 << 3)
	TDirUMASK = ^TDirMASK

	BytBlokMASK = uint8(0x01 << 2)
	TLenMASK    = uint8((0x01 << 2) - 1)
	TLenUMASK   = ^TLenMASK | ^BytBlokMASK

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

type ataCDB [16]byte

func (cdb *ataCDB) setProtocol(protocol ataProtocol) {
	cdb[1] = (cdb[1] & ProtocolUMASK) | (uint8(protocol) & ProtocolMASK)
}

func (cdb ataCDB) getProtocol() ataProtocol {
	return ataProtocol(cdb[1] & ProtocolMASK)
}

func (cdb *ataCDB) setExtendBit() {
	cdb[1] = cdb[1] | ExtendMASK
}

func (cdb *ataCDB) unsetExtendBit() {
	cdb[1] = cdb[1] & ExtendUMASK
}

func (cdb ataCDB) isExtendSet() bool {
	return (cdb[1] & ExtendMASK) == ExtendMASK
}

func (cdb *ataCDB) setMultiple(sectors uint8) {
	Pof2 := uint8(math.Log2(float64(sectors))) << MultiplePos

	cdb[1] = (cdb[1] & MultipleUMASK) | Pof2
}

func (cdb *ataCDB) resetMultiple() {
	cdb[1] = cdb[1] & MultipleUMASK
}

func (cdb ataCDB) getMultiple() uint8 {
	return 0x01 << (cdb[1] >> MultiplePos)
}

func (cdb *ataCDB) devToHostDir() {
	cdb[2] = cdb[2] | TDirMASK
}

func (cdb *ataCDB) hostToDevDir() {
	cdb[2] = cdb[2] & TDirUMASK
}

func (cdb *ataCDB) resetDir() {
	cdb.hostToDevDir()
}

func (cdb ataCDB) isDevDir() bool {
	return (cdb[2] & TDirMASK) == 0
}

func (cdb ataCDB) isHostDir() bool {
	return (cdb[2] & TDirMASK) != 0
}

func (cdb *ataCDB) setBlockSize(blocks uint8) {
	size := blocks&TLenMASK | BytBlokMASK

	cdb[2] = (cdb[2] & TLenUMASK) | size
}

func (cdb *ataCDB) setByteSize(bytes uint8) {
	cdb[2] = (cdb[2] & TLenUMASK) | (bytes & TLenMASK)
}

func (cdb *ataCDB) clearSize() {
	cdb[2] = cdb[2] & TLenUMASK
}

func (cdb ataCDB) isBlockCmd() bool {
	return (cdb[2] & BytBlokMASK) == BytBlokMASK
}

func (cdb ataCDB) getTLen() uint8 {
	return cdb[2] & TLenMASK
}

func makeAtaCDB() ataCDB {
	cdb := ataCDB{scsiAtaPassThrough16}

	return cdb
}

func __makeReset(protocol, second uint8) ataCDB {
	cdb := makeAtaCDB()

	offline := uint8(math.Log2(float64(second+2))) - 1

	cdb[1] = (cdb[1] & ProtocolUMASK) | (protocol & ProtocolMASK)
	cdb[2] = offline << OfflinePos

	return cdb
}

func makeHardReset(second uint8) ataCDB {
	return __makeReset(HardReset, second)
}

func makeSoftReset(second uint8) ataCDB {
	return __makeReset(SRST, second)
}

// T13/1699-D Revision 4a (p.109)
// Feature:            00h     [1:0]
// Count:              01h     [3:2]
// LBA:                02h-04h [9:4]
// Device and Command: 05h     [11:10]
type ata48BitCmd struct {
	feature word
	count   word
	lba     [3]word
	device  byte
	command byte
}

// IDENTIFY DEVICE - 0xEC, PIO Data-In
//   FEATURE: N/A
//   COUNT:   N/A
//   LBA:     N/A
//   DEVICE:  [4] Transport Dependent
//   COMMAND: [7:0] 0xEC

type SATADevice struct {
	StorageMeta
}

func newSATADev(path string) *SATADevice {
	sata := new(SATADevice)

	sata.devType = SATA
	sata.devPath = path

	return sata
}

func ScanSATA(storage map[string]StorageDevice) (map[string]StorageDevice, error) {
	files, err := GetDevFiles(SATA)
	if err != nil {
		return storage, err
	}

	for _, file := range files {
		sata := newSATADev(file)

		/*
			if err := sata.ScanSMART(); err != nil {
				// this device cannot support S.M.A.R.T
				continue
			}

			if _, ok := storage[sata.Serial()]; ok {
				continue
			}

			storage[sata.Serial()] = sata
		*/
		storage[sata.Device()] = sata
	}

	return storage, nil
}
