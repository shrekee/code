import yaml
def set_state(state):
    with open('02yaml.yaml') as f:
        doc = yaml.load(f)
        print("doc is: %s"%doc)

    #doc['cluster']['clueter01']['host'] = state
    print(doc['cluster']['cluster01']['host'])

    with open('my_file02.yaml', 'w') as f:
        yaml.dump(doc, f,default_flow_style=False)


if __name__ == "__main__":
    set_state("zz")
