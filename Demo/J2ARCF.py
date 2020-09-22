#!/usr/bin/env python3

##### J2A-RCF TOOL #############
##### msft-team@arista.com #####
##### VERSION 1.0 ##############

import csv

def parse():
    pv = ['protocol bgp', 'as-path', 'color', 'route-filter']

    if f != z:
        oper = "and"
    else:
        oper = "{"

    if row[0].find(pv[0]) > 1:
        v1 = row[0].find(pv[0])
        v2 = len(pv[0])
        v = ('source_protocol is BGP')
        print(v)
    elif row[0].find(pv[1]) > 1:
        v1 = row[0].find(pv[1])
        v2 = len(pv[1])
        v = ('as_path match as_path_list' + row[0][v1+v2:])
        print(v)
    elif row[0].find(pv[2]) > 1:
        v1 = row[0].find(pv[2])
        v2 = len(pv[2])
        v = ('ext_community match ext_community_list_COLOR_' + row[0][v1+v2+1:])
        print(v)
    elif row[0].find(pv[3]) > 1:
        v0 = row[0].split()
        v1 = row[0].find(pv[3])
        v2 = len(pv[3])
        v3 = row[0][v1:]
        v4 = row[0][v1:].find('/')
        v5 = v3[v4:v4+3]
        v6 = v3[v2+1:v4]
        global rp
        if (v0[-1]) == 'orlonger':
            rf.append('ip prefix-list' + f.upper() + 'permit '  + v6 + v5 + ' ge ' + v5[1:])
            if rp == True:
                print('prefix match prefix_list', f[1:].upper())
                rp = False
        elif (v0[-1]) == 'exact':
            rf.append('ip prefix-list' + f.upper() + 'permit '  + v6 + v5 + ' ge ' + v5[1:])
            if rp == True:
                print('prefix match prefix_list', f[1:].upper())
                rp = False
        else:
            print('# COULD NOT FIND PREFIX MATCH')
    else:
        print("#UNDEFINED-FROM-POLICY")

def setvalue():
    sv = ['local-preference', 'accept', 'as-path-prepend', 'apply-groups', 'next term', 'reject']

    if row[0].find(sv[0]) > 1:
        v1 = row[0].find(sv[0])
        v2 = len(sv[0])
        v = ('local_preference =' + row[0][v1+v2:])
        print (v, ";")
    elif row[0].find(sv[1]) > 1:
        v1 = row[0].find(sv[1])
        v2 = len(sv[1])
        v = ('return true')
        print (v, ";")
    elif row[0].find(sv[2]) > 1:
        v1 = row[0].find(sv[2])
        v2 = len(sv[2])
        v = ('as_path_prepend ' + row[0][(v1+v2+2):-1])
        print (v, ";")
    elif row[0].find(sv[3]) > 1:
        v1 = row[0].find(sv[3])
        v2 = len(sv[3])
        v = ('# APPLY-GROUPS N/A')
        print (v)
    elif row[0].find(sv[4]) > 1:
        v1 = row[0].find(sv[4])
        v2 = len(sv[4])
        v = ('# MOVE TO NEXT TERM')
        print (v)
    elif row[0].find(sv[5]) > 1:
        v1 = row[0].find(sv[5])
        v2 = len(sv[5])
        v = ('return false')
        print (v, ";")
    else:
        print("#UNDEFINED-THEN-POLICY")

with open('/Users/paulc/Desktop/SWAN/JUNOS_RCF.csv') as f:
    reader = csv.reader(f)
    rp = False
    rf = []
    v = []
    y = "NEW POLICY"
    z = "NEW TERM"
    for row in reader:
        if row[0].find('policy-statement'):
            a = (row[0].find('policy-statement')) + 16
            b = (row[0].find('term'))
            c = (row[0][a:b])
            if c != y:
                if rf != []:
                    ip = '\n'.join([str(item) for item in rf])
                    print(ip)
                print("\n#########################FUNCTION########################\nfunction", c, "(){")
                v.clear()
                rp = True
                rf = []
                z = "NEW TERM"
            y = c
            d = (row[0].find('from'))
            e = (row[0].find('then'))
            if d > 0:
                f = (row[0][b+4:d])
                if f != z:
                    if rf != []:
                        ip = '\n'.join([str(item) for item in rf])
                        print(ip)
                    print ("#########################TERM############################\n#", f.upper(), "\n#########################################################")
                    v.clear()
                    rp = True
                    rf = []
                parse()
                z = f
            if e > 0:
                f = (row[0][b+4:e])
                if f != z:
                    if rf != []:
                        ip = '\n'.join([str(item) for item in rf])
                        print(ip)
                    print ("#########################TERM############################\n#", f.upper(), "\n#########################################################")
                    v.clear()
                    rp = True
                    rf = []
                setvalue()
                z = f
