# To do api

> Simple Go to do rest api

## Basic functionalities of to do rest api:

```bash
Has the following apis:
User
1. [POST]   /signup -> to create a new user
2. [POST]   /login -> to authenticate a user

Todo
1. [POST]   /todos -> to create a new todo                  [requires jwt]
2. [GET]    /todos -> to get a list of all todos            [requires jwt]
3. [GET]    /todos/:id  -> to get a todo by id              [requires jwt]
4. [PUT]    /todos/:id  -> To update a todo by id           [requires jwt] [only creator can execute]
5. [DELETE] /todos/:id  -> To delete a todo by id           [requires jwt] [only creator can execute]

User - todo
1. [POST] /todos/id/link -> to link a user* to a todo**     [requires jwt]
2. [POST] /todos/id/unlink -> to unlink a user* to a todo** [requires jwt]

* means user id taken from jwt
** means todo id is the id specified in the url
```

## Running the project locally:

```bash
1. Download and unzip project.<br>
   Open the project in a code editor e.g. Visual Studio Code.<br>
   Requires Go to be installed on your development setup.
   Installation instructions for Go can be found here: https://go.dev/doc/install

2. If Visual Studio Code is used, install the REST Client to use the apis in the api-test folder.
api-test folder contains the sample requests used to test the apis.
Otherwise, postman can be used too.

3. Using sqlite as a sql database for development.

4.
a) To run the api, type in terminal after opening the root folder:
    go install
    go run . (Ctrl+C to stop)

b) To build the app , type in terminal after opening the root folder:
    go build
    ./todo-app.exe (Ctrl+C to stop, assuming windows OS)
```
