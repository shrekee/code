#!/bin/bash

SOURCE_XML_FILE='/etc/libvirt/qemu/centos7.0.bak.xml'
SOURCE_IMG_FILE='/opt/vm-images/hosts/centos7.0-4.qcow2'

DEST_XML_DIR='/etc/libvirt/qemu'
DEST_IMG_DIR='/var/lib/libvirt/images'

# xml-source-file: centos7.0.bak.xml
#　使用root用户来执行此脚本
echo　"使用root用户来执行此脚本"

# 用来输入参数
function input(){
    read -p"$1" thing
    if test -z $thing;then
        input $1
    fi
    echo $thing
}

#res=$(input 请输入dom 名字)
domName=$(input 请输入dom名字)
domMemCurrent=$(input 内存)
domMemMax=$(input 最大内存)
domVcpu=$(input 虚拟cpu数)
domuuid=$(uuidgen)
domMac=$(openssl rand -hex 6 | sed -r -e"s/(..)/\1:/g" -e"s/.$//")

dom_full_path=${DEST_XML_DIR}${domName}.xml
cp -a ${SOURCE_XML_FILE} ${dom_full_path}
#NEW_XML_FILE=$(cp -a ${SOURCE_XML_FILE} ${dom_full_path})

sed -i -e "s/dom-name/${domName}/" -e "s/dom-uuid/${domuuid}/" ${dom_full_path}

echo -e "${domName}\t${domMemCurrent}"
echo "$domuuid"
echo "$domMac"

