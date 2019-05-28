#coding:utf8
import socket
import time
from multiprocessing import Process,Pipe
import os
import configparser


class ClientHa_Agent(object):
    def __init__(self):
        self.get_info()

    def ConfigParser(self):
        config_dic = {}
        config = configparser.ConfigParser()
        file = './setting.ini'
        config.read(file)
        config_dic['server_host'] = config.get('client_settings', 'server_host')
        config_dic['port'] = config.get('client_settings', 'port')
        config_dic['sendsocket_delay'] = config.get('client_settings', 'sendsocket_delay')
        config_dic['getvmxml_delay'] = config.get('client_settings', 'getvmxml_delay')
        return config_dic

    def get_info(self):
        self.config_dict = self.ConfigParser()
        self.host = self.config_dict.get('server_host')
        self.port = int(self.config_dict.get('port'))
        self.sendsocket_delay = self.config_dict.get('sendsocket_delay')
        self.getvmxml_delay = int(self.config_dict.get('getvmxml_delay'))

    def tcp_conn(self):
        s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
        s.connect((self.host, self.port))
        keepclass = "actived"
        while True:
            time.sleep(self.sendsocket_delay)
            s.send(bytes(keepclass,'UTF-8'))


    def send_xml(self):
        path = "/etc/libvirt/qemu"
        dirs = os.listdir(path)
        allvm_xml = []
        hostname = socket.gethostname()
        for file in dirs:
            if file.endswith('.xml'):
                allvm_xml.append(file)
                vm_filename = file[:-4]
                path = '/var/nfs/public/%s/%s' % (hostname, vm_filename)
                # read vm_xml
                xml_info = os.popen('cat /etc/libvirt/qemu/%s.xml' % vm_filename).read()
                isExists = os.path.exists(path)
                if not isExists:
                    os.makedirs(path)
                with open('%s/%s.xml' % (path, vm_filename), 'w') as f:
                    f.write(xml_info)


    def control_send_xml(self):
        #  循环向server发送    如果更改xml主动发送
        while 1:
            time.sleep(self.getvmxml_delay)
            self.send_xml()


    def run(self):
        # 进程1：发送tcp包确保在线
        p1 = Process(target=self.tcp_conn,)
        p1.start()
        # 进程2：发送xml文件给server端
        p2 = Process(target=self.control_send_xml)
        p2.start()


if __name__ == '__main__':
    client = ClientHa_Agent().run()