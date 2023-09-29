import requests
import json

# This script calls the canada holiday API (https://canada-holidays.ca/api)

response = requests.get("https://canada-holidays.ca/api/v1/provinces")

data = response.json()

# print(type(data["provinces"]))

# print(len(data["provinces"]))

# print(type((data["provinces"])[0]))
# print("**********")
# print((data["provinces"])[0])

print(((data["provinces"])[0]).keys())

# for prov in data["provinces"]:
#     print(prov['id'])


#print(response.json())

#print(response.json()['provinces'])

# for data in (response.json()['provinces']):
#     print(data)
