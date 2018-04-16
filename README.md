# README

minimumgo: Do minimum initialization to execute a single golang executable on linux kernal

__STATUS: This is still alpha version__

If you have a golang executalbe file and you don't need any other, import this package and execute your golang executalbe file as init.

## Install

```
go get github.com/tetsu-koba/minimumgo
```

## How to use

Blank import this package to your go source file which have main().

```
import _ "github.com/tetsu-koba/minimumgo"
```

You have to add this line before any other imports to initialize this package first.

See example/httpfileserver/httpfileserver.go

## Build

Build your golang executable and put it as `/sbin/init`, or add kernel boot parameter `init=` to start your program as init.

If you have a raspberry pi, see `example/initramfs/README.md` to make initramfs image for raspberry pi.


