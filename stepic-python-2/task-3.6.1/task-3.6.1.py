# put your python code here
import sys
import xml.etree.ElementTree as ET

colors_info = {
    'red': 0,
    'green': 0,
    'blue': 0
}


def get_weight(element, weight):
    current_color = element.attrib.get('color')
    colors_info.update({current_color: colors_info.get(current_color) + weight})
    for child in element:
        get_weight(child, weight + 1)


for xml_row in sys.stdin:
    get_weight(ET.fromstring(xml_row), 1)
result = ' '.join([str(x) for x in colors_info.values()])
print(result)
