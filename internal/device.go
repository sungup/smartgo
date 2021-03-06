package internal

import (
	"errors"
	"io/ioutil"
	"path"
	"regexp"
	"runtime"
)

type DeviceType string

const (
	NVMe = DeviceType("nvme")
	SATA = DeviceType("sata")

	deviceRoot = "/dev"
)

var (
	linuxNvmeMatch, _  = regexp.Compile("^nvme([0-9]*$)")
	linuxSataMatch, _  = regexp.Compile("^sd([a-z]$)")
	darwinSataMatch, _ = regexp.Compile("^disk([0-9]*$)")
)

type StorageDevice interface {
	Type() DeviceType
	Device() string
	Model() string
	Firmware() string
	Serial() string

	ScanSMART() error
}

type StorageMeta struct {
	StorageDevice

	devType  DeviceType
	devPath  string
	model    string
	firmware string
	serial   string
}

func (meta *StorageMeta) Type() DeviceType {
	return meta.devType
}

func (meta *StorageMeta) Device() string {
	return meta.devPath
}

func GetDevFiles(devType DeviceType) ([]string, error) {
	stats, err := ioutil.ReadDir(deviceRoot)
	if err != nil {
		return nil, err
	}

	regex := linuxSataMatch

	if runtime.GOOS == "darwin" {
		if devType != SATA {
			return nil, errors.New("unsupported device type on darwin")
		}

		regex = darwinSataMatch

	} else if runtime.GOOS == "linux" {
		if devType == NVMe {
			regex = linuxNvmeMatch
		} else {
			regex = linuxSataMatch
		}

	}

	files := make([]string, 0)

	for _, stat := range stats {
		if stat.IsDir() || !regex.MatchString(stat.Name()) {
			continue
		}

		files = append(files, path.Join(deviceRoot, stat.Name()))
	}

	return files, nil
}
