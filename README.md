# The Foo Go Webserver
Foo App is a simple Go webserver given to me as go testing challenge.
It is a small webserver that allows us to run CRUD(or CRD in our case) operations on a small data set.

## Project Requirements
In order to run this project you will need the following already installed 
 
 - [Docker](https://www.docker.com/)
 - [Docker Compose](https://docs.docker.com/compose/)
 - [make](https://www.gnu.org/software/make/) 

Please follow the documentation for installation for your operating system.

## Building, Running & Clean up
The following are a list of commands to help build run and maintain the project

### make build
```make build```
This should be the first command you run this builds the project.

### make run
```make run```
once this command has run your service should be ready to go. 
You should now be able to hit the service on port `8080`

### make clean
```make clean```
once you are done simply run  the clean command to bring the service down and have it clean up after it's self.


## Original Acceptance Criteria 
You will be building a very simple webservice in go that will store “foo” records in memory. Please feel free to use any public resources that you can find. 


### Requirements

1.	Please commit to a public Github repository.
2.	Please include a Makefile that supports “build”, “run”, and “clean” tasks. 
3.	The webservice should run on port 8080.
4.	The foo data structure only needs 2 fields. A “name” field and an “id” field. Both should be string data types. 
5.	It should support a POST endpoint (‘/foo’) that accepts a json foo object and responds with a 200 response code. The value of the id field should be added by this endpoint using a generated UUID.
6.	It should support a GET endpoint (‘foo/{id}’) that responds with a 200 response code if the record is found, or a 404 response code if not found.
7.	It should support a DELETE endpoint (‘foo/{id}’) that responds with a 204 response code if the record is found, or a 404 response code if not found.
8.	Please make sure that your code is formatted with gofmt before committing.
9.	Please send the url of your git repository when you are complete.



### Sample curl output

```$ curl -i -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' 'http://localhost:8080/foo' -d '{"name": "Jack"}'```

```
HTTP/1.1 200 
Content-Type: application/json
Transfer-Encoding: chunked
Date: Fri, 11 Feb 2022 15:49:25 GMT
{
  "id": "26baf48a-db0f-4884-9b89-820ce7596a6e",
  "name": "Jack"
}
```

```$ curl -i -X GET -H 'Accept: application/json' 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'```

```
HTTP/1.1 200 
Content-Type: application/json
Transfer-Encoding: chunked
Date: Fri, 11 Feb 2022 15:49:58 GMT
{
  "id": "26baf48a-db0f-4884-9b89-820ce7596a6e",
  "name": "Jack"
}
```

```$ curl -i -X DELETE 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'```
```
HTTP/1.1 204 
Date: Fri, 11 Feb 2022 15:50:34 GMT
```


```$ curl -i -X GET -H 'Accept: application/json' 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'```
```
HTTP/1.1 404 
Content-Length: 0
Date: Fri, 11 Feb 2022 15:50:42 GMT 
```