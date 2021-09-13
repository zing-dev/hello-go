#! /usr/bin/env python
# _*_ coding:UTF-8 _*_
import os
import pty
import select


def mkpty():
    # 打开虚拟终端
    master1, slave = pty.openpty()
    name1 = os.ttyname(slave)
    master2, slave = pty.openpty()
    name2 = os.ttyname(slave)
    print('虚拟设备名称: ', name1, name2)
    return master1, master2


if __name__ == "__main__":
    m1, m2 = mkpty()
    while True:
        rl, wl, el = select.select([m1, m2], [], [], 1)
        for master in rl:
            data = os.read(master, 128)
            print("read %d data." % len(data))
            if master == m1:
                os.write(m2, data)
            else:
                os.write(m1, data)
