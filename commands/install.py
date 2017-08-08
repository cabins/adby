# -*- coding: utf-8 -*-
# install.py
# 
#
# Copyright (C) 2017 KongLingCun
#
# This copyrighted material is made available to anyone wishing to use,
# modify, copy, or redistribute it subject to the terms and conditions of
# the GNU General Public License v.3, or (at your option) any later version.
# This program is distributed in the hope that it will be useful, but WITHOUT
# ANY WARRANTY expressed or implied, including the implied warranties of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General
# Public License for more details.  
#
#

import os
import sys
import subprocess

from download import download


def install(name):
    """
    name参数可以是包名，也可以是App名称。
    """
    os.popen('adb wait-for-device')
    apk = os.path.join('temp/' + download(name))

    print("\n正在使用ADB安装，部分手机需要在手机上进行手动确认……\n")
    sub = subprocess.Popen('adb install -r ' + apk, shell=True, stdout=subprocess.PIPE)
    c = sub.stdout.readline()
    while c:
        if c.strip() == 'Success'.encode('utf-8'):
            print('\n安装成功。\n')
        elif c.startswith('Failure'.encode('utf-8')):
            print('\n安装失败。', c.decode('utf-8'), '\n')
        c = sub.stdout.readline()


if __name__ == '__main__':
    install(sys.argv[1])