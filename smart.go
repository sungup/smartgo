package smartgo

/*
Scan all devices
*/

type DeviceType string

const (
	NVMe = DeviceType("nvme")
	SATA = DeviceType("sata")
)

type StorageDev interface {
	Type() DeviceType
	Device() string

	Close() error
}
