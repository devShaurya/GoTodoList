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
