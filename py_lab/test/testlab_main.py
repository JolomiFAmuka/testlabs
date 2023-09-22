import requests
import json

response = requests.get("https://api.stackexchange.com//2.3/questions?order=desc&sort=activity&site=stackoverflow")

for data in response.json()["items"]:
    print(data["title"])
    print(data["link"])
    print(data["owner"]["display_name"])
    print(data["owner"]["account_id"])
    print(" ")

#print(response)

#print(response.json()["items"])