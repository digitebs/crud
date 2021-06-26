Simple restful api written in golang that inserts in mongdb

##### Structure
```
.
├── config
├── controllers
├── models
├── routes 
└── main.go
```
##### Prerequisite
install golang: https://golang.org/dl/

install mongo db:
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/

once install create database and collection:
```shell
use db
db.createCollection("test")
```

##### Building

`go build .`

##### Running
`./crud`

##### Test api

http://localhost:8090/user/{id}

where {id} is id of the document. 
change the method for crud operation(POST,PUT,DELETE,GET)

id is automatically generated when you create a document

**Insert**
```shell
curl -X POST \
  http://localhost:8090/user/ \
  -d '{  
"name": "test",
"dob": "2020-01-01", 
"address": "sg", 
"description": "test", 
"createdAt": "2020-01-01" 
} 
'
```
**Update**
```shell

curl -X PUT \
  http://localhost:8090/user/xxx \
  -d '{  
"name": "me",
"dob": "2020-01-01", 
"address": "sg", 
"description": "test", 
"createdAt": "2020-01-01" 
}'
```

**Delete**
```shell

curl -X DELETE \
  http://localhost:8090/user/xxx

```
**Get**
```shell
  
curl -X GET \
  http://localhost:8090/user/xxx
```