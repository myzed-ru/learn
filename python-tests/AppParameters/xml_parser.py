import xml.etree.ElementTree as ET

app_params = dict()

tree = ET.ElementTree(file='./Local.1Node.xml')
parameters = next(iter(tree.getroot()))

for parameter in parameters:
    key = parameter.attrib.get('Name')
    value = parameter.attrib.get('Value')
    app_params.update({key: value})

print(app_params)
