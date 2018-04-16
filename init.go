package minimumgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func init() {
	err := mountfs()
	if err != nil {
		log.Println(err)
		return
	}
	err = setupNetwork()
	if err != nil {
		log.Println(err)
		return
	}

	if ntpServer := getKernelcmdline("minimumgo.ntp"); ntpServer != "" {
		Ntpdate(ntpServer)
	} else if rdateServer := getKernelcmdline("minimumgo.rdate"); rdateServer != "" {
		Rdate(rdateServer)
	}
}

func setupNetwork() error {
	hostnameb, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		log.Printf("Failed to read /etc/hostname. Ignore\n")
	}
	if err := unix.Sethostname(hostnameb); err != nil {
		log.Printf("Failed to set hostname. Ignore\n")
	}
	if !Exists("/etc/resolv.conf") && getKernelcmdline("ip") != "" {
		if err := os.Symlink("/proc/net/pnp", "/etc/resolv.conf"); err != nil {
			log.Printf("Failed to set /proc/resolv.conf\n")
		}
	}

	if !Exists("/etc/hosts") {
		hostsb := []byte("127.0.0.1 localhost\n::1 localhost\n")
		if hostnameb != nil {
			hostsb = append(hostsb, []byte("127.0.1.1 ")...)
			hostsb = append(hostsb, hostnameb...)
			hostsb = append(hostsb, '\n')
		}
		if err := ioutil.WriteFile("/etc/hosts", hostsb, 0644); err != nil {
			return fmt.Errorf("/etc/hosts: %v", err)
		}
	}
	return nil
}

func getKernelcmdline(key string) string {
	b, err := ioutil.ReadFile("/proc/cmdline")
	if err != nil {
		return ""
	}
	for _, v := range bytes.Split(b, []byte(" ")) {
		v2 := bytes.SplitN(v, []byte("="), 2)
		if string(v2[0]) == key {
			return strings.TrimSpace(string(v2[1]))
		}
	}
	return ""
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
