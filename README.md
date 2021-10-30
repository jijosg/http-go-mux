# http-go-mux [![Go](https://github.com/jijosg/http-go-mux/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/jijosg/http-go-mux/actions/workflows/go.yml)

A sample REST api using go and mux. It utilizes SQLLite for database storage.
Also intends to publish the executable as a release using GitHub actions.

## Build using Docker

`docker build -t http-go-mux:1.0 .`

`docker run -it --rm -p 8000:8000 http-go-mux:1.0`

### Perform CRUD on API
Creates User table 
```
> curl localhost:8000/create
Created table user
```
Insert Data into User table 
```
> curl -X POST -d '{"name" : "John"}' localhost:8000/insert
Inserted row for John
```
List Data from User table 
```
> curl -X GET localhost:8000/list
[
 {
  "id": 1,
  "name": "John"
 }
]
```

Delete Data from User table
```
> curl -X DELETE localhost:8000/delete/1
Deleted row for 1
```
