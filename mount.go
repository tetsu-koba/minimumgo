package minimumgo

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func mountfs() error {
	// ignore error
	_ = os.Mkdir("/tmp", 1777)
	_ = os.Mkdir("/etc", 0755)
	_ = os.Mkdir("/dev", 0755)
	_ = os.Mkdir("/proc", 0555)
	_ = os.Mkdir("/sys", 0555)
	_ = os.Mkdir("/boot", 0755)

	if err := unix.Mount("tmpfs", "/tmp", "tmpfs", unix.MS_RELATIME, "size=50M"); err != nil {
		return fmt.Errorf("tmpfs on /tmp: %v", err)
	}

	if err := unix.Mount("devtmpfs", "/dev", "devtmpfs", 0, ""); err != nil {
		if err != unix.EBUSY {
			return fmt.Errorf("devtmpfs: %v", err)
		}
	}

	_ = os.Mkdir("/dev/pts", 0755)
	if err := unix.Mount("devpts", "/dev/pts", "devpts", 0, ""); err != nil {
		return fmt.Errorf("devpts: %v", err)
	}

	if err := unix.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		if err != unix.EBUSY {
			return fmt.Errorf("proc: %v", err)
		}
	}

	if err := unix.Mount("sysfs", "/sys", "sysfs", 0, ""); err != nil {
		if err != unix.EBUSY {
			return fmt.Errorf("sys: %v", err)
		}
	}

	if err := unix.Mount("/dev/mmcblk0p1", "/boot", "vfat", unix.MS_RDONLY, ""); err != nil {
		log.Printf("Could not mount boot partition: %v", err)
	}

	return nil
}
