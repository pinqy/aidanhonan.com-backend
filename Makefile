BINARY_NAME=app
COVERAGE_FOLDER=coverage
GO_COVERAGE_FILE=coverage/coverage.out
LCOV_COVERAGE_FILE=coverage/coverage.lcov
 
all: build test_cov_pretty

release: build test_cov
 
build:
	go build -tags netgo -ldflags '-s -w' -o ${BINARY_NAME}

test:
	go test ./...

test_cov:
	mkdir -p ${COVERAGE_FOLDER}
	go test -coverprofile ${GO_COVERAGE_FILE} -v ./...
	go tool cover -func=coverage/coverage.out
	
test_cov_pretty:
	mkdir -p ${COVERAGE_FOLDER}
	go test -coverprofile ${GO_COVERAGE_FILE} -v ./... | sh prettify_test.sh
	gcov2lcov -infile=${GO_COVERAGE_FILE} -outfile=${LCOV_COVERAGE_FILE}
	go tool cover -func=coverage/coverage.out
 
run:
	build
	./${BINARY_NAME}
 
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf ${COVERAGE_FOLDER}