# put your python code here
import sys
import json


def create_relation(base_dict, parent_name, child_name=None):
    child_name = child_name if child_name else parent_name
    if parent_name not in base_dict.keys():
        base_dict.update({parent_name: [child_name]})


def add_parents_info(base_dict, row_base_dict, parent, child):
    if parent in base_dict.keys():
        children = base_dict.get(parent)
        if child not in children:
            children.append(child)
            base_dict.update({parent: children})
    else:
        base_dict.update({parent: [parent, child]})
    # print('----')
    # print(f'parent: {parent} , child: {child}')
    # print(base_dict)
    for grandparent in row_base_dict.get(parent):
        # print(f'add_parents_info(base_dict, {grandparent}, {child})')
        add_parents_info(base_dict, row_base_dict, grandparent, child)


json_obj = json.load(sys.stdin)
row_relations = dict()
for class_info in json_obj:
    row_relations.update({class_info.get('name'): class_info.get('parents')})

relations = dict()
for child_info, parents_info in row_relations.items():
    create_relation(relations, child_info)
    for parent_info in parents_info:
        add_parents_info(relations, row_relations, parent_info, child_info)


# print(relations)
result_dict = dict()
for class_name in sorted(relations.keys()):
    print(f'{class_name} : {len(relations.get(class_name))}')
