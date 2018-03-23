# Blurb

This repo contains a workaround the the clouddriver issues contained in this post: https://github.com/spinnaker/spinnaker/issues/2039

## Running the solution

A kubernetes cronjob has been setup to delete the clouddriver service every 24 hours. To run the solution:

`kubectl apply -f kubernetes-cronjob.yaml`

The base image is available on `https://hub.docker.com`

`docker pull tensaibankai/clouddriver-pod-delete`

## BUILD

### Application

To build this application you need to have go installed.

Install the dependencies with:

`go get k8s.io/client-go/...`

Build the application with:

`GOOS=linux go build -o ./app .`

Run the application:

`./app`

### Docker image

Build the docker image with:

`docker build -t clouddriver-pod-delete .`