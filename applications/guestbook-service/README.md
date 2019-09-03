
# guestbook-ui

## Docker
```
docker build . -t $DOCKER_REGISTRY_USER/guestbook-service:latest

docker run -ti --rm --name guestbook-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/guestbook-service:latest

docker run -d --name guestbook-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/guestbook-service:latest
```

## REST APIs
```
curl http://localhost:8080/
```

## Metrics
```
curl http://localhost:8080/actuator
```

## Kubernetes probes
```
curl http://localhost:8080/actuator/health
curl http://localhost:8080/
```
