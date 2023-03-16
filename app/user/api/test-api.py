import json

import requests


def test_user_get_info():
    url = "http://127.0.0.1:8888/dev/v1/user/getInfo"
    data = {
        "userId": 80
    }
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.post(url, headers=headers, data=json.dumps(data))
    print(res.text)

def test_user_insert_info():
    url = "http://127.0.0.1:8888/dev/v1/user/insertInfo"
    data = {
        "name": "123"
    }
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.post(url, headers=headers, data=json.dumps(data))
    print(res.text)

def test_bike_rpc_getInfo ():
    url = "http://127.0.0.1:8888/dev/v1/bike/rpc/getInfo"
    data = {
        "bikeId": 2
    }
    headers = {
        'Content-Type': 'application/json'
    }

    res = requests.post(url, headers=headers, data=json.dumps(data))
    print(res.text)

if __name__ == '__main__':
    # test_user_get_info()
    # test_user_insert_info()
    test_bike_rpc_getInfo()