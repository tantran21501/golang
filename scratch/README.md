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


go build: Biên dịch gói hoặc chương trình.
Ví dụ: go build -o myapp main.go sẽ biên dịch file main.go và tạo ra file thực thi myapp.
go clean: Xóa các file build, test và module.
Ví dụ: go clean -i ./... sẽ xóa tất cả các file build, test và module trong dự án.
go doc: Hiển thị documentation.
Ví dụ: go doc fmt.Println sẽ hiển thị documentation cho function Println trong package fmt.
go env: Hiển thị giá trị biến môi trường Go.
Ví dụ: go env GOROOT sẽ hiển thị đường dẫn đến thư mục cài đặt Go.
go fix: Cập nhật code theo các đề xuất.
Ví dụ: go fix ./... sẽ cập nhật tất cả code trong dự án theo các đề xuất.
go fmt: Format code theo chuẩn Go.
Ví dụ: go fmt ./... sẽ format tất cả các file Go trong dự án.
go generate: Tự động tạo code.
Ví dụ: go generate ./... sẽ tự động tạo code cho tất cả các file trong dự án.
go get: Tải và cài đặt package.
Ví dụ: go get github.com/gin-gonic/gin sẽ tải về và cài đặt package Gin.
go install: Biên dịch và cài đặt gói hoặc chương trình.
Ví dụ: go install github.com/user/myapp sẽ biên dịch và cài đặt ứng dụng myapp.
go list: Liệt kê các package.
Ví dụ: go list ./... sẽ liệt kê tất cả các package trong dự án.
go mod: Quản lý các dependency.
go mod tidy: Cập nhật và tải về các dependency cần thiết.
go mod vendor: Tạo thư mục vendor chứa các dependency.
go mod download: Tải về các dependency.
go mod edit: Chỉnh sửa file go.mod.
go mod graph: Hiển thị dependency graph.
go mod init: Khởi tạo file go.mod.
go mod verify: Kiểm tra tính toàn vẹn của các dependency.
go mod why: Hiển thị lý do package được import.
go run: Biên dịch và chạy chương trình.
Ví dụ: go run main.go sẽ biên dịch và chạy file main.go.
go test: Chạy các test case.
Ví dụ: go test ./... sẽ chạy tất cả các test case trong dự án.
go tool: Các công cụ hỗ trợ khác.
Ví dụ: go tool pprof sẽ chạy profiler.
go version: Hiển thị phiên bản Go.
Ví dụ: go version sẽ hiển thị phiên bản Go hiện tại.
go vet: Kiểm tra code có vấn đề.
Ví dụ: go vet ./... sẽ kiểm tra tất cả code trong dự án.