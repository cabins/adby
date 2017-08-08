# -*- coding: utf-8 -*-
# info.py
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
import re
from tabulate import tabulate

from search import search


def local_apps():
    apps = {}
    # 获取所有安装的软件包
    app_info = os.popen('adb shell pm list packages').readlines()
    app_info = [x.strip().split(':')[1] for x in app_info if x.strip() != '']

    # 拉取文件并分析
    os.mkdir('pull_temp')
    for pkg in app_info:
        print(type(pkg))
        path = os.popen('adb shell pm path ' + pkg).readlines()[0].strip().split(':')[1]
        os.system('adb pull ' + path + ' pull_temp/' + pkg + '.apk')
        
        infos = os.popen('aapt d badging pull_temp/' + pkg + '.apk').readlines()
        for item in infos:
            if not (item.startswith('package:') or item.startswith('application: ')):
                infos.remove(item)
        
        version = re.search(r"versionName='(.*)' ", infos[0]).group(1)
        pkgName = re.search(r"name='(.*)' ", infos[0]).group(1)
        appName = re.search(r"label='(.*)' ", infos[1]).group(1)

        apps[pkg] = {
            'version': version,
            'appName': appName
        }

    print(apps)




def local_info(pkgname):
    """
    获取手机已安装App信息。
    """
    app_info = os.popen('adb shell dumpsys package ' + pkgname).readlines()

    if len(app_info) == 0:
        return {
        "version": '未安装',
        "firstInstallTime": '-',
        "lastUpdateTime": '-'
    }
    else:
        app_info = [x.strip() for x in app_info if '=' in x]

        infos = {}

        for info in app_info:
            [k, v] = info.split('=')[0:2]
            infos[k] = v

        return {
            "version": infos['versionName'] + infos['versionCode'].split()[0],
            "firstInstallTime": infos['firstInstallTime'],
            "lastUpdateTime": infos['lastUpdateTime']
        }



def info(name):
    """
    name可以是包名或者应用名。

    用于查询软件包是否安装。
    """
    candidate = search(name, verbose=False)[0]
    pkgName = candidate['pkgName']
    appName = candidate['appName']
    rversion = candidate['version']

    app_info = os.popen('adb shell dumpsys package ' + pkgName).readlines()

    lversion = local_info(pkgName)['version']

    print(tabulate([[appName, pkgName, lversion, rversion]], headers=['应用名称', '包名', '已安装版本', '最新版本']))



if __name__ == '__main__':
    info(sys.argv[1])