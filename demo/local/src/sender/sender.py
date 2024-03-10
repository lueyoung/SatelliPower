#!/usr/bin/env python

import socket
from typed import *
from time import sleep
from random import random

def main():
    srv = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    dest = ("127.0.0.1", 8080)
    counter = 0
    while True:
        method = random() 
        if method < 0.5:
            d = get_dict1()
        else:
            d = get_dict2()
        d['ID'] = str(counter)
        counter += 1
        msg = make_json(d, id=counter)
        print(msg)
        srv.sendto(msg.encode(), dest)
        sleep(1)

if __name__ == "__main__":
    main()
