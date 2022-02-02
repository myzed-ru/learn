import requests

with open('dataset_24476_3.txt', 'r') as row_file:
    numbers = row_file.readlines()
numbers = [x.rstrip() for x in numbers]


def is_interesting(n):
    response = requests.get(f"http://numbersapi.com/{n}/math", {"json": "true"})
    return 'Interesting' if response.json().get('found') else 'Boring'


with open('result.txt', 'w') as write_file:
    for number in numbers:
        write_file.write(is_interesting(number) + '\n')

