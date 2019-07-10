# coding:utf-8

import threading
from queue import Queue

def fun1():
    print("func1...")
    fun2()

def fun2():
    print("func2...")


if __name__ == "__main__":
    t = threading.Thread(target=fun1)
    t.start()
    t.join()

    print("All Done...")
