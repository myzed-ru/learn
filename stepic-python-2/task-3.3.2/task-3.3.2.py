# put your python code here
import re
import requests
import sys

list_unique_domains = []
# pattern = r"<a.*href\s*=\s*[\"\']?(?:\w*://|\s*)([^\.]{2,}[^\/]?[a-zA-Z]+[\.\w-]+)[/:\w\.\s]*[\"\']?\s*\>"
pattern = r"<a.*href\s*=\s*[\"\']?(?:\w*://|\s*)([^\.]{2,}[^\/]?[a-zA-Z]+[\.\w-]+).*[\"\']?\s*\>"

for base_url in sys.stdin:
    base_url = base_url.strip()
    try:
        res_base = requests.get(base_url)
        if res_base.status_code == 200:
            content = res_base.content.decode("utf-8")
            all_lines = content.splitlines()
            for line in all_lines:
                domain = re.findall(pattern, line)
                # print(f"row line: {line}")
                if domain and list(domain)[0] not in list_unique_domains:
                    # print(f"domain: {list(domain)[0]}")
                    list_unique_domains.append(list(domain)[0])

            for element in sorted(list_unique_domains):
                print(element)

        else:
            print('Error at connect to ', base_url)
    except Exception as e:
        print(e)
        print('Error create connect to ', base_url)

