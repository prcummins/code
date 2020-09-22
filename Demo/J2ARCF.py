#!/usr/bin/env python3

import csv

def parse():
    pv = ['protocol bgp', 'as-path', 'color']
    if row[0].find(pv[0]) > 1:
        v1 = row[0].find(pv[0])
        v2 = len(pv[0])
        v.append('source_protocol is BGP')
    elif row[0].find(pv[1]) > 1:
        v1 = row[0].find(pv[1])
        v2 = len(pv[1])
        v.append('as_path match as_path_list' + row[0][v1+v2:])
    elif row[0].find(pv[2]) > 1:
        v1 = row[0].find(pv[2])
        v2 = len(pv[2])
        v.append('ext_community match ext_community_list_COLOR' + row[0][v1+v2:])
    else:
        print("UNDEFINED")

def setvalue():
    sv = ['local preference']
    if row[0].find(sv[0]) > 1:
        v1 = row[0].find(sv[0])
        v2 = len(sv[0])
        v.append('local_preference =' + row[0][v1+v2:])
        print(v2)
    #print(row[0])

with open('/Users/paulc/Desktop/SWAN/JUNOS_RCF.csv') as f:
    reader = csv.reader(f)
    cmd = ""
    v = []
    y = "NEW POLICY"
    z = "NEW TERM"
    for row in reader:
        if row[0].find('policy-statement'):
            a = (row[0].find('policy-statement')) + 16
            b = (row[0].find('term'))
            c = (row[0][a:b])
            if c != y:
                print("function", c, "(){")
            y = c
            d = (row[0].find('from'))
            e = (row[0].find('then'))
            if d > 0:
                f = (row[0][b+4:d])
                if f != z:
                    print ("###TERM###", f)
                    v.clear()
                parse()
                cmd = " and \n".join(v)
                z = f
            elif d < 0:
                if cmd != "":
                    print("if", cmd, "{")
                    v.clear()
                f = (row[0][b+4:e])
                if f != z:
                    print ("###TERM###", f)
                    v.clear()
                setvalue()
                cmd = " and \n".join(v)
                z = f
