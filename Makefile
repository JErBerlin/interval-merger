# Makefile

.PHONY: generate_tests test benchmark build run

generate_tests: 
	# Generate some test files to run the end to end tests and benchmarks
	@test -f ./testdata/low.xs.txt || ./generate_tests.sh low xs
	@test -f ./testdata/low.s.txt || ./generate_tests.sh low s
	@test -f ./testdata/low.l.txt || ./generate_tests.sh low l
	@test -f ./testdata/mid.m.txt || ./generate_tests.sh mid m
	@test -f ./testdata/mid.l.txt || ./generate_tests.sh mid l
	@test -f ./testdata/high.xl.txt || ./generate_tests.sh high xl

test:
	@go test

benchmark:
	@go test -bench .

build:
	@go build

run:
	@./intervals_merger || go run

