package scsi

import (
	"bytes"
	"encoding/binary"
	"github.com/sungup/smartgo"
	"github.com/sungup/smartgo/internal"
	"github.com/sungup/smartgo/internal/ata"
	"golang.org/x/sys/unix"
	"unsafe"
)

type Device struct {
	smartgo.StorageDev

	devPath string
	fd      int
}

/*
internal
*/
func (dev *Device) sendCDB(cdb AtaCDB, resp *[]byte) error {
	sense := make([]byte, sgCDBSenseLen)

	header := sgIoHdr{
		interfaceId:    sgInterfaceIdOrig,
		dxferDirection: sgDeferFromDev,
		cmdLen:         uint8(len(cdb)),
		mxSbLen:        sgCDBSenseLen,

		dxferLen: uint32(len(*resp)),
		dxferp:   uintptr(unsafe.Pointer(&(*resp)[0])),

		cmdp:    uintptr(unsafe.Pointer(&cdb[0])),
		sbp:     uintptr(unsafe.Pointer(&sense[0])),
		timeout: sgDefaultTimeout,
	}

	if err := internal.IoCtl(uintptr(dev.fd), sgIo, header.toPtr()); err != nil {
		return err
	}

	if !header.isInfoOk() {
		return newSgIoError(&header, sense)
	}

	return nil
}

func (dev *Device) getIdentify() (ata.DevIdentify, error) {
	var identify ata.DevIdentify

	buffer := make([]byte, unsafe.Sizeof(identify))

	if err := dev.sendCDB(makeIdentify(), &buffer); err != nil {
		return identify, err
	}

	if err := binary.Read(bytes.NewReader(buffer), binary.BigEndian, &identify); err != nil {
		return identify, err
	}

	return identify, nil
}

/*
interfaces
*/
func (dev Device) Type() smartgo.DeviceType {
	// TODO need to change scsi and sata device
	return smartgo.SATA
}

func (dev Device) Device() string {
	return dev.devPath
}

func (dev *Device) Close() error {
	return unix.Close(dev.fd)
}

func New(devPath string) (*Device, error) {
	var err error

	dev := new(Device)

	dev.devPath = devPath
	dev.fd, err = unix.Open(dev.devPath, unix.O_RDWR, 0600)

	return dev, err
}
