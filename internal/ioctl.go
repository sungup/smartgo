package internal

import (
	"golang.org/x/sys/unix"
)

func IoCtl(fd, cmd, ptr uintptr) error {
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, cmd, ptr); err != 0 {
		return err
	}

	return nil
}
