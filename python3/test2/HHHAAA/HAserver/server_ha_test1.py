from multiprocessing import Process,Pipe
import time

import sys

import copy

from TCPserver.tcpserver import tcp_server
import libvirt
import configparser
import os


class ServerHa_Agent(object):
    def __init__(self):
        self.config_dict = self.ConfigParser()
        self.missing_host = []
        self.host_allvmstart = []

    def _get_conn(self,ip):
        print('正在连接远程libvirt')
        try:
            conn = libvirt.open('qemu+ssh://zajixuan@%s/system'%ip)
        except:
            print('Failed to connection to the KVM')
            sys.exit(1)
        self.conn = conn

    def ConfigParser(self):
        config_dic = {}
        config = configparser.ConfigParser()
        file = './setting.ini'
        config.read(file)
        config_dic['host'] = config.get('host_cluster', 'server_host')
        config_dic['port'] = config.get('host_cluster', 'port')
        config_dic['host_cluster'] = config.get('host_cluster', 'host_cluster')
        config_dic['maxconn'] = config.get('host_cluster', 'maxconn')
        return config_dic

    def check_socket(self,pipe1):
        host = self.config_dict.get('host')
        port = self.config_dict.get('port')
        maxconn = self.config_dict.get('maxconn')
        all_cluster_hostip = self.config_dict.get('host_cluster')
        all_cluster_hostip = all_cluster_hostip.replace(' ','').split(',')
        tcp_server(host,int(port),all_cluster_hostip,pipe1,int(maxconn))

    def ping_host(self,missing_hostip):
        # IP可以ping通 但是socket接受不到心跳
        ipalive_losetcp = set()
        # IP 不同  没有心跳
        iptcp_alllose = set()
        for ip in missing_hostip:
            cmd = 'ping -c 2 %s' % ip
            res = os.popen(cmd).read()
            if '100% packet loss' in res:
                iptcp_alllose.add(ip)
            else:
                ipalive_losetcp.add(ip)
        # 把vm 都启动的host 删除
        for host in self.host_allvmstart:
            iptcp_alllose.remove(host)
        return ipalive_losetcp,iptcp_alllose


    def open_vm(self, vm, xml):
        # 先获取最优host主机 然后通过libvirt连接根据xml文件开启VM
        try:
            dom = self.conn.createLinux(xml,0)
            print(dom.name())
            print('vm:%s 重启成功'%vm)
        except Exception as e:
            print(e)


    def check_migrate_state(self):
        dom_list = self.conn.listAllDomains(0)
        # 获取最优host 目前所有vm
        #  vm_list 是最优host 所有虚拟机
        besthost_vm_list = [ x.name() for x in dom_list]
        return besthost_vm_list

    def _get_best_host(self):
        # 获取最优的host主机
        best_host = '192.168.43.147'
        return best_host


    def migrate_vm(self,iptcp_alllose):
        iptcp_alllose_cp = copy.deepcopy(iptcp_alllose)

        for host in iptcp_alllose_cp:
            vms = os.popen('ls /sdb_data1/nfs/%s'%host).read()
            # host_vms  可以是host中所有vm的xml
            host_vms = vms.strip().split('\n')
            print('--------migrating host:%s... ---------'%host)
            for vm in host_vms:
                # 检查IP和心跳都没的机器中的vm 是否迁移成功
                besthost_vm_list = self.check_migrate_state()
                if set(host_vms).issubset(set(besthost_vm_list)):
                    print('host:%s 中vm已经都启动完毕...'%host)
                    self.host_allvmstart.append(host)
                if vm in besthost_vm_list:
                    print('vm：%s already start' % vm)
                    continue
                print('---------migrating vm：%s'%vm)
                with open('/sdb_data1/nfs/%s/%s/%s.xml'%(host,vm,vm),'r') as f:
                    xml = f.read()
                    # 根据xml开启vm
                    # conn = libvirt.open('qemu:///system')
                    # dom = conn.createLinux(xml_info, 0)
                    # print(dom.name())
                    self.open_vm(vm, xml)

    def run(self):
        #1 通过发包检测host 是否在线
        pipe1,pipe2 = Pipe()

        p1 = Process(target=self.check_socket,args=(pipe1,))
        p1.start()
        #2 获得不在线的host
        while 1:
            missing_hostip = pipe2.recv()
            if missing_hostip:
                # print('########### losing socket host is :',missing_hostip)
                #  在进行ping等检测
                ipalive_losetcp, iptcp_alllose = self.ping_host(missing_hostip)
                print('\n---------------------------------')
                print(time.strftime("-------%Y-%m-%d %H:%M:%S-------", time.localtime()))
                print('ipalive_losetcp:',ipalive_losetcp)
                print('iptcp_alllose  :',iptcp_alllose)
                print('---------------------------------')
                #3计算其他host中最优host   通过xml配置vm并开机
                if iptcp_alllose:
                    # 先获取最优的host
                    best_host = self._get_best_host()
                    #  libvirt连接最优host 获取conn连接对象
                    self._get_conn(best_host)
                    #  进行 迁移
                    self.migrate_vm(iptcp_alllose)


if __name__ == '__main__':

    a = ServerHa_Agent().run()