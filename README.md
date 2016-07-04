statusAS-api
============

Minimalistic API to monitor the status of your services.

You need a MongoDB database, with a DB named "statusAS" and a collection named "services", anyway you can change whatever you like in the configuration.

Create a Service:
```
$ curl -i -X POST http://192.168.125.1:8080/api/v1/services -d '{"name": "service X", "description": "service X description", "link": "https://serviceX.yourcom.com", "status": 0, "group_id": 0, "messages": []}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/v1/services/5779ccca802abd1464cc9e45
Date: Mon, 04 Jul 2016 02:41:14 GMT
Content-Length: 0
```

Get Services:
```
$ curl -X GET http://192.168.125.1:8080/api/v1/services
[
    {
        "id": "5779b85b802abd244c797ca2",
        "name": "Service #1",
        "description": "Service1 Description",
        "link": "https://service1.local",
        "status": 1,
        "group_id": 0,
        "messages": [
            "service component A is down",
            "service componente Y is down"
        ],
        "created_at": "2016-07-03T20:44:03.944-04:30",
        "updated_at": "2016-07-03T21:17:15.517-04:30"
    },
.
.
.
]
```

Update Service:
```
$ curl -i -X PUT http://192.168.125.1:8080/api/v1/services/5779ccca802abd1464cc9e45 -d '{"name": "service X", "description": "service X description", "link": "https://serviceX.yourcom.com", "status": 1, "group_id": 0, "messages": ["Service component Y is down", "Service component Z is down"], "created_at": "2016-07-03T20:44:03.944-04:30"}'
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Mon, 04 Jul 2016 02:44:38 GMT
```

Delete a component:
```
$ curl -i -X DELETE http://192.168.125.1:8080/api/v1/services/5779ccca802abd1464cc9e45
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Mon, 04 Jul 2016 02:45:29 GMT
```

TODO: add TLS and a Token for POST, PUT and DELETE.
