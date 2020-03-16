package internal

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
