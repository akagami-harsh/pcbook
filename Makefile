gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/todo.proto
clean:
	rm pb/*.go
run:
	go run main.go