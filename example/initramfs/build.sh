#!/bin/sh

GOSRC=../httpfileserver/httpfileserver.go
export GOARCH=arm
export GOARM=6

[ -d _root ] || mkdir _root 
go build -ldflags '-s -w' -o _root/init $GOSRC
fakeroot ./make_initramfs 
