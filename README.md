# NGINX BASICS

This repo is a simple boilerplate to study the Nginx basic utilities when I was reading [this arcticle.](https://dev.to/mauricioabreu/uma-introducao-ao-nginx-1jdg)

Here we have a simple Golang app with docker and a simple nginx config to study about load balancing and cache.

### Dependencies
- Docker
- Golang
- Nginx

### How to run

```
docker-compose up
```

### Load Balancer
To learn about load balancer, this project start up two containers that uses the same app build to test in nginx upstream servers.

### Proxy Cache
To learn about cache requests in nginx I used properties inside location to configure.

### Issues to solve
- [ ] When changes a golang code the container doesn't refresh with the new changes
(Temporary solution) run to clean cache:
```
docker-compose build --no-cache
```