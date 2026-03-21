build:
	go build -v -o bin/invoice_gen *.go

run: build
	./bin/invoice_gen

sqlc:
	sqlc generate

clear-cache:
	rm -rf ~/.local/share/parmaan-patr

clear-invoice:
	rm ./invoices/*
