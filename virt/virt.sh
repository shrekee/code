#!/bin/bash
# KVM虚拟机管理工具
#

workdir=/home
xml=/etc/libvirt/qemu
img=/var/lib/libvirt/images
vmxmlroot=$workdir/vmm
vmdiskroot=$workdir/vmm

menu() {
	cat <<- EOF
	------------------------------------------
	A.部署CentOS7u5系统
	B.部署Win2k8r2系统
	------------------------------------------
	EOF
}
menu

read -p "请选择需要安装的操作系统:" choose


case $choose in
	A|a)
		read -p "请输入您想使用的域名称:" name

		vmname=${name}
		vmuuid=$(uuidgen)
		vmdisk=${name}.qcow2
		vmmac="52:54:$(dd if=/dev/urandom count=1 2>/dev/null | md5sum | sed -r 's/^(..)(..)(..)(..).*$/\1:\2:\3:\4/')"
		
		cp $vmxmlroot/centos7u5/centos7u5.xml $xml/${vmname}.xml
		qemu-img create -f qcow2 -b $vmdiskroot/centos7u5/centos7.5.qcow2 $img/${vmdisk} &> /dev/null
		sed -ri "s#vmname#${vmname}#" $xml/${vmname}.xml
		sed -ri "s#vmuuid#${vmuuid}#" $xml/${vmname}.xml
		sed -ri "s#vmdisk#$img/${vmdisk}#" $xml/${vmname}.xml
		sed -ri "s#vmmac#$vmmac#" $xml/${vmname}.xml
		virsh define $xml/${vmname}.xml 
		;;

	B|b)

		echo "win2k8r2已安装完成"
		;;
	*)
		echo "请按菜单选择（A|B）"
esac
virt-manager
