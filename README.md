# BUILD

To build this application you need to have go installed.

Install the dependencies with:

`go get k8s.io/client-go/...`

Build the application with:

`GOOS=linux go build -o ./app .`

Run the application:

`./app`