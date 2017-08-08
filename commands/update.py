# -*- coding: utf-8 -*-
# update.py
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

from search import search
from info import local_info
from install import install
from utils import format_version, cmp_version


def update(*names):
    """
    接受包名或者应用名作为参数，如果参数不传，则更新全部应用。
    """

    packages = []
    skip_pkgs = []
    if len(names) == 0:
        # 更新全部
        print("Getting all the pkgs.")
        lists = os.popen('adb shell pm list packages -3').readlines()
        print("Here")
        packages = [x.strip().split(':')[1] for x in lists if x.startswith('package:')]
        print("Done.")
    else:
        packages.append(search(name, False)[0]['pkgName'])

    for pkgName in packages:
        # 检查需要更新的包，记录
        print(pkgName)
        lversion = local_info(pkgName)['version'].split('.')
        rversion = search(pkgName, False)[0]['version'].split('.')
        try:
            lversion = [int(x) for x in lversion]
            rversion = [int(x) for x in rversion]
            (lversion, rversion) = format_version(lversion, rversion)

            if cmp_version(lversion, rversion) != 1:
                packages.remove(pkgName)
        except ValueError:
            skip_pkgs.append(pkgName)

    if len(packages) == 0:
        print("\n所有应用都是最新版本。\n")
        return
    
    for package in packages:
        # install(package)
        print(package)

    print("\n以下应用未更新，请手动完成更新:")
    print(skip_pkgs)
    


if __name__ == '__main__':
    update()