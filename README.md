# zopsmart-assignment

## Requriements

- Must have GoLang Compiler/ToolChain installed

## How was the project made?

- Created a github repository and cloned it to local system.
- Initialized the "Go Module" using `go mod init github.com/moulikchaturvedi/zopsmart-assignment`. This simplifies the remote downloading of packages.
- Created the file `main.go`, with basic configuration (including GoFR).
- For necessary libraries, `go mod tidy`.
- `docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=password -p 2001:3306 -d mysql/mysql-server` to start the mysql docker server.
- Create a database and table on the docker server. `myysql -u root -p password` `CREATE DATABASE Students;` and `CREATE TABLE students (id int auto_increment, name varchar(255), class varchar(255), PRIMARY KEY (id));`. Created Dummy Data, `INSERT INTO students (name, class) VALUES ("moulik","fourth");`

## To run

- To run can use the command `go run main.go`.
- We can also the command `go build` which will compile the code in an executable program. Then we can run the executable file.
