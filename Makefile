# Makefile

.PHONY: generate_tests test benchmark benchmem build run

# default target generates tests if not existing and run those tests
all: generate_tests test

generate_tests: 
	# Generate some test files to run the end to end tests and benchmarks
	@test -f ./testdata/low.xs.txt || ./generate_tests.sh low xs
	@test -f ./testdata/low.s.txt || ./generate_tests.sh low s
	@test -f ./testdata/low.l.txt || ./generate_tests.sh low l
	@test -f ./testdata/mid.m.txt || ./generate_tests.sh mid m
	@test -f ./testdata/mid.l.txt || ./generate_tests.sh mid l
	@test -f ./testdata/high.xl.txt || ./generate_tests.sh high xl

test:
	# Run end to end tests
	@go test

benchmark:
	# Run performance benchmarks
	@go test -bench .

benchmem:
	# Run memory benchmarks
	@go test -bench=. -benchmem

build:
	@go build

run:
	@./interval-merger || go run main.go

