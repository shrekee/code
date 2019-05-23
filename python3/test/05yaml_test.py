import yaml
import re
def set_state(state):
    with open('02yaml.yaml') as f:
        doc = yaml.load(f)
        print("doc is: %s"%doc)

    #doc['cluster']['clueter01']['host'] = state
    print("clusters\t:{0}".format(doc['cluster']['cluster01']))
    print("hosts\t:{0}".format(doc['cluster']['cluster01']['host']))
    print("type of hosts\t:{0}".format(type(doc['cluster']['cluster01']['host'])))

    #for host in doc['cluster']['cluster01']['host']:
    #    for key in doc['cluster']['cluster01']['host'].keys():
    #        if re.match('^host\d+',key):
    #            print(host)
    #            print(key)
                #print(key.get('name'))
                #print(host[key])
                #print("host type is : %s"%type(host))
                #print(host.get(key))
    #    yaml.dump(doc, f,default_flow_style=False)
    for key in doc['cluster']['cluster01']['host']:
            if re.match('^host\d+',key):
                print(doc['cluster']['cluster01']['host'].get(key).get('name'))
                print("type is\t:%s"%type(doc['cluster']['cluster01']['host'].get(key).get('name')))


if __name__ == "__main__":
    set_state("zz")
