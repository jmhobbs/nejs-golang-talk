.PHONY: serve

serve:
	go run main.go

all: net-http.js xhr.js main net-http

main: main.go
	go build -o main main.go

net-http.js: net-http.go
	gopherjs build -o net-http.js net-http.go

net-http: net-http.go
	go build -o net-http net-http.go

xhr.js: xhr.go
	gopherjs build -o xhr.js xhr.go
