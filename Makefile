compile:
	protoc api/v1/*.proto \
                --go_out=. \
                --go-grpc_out=. \
                --go_opt=paths=source_relative \
                --go-grpc_opt=paths=source_relative \
                --proto_path=.
test_cover:
	go test -coverprofile=coverage.out ./...

test:
	go test -v ./...

coverage:
	go tool cover -html=coverage.out

clean:
	rm -f coverage.out