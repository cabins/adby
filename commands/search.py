# -*- coding: utf-8 -*-
# search.py
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

import sys
import json
from urllib.request import urlopen
from urllib.parse import quote

from tabulate import tabulate


def search(name, verbose=True):
    candidates = []
    url = ''.join(['http://sj.qq.com/myapp/searchAjax.htm?kw=', quote(name)])
    res = json.loads(urlopen(url).read())
    for candidate in res['obj']['items'][0:5]:
        candidates.append({
            'appName': candidate['appDetail']['appName'],
            'pkgName': candidate['appDetail']['pkgName'],
            'version': candidate['appDetail']['versionName'] + str(candidate['appDetail']['versionCode']),
            'fileSize': candidate['appDetail']['fileSize'],
            'authorName': candidate['appDetail']['authorName'],
            'apkMd5': candidate['appDetail']['apkMd5'],
            'apkUrl': candidate['appDetail']['apkUrl'],
            'categoryName': candidate['appDetail']['categoryName'],
            'newFeature': candidate['appDetail']['newFeature'],
            'description': candidate['appDetail']['description']
            })
    
    print()
    table = []
    for index in range(len(candidates)):
        item = candidates[index]
        table.append([index, item['appName'], item['pkgName'], item['version'], item['authorName'], item['apkMd5']])
    
    if verbose:
        print(tabulate(table, headers=['序号', '应用名称', '包名', '版本号', '开发者', 'MD5']))


    return candidates


if __name__ == '__main__':
    search(sys.argv[1])
