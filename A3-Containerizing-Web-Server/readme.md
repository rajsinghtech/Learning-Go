# Containerizing Go Web Server
This is a simple Go web server that serves a home page and a about page. To do this we utilize the Golang image on docker hub.

## Explanation
Using the Golang container on DockerHub we build our package 'go build -o webserver .'. We then create a `docker-compose.yaml` to build and host the container locally.
We use the alpline image specifically to reduce the size of the container significantly.
 
Initialize Project
```
docker compose up
```


[Golang Image](https://hub.docker.com/_/golang)
