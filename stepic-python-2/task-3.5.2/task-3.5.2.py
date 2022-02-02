import requests
from operator import itemgetter

MY_TOKEN = 'eyJhbGciOiJIUzI1NiJ9.eyJyb2xlcyI6IiIsInN1YmplY3RfYXBwbGljYXRpb24iOiI2MDdhODRlZjA2MWI2ZDAwMGUyNzAxNTAiLCJleHAiOjE2MTkyNDY5NTksImlhdCI6MTYxODY0MjE1OSwiYXVkIjoiNjA3YTg0ZWYwNjFiNmQwMDBlMjcwMTUwIiwiaXNzIjoiR3Jhdml0eSIsImp0aSI6IjYwN2E4NGVmZjI0NjJhMDAwZTQxZmVjOSJ9.HkPkv_sOK-wNYxVglPm-oVjprowtEiqdufY1L721K6Y'


def get_artist_info(artist_id):
    response = requests.get(url=f"https://api.artsy.net/api/artists/{artist_id}", headers={"X-Xapp-Token": MY_TOKEN})
    response.encoding = 'utf-8'
    row_dict = response.json()
    return row_dict.get('sortable_name'), int(row_dict.get('birthday'))


with open('dataset_24476_4.txt', 'r') as row_file:
    artists = row_file.readlines()

artists = [x.rstrip() for x in artists]
artists_info = list()
for artist in artists:
    artists_info.append(get_artist_info(artist))


sorted_by_age = sorted(artists_info, key=itemgetter(1, 0))
print(artists_info)
print(sorted_by_age)

with open('result.txt', 'w', encoding='utf-8') as write_file:
    for artist in sorted_by_age:
        write_file.write(artist[0] + '\n')