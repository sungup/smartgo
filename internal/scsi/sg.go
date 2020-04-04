package scsi

import (
	"fmt"
	"unsafe"
)

const (
	sgInterfaceIdOrig = 'S'

	sgDeferNone      = -1
	sgDeferToDev     = -2
	sgDeferFromDev   = -3
	sgDeferToFromDev = -4

	sgDefaultTimeout = 1000 // 1sec

	sgCDBSenseLen = 32

	// synchronous SCSI command IOCTL, (only in version 3 interface)
	sgIo = 0x2285

	// reference: https://www.tldp.org/HOWTO/SCSI-Generic-HOWTO/x364.html
	sgInfoOkMask = 0x1
	sgInfoOk     = 0x0
	sgInfoCheck  = 0x1
)

type sgIoHdr struct {
	interfaceId    int32
	dxferDirection int32
	cmdLen         uint8
	mxSbLen        uint8
	iovecCount     uint16
	dxferLen       uint32
	dxferp         uintptr

	cmdp         uintptr
	sbp          uintptr
	timeout      uint32
	flags        uint32
	packId       int32
	usrPtr       uintptr
	status       uint8
	maskedStatus uint8
	msgStatus    uint8
	sbLenWr      uint8
	hostStatus   uint16
	driverStatus uint16
	resid        int32
	duration     uint32
	info         uint32
}

func (header sgIoHdr) isInfoOk() bool {
	return header.info&sgInfoOkMask == sgInfoOk
}

func (header *sgIoHdr) toPtr() uintptr {
	return uintptr(unsafe.Pointer(header))
}

type sgIoError struct {
	scsi   uint8  // scsi status
	host   uint16 // host status
	driver uint16 // driver status

	sense []byte
}

func (err sgIoError) Error() string {
	return fmt.Sprintf(
		"error status (scsi: %#02x, host: %#02x, driver: %#02x",
		err.scsi,
		err.host,
		err.driver,
	)
}

func newSgIoError(header *sgIoHdr, sense []byte) sgIoError {
	err := sgIoError{
		scsi:   header.status,
		host:   header.hostStatus,
		driver: header.driverStatus,
		sense:  sense,
	}

	return err
}
