package ata

type Command byte

const (
	IdentifyCmd = Command(0xEC)
)

// T13/1699-D Revision 4a (p.109)
// Feature:            00h     [1:0]
// Count:              01h     [3:2]
// LBA:                02h-04h [9:4]
// Device and Command: 05h     [11:10] 10 is device, 11 is command
type Cmd48bit [12]byte

// IDENTIFY DEVICE - 0xEC, PIO Data-In
//   FEATURE: N/A
//   COUNT:   N/A
//   LBA:     N/A
//   DEVICE:  [4] Transport Dependent
//   COMMAND: [7:0] 0xEC

func (ataCmd *Cmd48bit) SetCommand(cmd Command) {
	ataCmd[11] = byte(cmd)
}
