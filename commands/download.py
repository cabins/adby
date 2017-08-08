# -*- coding: utf-8 -*-
# download.py
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
import hashlib
from urllib.request import urlretrieve
from search import search


def progress(a, b, c):
    """
    下载进度回调方法。

    a:已经下载的数据块
    b:数据块的大小
    c:远程文件的大小
    """

    per = 100.0 * a * b / c
    if per > 100:
        per = 100

    print('%.2f%%' % per, a*b, '/', c, end='\r')


def get_md5(filename, _md5):
    """
    校验MD5
    """
    f = open(filename, 'rb')
    md5 = hashlib.md5(f.read()).hexdigest().upper()
    f.close()
    return md5 == _md5.upper()


def download(name):
    """
    下载，接受的参数是包名或者应用名称。
    """
    candidate = search(name)[0]
    url = candidate['apkUrl']
    pkgName = candidate['pkgName']
    md5 = candidate['apkMd5']

    os.mkdir('temp')
    filename = os.path.join('temp/', pkgName + '.apk')

    if os.path.exists(filename) and os.path.isfile(filename) and get_md5(filename, md5):
        print("\n文件存在，校验通过。\n")
    else:
        print("\n已将第0个候选设置为下载目标，若不正确，请输入准确的包名或软件名称作为参数。\n")
        print("\n正在从以下地址下载： " + url + "\n")

        urlretrieve(url, filename=os.path.join('temp/', pkgName + ".apk"), reporthook=progress)
        print("\n下载完成.\n")

    return pkgName + ".apk"

if __name__ == '__main__':
    download(sys.argv[1])