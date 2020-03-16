package smartgo

import (
	"github.com/sungup/smartgo/internal"
)

/*
Scan all devices
*/
func ScanDevice() (map[string]internal.StorageDevice, error) {
	storage := make([]internal.StorageDevice, 0)
	var err error

	if storage, err = internal.ScanNVMe(storage); err != nil {
		return nil, err
	}

	if storage, err = internal.ScanSATA(storage); err != nil {
		return nil, err
	}

	for _, device := range storage {
		if err := device.ScanSMART(); err != nil {
			return nil, err
		}
	}

	return storage, nil
}
