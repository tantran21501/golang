
Config sqlc.yaml: https://docs.sqlc.dev/en/stable/reference/config.html#go

Install sqlc:
    - brew install sqlc
Generate sql: 
    /* At folder scratch command line */
    sqlc generate
Install Goose:
 - go install github.com/pressly/goose/v3/cmd/goose@latest
 - brew install goose
At schema folder:
    /* Run file 001_users.sql */
 - goose postgres postgresql://root:root@localhost:15432/database up 
 - goose postgres postgresql://root:root@localhost:15432/database down
 - goose postgres postgresql://root:root@localhost:15432/database reset

Go command:
go build : Compile file .go to file run
go run <filename>: Run file .go
go get <github.com/joho/godotenv> : Get and install package godotenv from github 
go mod tidy : Add library/dependencies into vendor folder in project
go mod vendor: Use library in folder vendor for project
go fmt ./... : format code all project
go doc <fmt.Println>: Show document for function Println in fmt lib