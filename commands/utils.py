# -*- coding: utf-8 -*-
# utils.py
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

def format_version(lversion, rversion):
    """
    lversion and rversion both should be list such as [3, 4, 0], which means 3.4.0
    """

    llen = len(lversion)
    rlen = len(rversion)

    if llen == rlen:
        return (lversion, rversion)
    else:
        if llen > rlen:
            tmp = [0] * llen
            tmp[0:rlen-1] = rversion[0:rlen-1]
            tmp[-1] = rversion[-1]
            return (lversion, tmp)
        else:
            tmp = [0] * rlen
            tmp[0:llen-1] = lversion[0:llen-1]
            tmp[-1] = lversion[-1]
            return (tmp, rversion)


def cmp_version(lversion, rversion):
    """
    lversion and rversion both should be list such as [3, 4, 0], which means 3.4.0

    Before use this method, you should format version with format_version method.

    版本号对比，返回值如下：
    1. 本地版本＞远程版本，返回-1
    2. 本地版本＝远程版本，返回0
    3. 本地版本＜远程版本，返回1
    """
    length = len(lversion)
    for i in range(length):
        if lversion[i] > rversion[i]:
            return -1
        elif lversion[i] < rversion[i]:
            return 1
        else:
            continue
    return 0


if __name__ == '__main__':
    (lversion, rversion) = format_version([1,4,259], [2,3,0,1,262])
    print(cmp_version(lversion, rversion))