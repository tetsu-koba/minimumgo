# Make initramfs for raspberry pi 1/2/3/zero

## Set up raspberry pi to use initramfs

Add this line to `/boot/config.txt`

```
initramfs initramfs.img followkernel
```

Add `ip=dhcp` and specify ntp server name by `minimumgo.ntp=` in `/boot/cmdline.txt`

For example,

```
console=serial0,115200 ip=dhcp minimumgo.ntp=ntp.nict.jp
```

## Make initramfs image

### Specify go src file

Modify `GOSRC` in ./build.sh

### Put files to __root dir (optional)

for example, `/etc/hostname` to specify hostname.

### Build and install

Execute `./build.sh`

Copy `initramfs.img` to /boot of raspberry pi SD card
