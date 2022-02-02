import csv
import operator
count_crimes = dict()


with open('Crimes.csv', "r") as csv_file:
    reader = csv.DictReader(csv_file)
    for row in reader:
        current_crime_type = row.get('Primary Type')
        if current_crime_type in count_crimes.keys():
            count_crimes.update({current_crime_type: count_crimes.get(current_crime_type) + 1})
        else:
            count_crimes.update({current_crime_type: 1})

print(max(count_crimes.items(), key=operator.itemgetter(1))[0])
print(count_crimes)