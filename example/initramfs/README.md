# Make initramfs for raspberry pi 1/2/3/zero

## Set up raspverry pi to use initramfs

Add this line to /boot/config.txt

```
initramfs initramfs.img followkernel
```

## Make initramfs image

### Specify go src file

Modify `GOSRC` in build.sh

### Put files to __root dir (optional)

for example, /etc/hostname

### Build

Execute `./build.sh`

### Copy `initramfs.img` to /boot of raspberry pi SD card

