package minimumgo

import (
	"log"

	"github.com/beevik/ntp"
	"golang.org/x/sys/unix"
)

func Ntpdate(server string) (err error) {
	d, err := ntp.Time(server)
	if err != nil {
		log.Printf("Ntpdate: %v\n", err)
		return err
	}
	tv := unix.NsecToTimeval(d.UnixNano())
	return unix.Settimeofday(&tv)
}
