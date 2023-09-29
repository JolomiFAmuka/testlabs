import requests
import json

response = requests.get("https://api.stackexchange.com//2.3/questions?order=desc&sort=activity&site=stackoverflow")

for data in response.json()["items"]:
    if data['answer_count'] == 0:
        print(f"Question: {data['title']}")
        print(data['link'])
    else:
        print("skipped")
    print("*****")

# for data in response.json()["items"]:
#     # print(data["title"])
#     print(data["link"])
# #     print(data["owner"]["display_name"])
# #     print(data["owner"]["account_id"])
# #     print(" ")

#print(response.json())

#print(((response.json()["items"])[1]).keys())