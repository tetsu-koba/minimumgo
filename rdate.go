package minimumgo

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/sys/unix"
)

const (
	timeservice_port = 37
	timeout          = 5 * time.Second
)

func Rdate(server string) (err error) {
	d, err := getDate(server)
	if err != nil {
		return err
	}
	tv := unix.NsecToTimeval(d.UnixNano())
	err = unix.Settimeofday(&tv)
	return err
}

func getDate(server string) (tUtc time.Time, err error) {
	addr := fmt.Sprintf("%s:%d", server, timeservice_port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		log.Println("rdate:", err)
		return
	}
	defer conn.Close()

	var t uint32
	err = binary.Read(conn, binary.BigEndian, &t)
	if err != nil {
		log.Println("binary.Read failed:", err)
		return
	}
	tUtc = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(t) * time.Second)
	return
}
