#!/bin/bash
rm -frv build
make clean && make mrproper
mkdir -pv build/out/modules build/out/dt_image
export ARCH=arm
export CROSS_COMPILE=../arm-eabi-4.9/bin/arm-eabi-
export STRIP=../arm-eabi-4.9/bin/arm-eabi-strip
make pxa1088_degaswifi_usa_defconfig
make menuconfig && wait
make CONFIG_NO_ERROR_ON_MISMATCH=y && make modules
./tools/dtbTool -o build/out/dt_image/boot.img-dt -p ./scripts/dtc/ ./arch/arm/boot/dts/
cp arch/arm/boot/*zImage build/out/boot.img-zImage
find -type f -name *.ko -exec cp {} build/out/modules/ \;
ls -al build/out/modules/
cd build/out/modules/
$STRIP --strip-unneeded *.ko
cd ../../../
ls -al build/out/modules/ build/out/dt_image/ build/out/
echo Done !
