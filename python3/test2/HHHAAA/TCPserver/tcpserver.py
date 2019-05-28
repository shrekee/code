#!coding=utf-8

import socket
import threading
import time

def mes_handle(conn,connect_num,ip):
    while True:
        try:
            bytes = conn.recv(1024)
            # print('revice:',bytes.decode('utf-8'))
            if len(bytes) == 0:
                conn.close()
                connect_num.remove(ip)
                print('ip:%s is losing heartbeat '%ip)
        except Exception as e:
            try:
                connect_num.remove(ip)
            except:
                pass
            print('####error: %s:'%ip, e)
            break

def run(s,connect_num,):
    while True:
        conn, address = s.accept()
        ip = address[0]
        connect_num.append(ip)
        # 超时时间
        conn.settimeout(20)
        # 接受数据大小
        thread = threading.Thread(target=mes_handle,args=(conn,connect_num,ip))
        thread.setDaemon(True)
        thread.start()


def tcp_server(host,port,all_cluster_hostip,pipe1,maxconn=10):
    addr = (host,port)
    connect_num = []
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind(addr)
    s.listen(maxconn)
    print('Tcp server is start,wait for connection')
    thread = threading.Thread(target=run,args=(s,connect_num))
    thread.setDaemon(True)
    thread.start()
    while True:
        time.sleep(10)
        missing_host = [x for x in all_cluster_hostip if x not in connect_num]
        print('运行正常的主机列表：',list(set(connect_num)))
        if missing_host:
            pipe1.send(missing_host)
