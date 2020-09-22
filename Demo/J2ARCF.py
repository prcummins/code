#!/usr/bin/env python3

import csv

def parse():
    pv = ['protocol bgp', 'as-path', 'color']

    if f != z:
        oper = "and"
    else:
        oper = "{"

    if row[0].find(pv[0]) > 1:
        v1 = row[0].find(pv[0])
        v2 = len(pv[0])
        v = ('source_protocol is BGP')
        print(v, oper)
    elif row[0].find(pv[1]) > 1:
        v1 = row[0].find(pv[1])
        v2 = len(pv[1])
        v = ('as_path match as_path_list' + row[0][v1+v2:])
        print(v, oper)
    elif row[0].find(pv[2]) > 1:
        v1 = row[0].find(pv[2])
        v2 = len(pv[2])
        v = ('ext_community match ext_community_list_COLOR' + row[0][v1+v2:])
        print(v, oper)
    else:
        print("UNDEFINED")

def setvalue():
    print(row[0])
    sv = ['local-preference', 'accept']
    if row[0].find(sv[0]) > 1:
        v1 = row[0].find(sv[0])
        v2 = len(sv[0])
        v = ('local_preference =' + row[0][v1+v2:])
        print (v, ";")
    if row[0].find(sv[1]) > 1:
        v1 = row[0].find(sv[1])
        v2 = len(sv[1])
        v = ('return true')
        print (v, ";")
    #else:
        #print("UNDEFINED")

with open('/Users/paulc/Desktop/SWAN/JUNOS_RCF.csv') as f:
    reader = csv.reader(f)
    v = []
    y = "NEW POLICY"
    z = "NEW TERM"
    for row in reader:
        if row[0].find('policy-statement'):
            a = (row[0].find('policy-statement')) + 16
            b = (row[0].find('term'))
            c = (row[0][a:b])
            if c != y:
                print("\n#########################FUNCTION########################\nfunction", c, "(){")
                v.clear()
                z = "NEW TERM"
            y = c
            d = (row[0].find('from'))
            e = (row[0].find('then'))
            if d > 0:
                f = (row[0][b+4:d])
                if f != z:
                    print ("#########################TERM############################\n#", f, "\n#########################################################")
                    v.clear()
                parse()
                z = f
            elif e > 0:
                f = (row[0][b+4:e])
                if f != z:
                    print ("#########################TERM############################\n#", f, "\n#########################################################")
                    v.clear()
                setvalue()
                z = f
