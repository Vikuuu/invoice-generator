build:
	go build -v -o bin/invoice_gen *.go

run: build
	./bin/invoice_gen
