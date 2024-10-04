# M-Golang-rest-mysql: Minimalist Golang REST MySQL

The aim here is to make a small, simple, and easy to plug REST service with access to mysql for
data.

The focus is just on doing CRUD operation, but without exact map with database.

From the point of view of an application, in fact a REST API is a layer of abstraction.

Whenever someone use it as a direct map between db table and the exposed entry point, she/he
is giving up the opportunity to abstract from the other level.

## Real example

What I am going to define here, as example, is a model that does not just map table, so
I propose:

- (lookup) a `JOIN`-query for lookup
- (update) a multiple update for update, in a transaction
- (lookup) a request to outer service for data retrieval and integration of mysql data

This almost all cases that make sense as example of microservice in a microservice environment,
apart from consuming message from a message broker.

## Folder structure

`model/` contains the module that access the db
`config/` contains module first init()'ed when service go up and running, so the db connection


## Other resources (and inspirations)


https://github.com/boantp/go-mysql-rest-api
https://github.com/Kivanc10/golang-rest-api-with-mysql
https://github.com/Fikri1234/CRUD_REST_MUX_MySQL
https://github.com/phathv/simple-restful-api-with-golang-mysql
https://github.com/wisnuuakbr/simple-rest-api-go
https://github.com/kelvins/GoApiTutorial
https://github.com/omjigupta/golang-gin-restapi-mysql