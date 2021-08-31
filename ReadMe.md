# Todo List App in Go

Made using go, gorm and gorilla-mux. Uses mysql as the database. Mysql version is 8.0.26 for macos11.3 on x86_64 (Homebrew). Go version used is go1.17 darwin/amd64.

## Steps to Configure

1. Run `go mod download` to download all the required packages.
2. `conf.yaml` file has all the configurations required to connect to database, port on which server should run as well as token for request verification.
3. `dbName` in **conf.yaml** is the schema name in which table for this server should be created. Table creation is done using migration. Please make a schema in your mysql and set its name in `dbName`
4. Other configurations related to database should also be provided in **conf.yaml** as they might differ from what I used.
5. If `AuthToken` is changed inside the **conf.yaml**, then please also change the `bearerToken` global variable in the postman collection.
6. Similarly if `port` is changed in **conf.yaml**, please change the `baseUrl` global variable in the postman collection accordingly.

## Steps to run

1. After successfully configuration, run `go build -o goTodoList` inside this project folder. This creates the executable named `goTodoList` inside this folder.
2. Run `./goTodoList` for starting the server at port configured in **conf.yaml**.

## Steps to verify

1. After successfully configuration, verfication can be done using postman.
2. Import the collection in postman.
3. If changes were made in **conf.yaml**, please make sure to take care of steps **5** and **6** in **Steps to Configure**.
4. Various request and response examples will help you to verify the submission.
5. There is a request titled `Missing Authorzation`, its examples shows the response if Auth header is not given as well as if wrong token is given. You can change Api url if you want to check other requests too.
6. Database constraints and api requests structure (along with parameters) are explained in necessary comments. **Please read those comments** to understand and run requests in postman accordingly. Only `POST`, `PUT` and filter requests need some explaination, so only those are explained in detail
7. `POST` and `PUT` requests will be checked to follow constraints specified in the model `TodoItem`. If they follow the constraints request will be successful else an error will be returned. The request response for requests that don't follow constraints is similar in `POST` and `PUT` are same so I have not added all failed request for `PUT`. I have added failed requests for many cases with their responses.
