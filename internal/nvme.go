package internal

type NVMeDevice struct {
	StorageMeta
}

func newNVMeDev(path string) *NVMeDevice {
	nvme := new(NVMeDevice)

	nvme.devType = NVMe
	nvme.devPath = path

	return nvme
}

func ScanNVMe(storage map[string]StorageDevice) (map[string]StorageDevice, error) {

	return storage, nil
}

/*
 * inherited interface methods
 */
