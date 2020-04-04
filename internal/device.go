package internal

import (
	"errors"
	"io/ioutil"
	"path"
	"regexp"
	"runtime"

	"github.com/sungup/smartgo"
)

const (
	deviceRoot = "/dev"
)

var (
	linuxNvmeMatch, _  = regexp.Compile("^nvme([0-9]*$)")
	linuxSataMatch, _  = regexp.Compile("^sd([a-z]$)")
	darwinSataMatch, _ = regexp.Compile("^disk([0-9]*$)")
)

func GetDevFiles(devType smartgo.DeviceType) ([]string, error) {
	stats, err := ioutil.ReadDir(deviceRoot)
	if err != nil {
		return nil, err
	}

	regex := linuxSataMatch

	if runtime.GOOS == "darwin" {
		if devType != smartgo.SATA {
			return nil, errors.New("unsupported device type on darwin")
		}

		regex = darwinSataMatch

	} else if runtime.GOOS == "linux" {
		if devType == smartgo.NVMe {
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
