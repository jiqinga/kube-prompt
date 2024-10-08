#!/bin/bash
set -ex
cd /tmp
wget https://github.com/upx/upx/releases/download/v4.2.4/upx-4.2.4-amd64_linux.tar.xz
tar -xvJf upx-4.2.4-amd64_linux.tar.xz
echo $PATH
which upx
if dpkg -s upx &> /dev/null; then
    apt-get remove -y upx
else
    echo "upx is not installed"
fi
which upx