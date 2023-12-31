from fastapi import FastAPI
from fastapi.params import Body

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Welcome to my API!!!!!!"}

@app.get("/posts")
def get_posts():
    return {"data": "this is your post"}

@app.post("/createposts")
def create_posts(payload: dict = Body(...)):
    print(payload)
    return{"message": f"title {payload['title']} content {payload['content']}"}