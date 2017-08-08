# -*- coding: utf-8 -*-
# cleancache.py
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

def clean_cache():
    print('\n正在清理缓存……\n')
    for apk in os.listdir('temp'):
        print(apk)
        os.remove('temp/' + apk)
    print('\n清理完成。')


if __name__ == '__main__':
    clean_cache()