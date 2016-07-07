statusAS-api
============

Minimalist API to monitor the status of your services.

You need a MongoDB database, with a DB named "statusAS" and two collections named "services" and "tokens", anyway you can change whatever you like in the configuration.

Config file (config.toml):
```toml
# Example of Configuration
[apiserver]
name = "server1.mydom.local"
port = 8080
mongoserver = "serverdb.mydom.local"
mongoport = 27017
```

In case you delete or lose the config file you can generate another from a template.
```
$ ./statusAS-api -template
config.toml created.
```

Before you run the server, you need to add a private key (named "server.key") and a cert (named "server.pem"), you can create it using openssl.
```
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

Now we need to add permanent (not the best choice, but becomes easier, probably in another version we can work with full jwt) tokens for POST, PUT and DELETE methods, so in mongoDB we have to create a tokens collection.

In your cli create a base64 code.
```
$ echo -n 'user1 mydom.local' | base64
dXNlcjEgbXlkb20ubG9jYWw=
$ echo -n dXNlcjEgbXlkb20ubG9jYWw= | base64 -d
user1 mydom.local
```

In mongoDB create the collections and insert the token.
```
> use statusAS
switched to db statusAS
> db.createCollection("services")
{ "ok" : 1 }
> db.createCollection("tokens")
{ "ok" : 1 }
> db.tokens.insert({"token": "dXNlcjEgbXlkb20ubG9jYWw="})
WriteResult({ "nInserted" : 1 })
> db.tokens.find()
{ "_id" : ObjectId("577d941861442012e950526f"), "token" : "dXNlcjEgbXlkb20ubG9jYWw=" }
```

Now lets test the api with curl, remember if you have some trouble working with curl and your self signed cert, then add option (-k) to curl.

```
$ go build
$ ./statusAS-api
```

Create a Service:
```
$ curl -i -H "X-StatusAS-Token: dXNlcjEgbXlkb20ubG9jYWw=" -X POST https://server1.mydom.local:8080/api/v1/services -d '{"name": "service X", "description": "service X description", "link": "https://serviceX.yourcom.com", "status": 0, "group_id": 0, "messages": []}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/v1/services/5779ccca802abd1464cc9e45
Date: Mon, 04 Jul 2016 02:41:14 GMT
Content-Length: 0
```

Get Services:
```
$ curl -X GET https://server1.mydom.local:8080/api/v1/services
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
$ curl -i -H "X-StatusAS-Token: dXNlcjEgbXlkb20ubG9jYWw=" -X PUT https://server1.mydom.local:8080/api/v1/services/5779ccca802abd1464cc9e45 -d '{"name": "service X", "description": "service X description", "link": "https://serviceX.yourcom.com", "status": 1, "group_id": 0, "messages": ["Service component Y is down", "Service component Z is down"], "created_at": "2016-07-03T20:44:03.944-04:30"}'
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Mon, 04 Jul 2016 02:44:38 GMT
```

Delete a component:
```
$ curl -i -H "X-StatusAS-Token: dXNlcjEgbXlkb20ubG9jYWw=" -X DELETE https://server1.mydom.local:8080/api/v1/services/5779ccca802abd1464cc9e45
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Mon, 04 Jul 2016 02:45:29 GMT
```

Models:
```go
type Service struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Link        string        `bson:"link" json:"link"`
	Status      int           `bson:"status" json:"status"`
	GroupID     int           `bson:"group_id" json:"group_id"`
	Messages    []string      `bson:"messages" json:"messages"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}
```

Status codes should be 0 ("OPERATIONAL") or 1 (NOT OPERATIONAL).
