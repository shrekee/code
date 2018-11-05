#!/bin/bash
# KVM 虚拟机管理工具


workdir=`pwd`
xml=/etc/libvirt/qemu
img=/var/lib/libvirt/images


menu() {
cat <<-  EOF

--------------------------------------------------
A: 安装Centos7.5
B: 安装win2K8
C: 退出

--------------------------------------------------
  EOF

}



