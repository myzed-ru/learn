# put your python code here
import re
import requests
import sys

pattern = r"<a.+href[\s]*=[\s]*[\"\']?(\w*:?//|\s*)([^\.]{1,2}[^\/]?[\.\w#]+)[/:\w\.\s]*[\"\']?.*\>"

for base_url in sys.stdin:
    base_url = base_url.strip()
    try:
        res_base = requests.get(base_url)
        # res_parent = requests.get(parent_url, verify=False)

        if res_base.status_code == 200:
            list_domains = re.findall(pattern, str(res_base.content))

        list_unique_domains = []
        for element in list_domains:
            if element[1] not in list_unique_domains:
                list_unique_domains.append(element[1])

        for element in sorted(list_unique_domains):
            print(element)

        else:
            print('Error at connect to ', base_url)
    except:
        print('Error create connect to ', base_url)





