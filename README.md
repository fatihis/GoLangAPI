# GoLang REST API

REST API W / GoLang and MongoDB

## Description

Simple REST API over CRUD Operations through MongoDB Atlas

## Getting Started

### Dependencies
	github.com/gofiber/fiber v1.14.6
	github.com/mitchellh/mapstructure v1.4.1
	go.mongodb.org/mongo-driver v1.7.1

### Executing program

For Unix Users
```
$ chmod +x hello.go
$ ./hello.go
```

For Unix Users
```
$ go run ./main.go
```
## API ENDPOINTS
### Postman Workspace
https://www.postman.com/flight-geologist-71319485/workspace/fill-labs-public-repo/overview
### Get All Persons
- Path : `/getAll`
- Method: `GET`

### Create Post
- Path : `/create`
- Method: `POST`
- Fields: `_id, FirstName, LastName, Email, Age`

### Details a Post
- Path : `/get/{id}`
- Method: `GET`

### Update Post
- Path : `/update/{id}`
- Method: `PUT`
- Fields: `_id, FirstName, LastName, Email, Age`


### Delete Post
- Path : `/delete/{id}`
- Method: `DELETE`


## Help

For requests and help, please contact at bfatihisik@gmail.com 

## Authors

Contributors names and contact info


 [@fatihis](https://github.com/fatihis)

## Version History

* 0.1
    * Initial Release
    * See [commit change]() or See [release history]()


