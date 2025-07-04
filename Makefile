BINARY_NAME=app
COVERAGE_FILE=coverage.out
 
all: build test
 
build:
	go build -o ${BINARY_NAME}

test:
	go test ./...
	
test_cov:
	go test -coverprofile ${COVERAGE_FILE} ./...
 
run:
	go build -tags netgo -ldflags '-s -w' -o ${BINARY_NAME}
	./${BINARY_NAME}
 
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f ${COVERAGE_FILE}