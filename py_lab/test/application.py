from flask import Flask
from flask_sqlalchemy import SQLAlchemy


app = Flask(__name__)
db = SQLAlchemy(app)

class Drink(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(80), unique=True, nullable=False)
    description = db.Column(db.String(120))

    def __repr__(self):
        return f"{self.name} - {self.description}"


@app.route("/")
def index():
    return "Hello from this side"

@app.route("/drinks")
def get_drinks():
    return {"drinks": "drink data"}

if __name__=="__main__":
    app.run(debug = True)