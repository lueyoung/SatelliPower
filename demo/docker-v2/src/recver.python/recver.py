#!/usr/bin/env python

import socket
from info import get_port

def main():
    recver = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    addr = ("0.0.0.0", get_port())
    recver.bind(addr)
    print(addr)
    while True:
        msg, src = recver.recvfrom(2048)
        print(msg.decode())
        print(addr)

if __name__ == "__main__":
    main()
