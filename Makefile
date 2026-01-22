.PHONY: test coverage lint run

test: lint coverage
	go test -cover -count=1 -race ./...

coverage:
	rm -rf ./coverage
	mkdir -p ./coverage
	go test -coverprofile=./coverage/coverage.out ./...
	go tool cover -func=./coverage/coverage.out
	go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	go tool cover -func=./coverage/coverage.out > ./coverage/func_coverage.txt
	@echo "Coverage report generated at ./coverage/coverage.html"
	@echo "Function coverage report saved to ./coverage/func_coverage.txt"

coverage-open: coverage
	@open ./coverage/coverage.html

lint:
	golangci-lint run ./...

run:
	mkdir -p ./db
	go run ./cmd/api -database-url "./db/myapp.db"

doc:
	go install golang.org/x/tools/cmd/godoc@latest
	@echo "Documentation running at http://localhost:6060/pkg/github.com/markbates/myapp"
	godoc -http=:6060