from flask import Flask, jsonify, request, render_template

app =Flask(__name__)

stores = [
    {
        'name': 'Bibbi stores',
        'items': [
            {
                'name': 'food warmer',
                'price': 15.59
            }
        ]
    }
]

@app.route('/')
def home():
    return render_template('index.html')

#POST - used to receive data (server perspective)
#GET - used to send data (server perspective)

# POST /store data: {name}
@app.route('/store', methods=['POST'])
def create_store():
    request_data = request.get_json()
    new_store = {
        'name': request_data['name'],
        'items': []
    }
    stores.append(new_store)
    return jsonify(new_store)

# GET /store/<string:name>
@app.route('/store/<string:name>') #'http://127.0.0.1:5000/store/some_name'
def get_store(name):
    for store in stores:
        if store['name'] == name:
            return jsonify(store)
    return jsonify({'messge':'No such store name found'})

# GET /store
@app.route('/store')
def get_stores():
    return jsonify({'stores': stores})

# POST /store/<string:name>/item
@app.route('/store/<string:name>/item', methods=['POST'])
def create_item_in_store(name):
    request_data = request.get_json()
    for store in stores:
        if store['name'] == name:
            new_item = [
                {
                    'name': request_data['name'],
                    'price': request_data['price']
                }
            ]
            store['items'].append(new_item)
            return jsonify(new_item)
    return jsonify({'messge':'store not found'})    

# GET /store/<string:name>/item
@app.route('/store/<string:name>/item') 
def get_item_in_store(name):
    for store in stores:
        if store['name'] == name:
            return jsonify({'items': store['items']})
    return jsonify({'messge':'store not found'})

app.run(port=5000)