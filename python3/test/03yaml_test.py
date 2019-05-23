import yaml
def set_state(state):
    with open('01yaml.yaml') as f:
        doc = yaml.load(f)
        print("doc is: %s"%doc)

    doc['B']['D'] = state

    with open('my_file.yaml', 'w') as f:
        yaml.dump(doc, f,default_flow_style=False)


if __name__ == "__main__":
    set_state("zz")
