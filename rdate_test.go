package minimumgo

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	d, err := getDate("192.168.10.104")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println("time.LoadLocation failed:", err)
		return
	}
	tLocal := d.In(loc)
	fmt.Printf("time=%v\n", tLocal)
}
