#!/bin/sh

COPY_LOCALTIME=1
#EXTRACT_ZONEINFO=1
COPY_CA_CERTIFICATE=1

ROOT=_root
GOSRC=../httpfileserver/httpfileserver.go
export GOARCH=arm
export GOARM=6

[ -d $ROOT ] || mkdir $ROOT

if [ "$COPY_LOCALTIME" ]; then
    [ -d $ROOT/etc ] || mkdir $ROOT/etc
    cp /etc/localtime $ROOT/etc/
fi

if [ "$EXTRACT_ZONEINFO" ]; then
    [ -r zoneinfo.zip ] || wget https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip
    [ -d $ROOT/usr/share/zoneinfo ] || mkdir -p $ROOT/usr/share/zoneinfo
    unzip zoneinfo.zip -d $ROOT/usr/share/zoneinfo/
fi

if [ "$COPY_CA_CERTIFICATE" ]; then
    [ -d $ROOT/etc/ssl/certs ] || mkdir -p $ROOT/etc/ssl/certs
    cp  /etc/ssl/certs/ca-certificates.crt $ROOT/etc/ssl/certs/
fi

go build -ldflags '-s -w' -o $ROOT/init $GOSRC
fakeroot ./make_initramfs 
