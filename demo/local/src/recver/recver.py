#!/usr/bin/env python

import socket

def main():
    PORT = 8081
    recver = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    addr = ("127.0.0.1", PORT)
    recver.bind(addr)
    while True:
        msg, src = recver.recvfrom(2048)
        print(msg.decode())

if __name__ == "__main__":
    main()
