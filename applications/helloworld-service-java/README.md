
# helloworld-service-java

## Docker
```
docker build . -t $DOCKER_REGISTRY_USER/helloworld-service-java:latest

docker run -ti --rm --name helloworld-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/helloworld-service-java:latest

docker run -d --name helloworld-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/helloworld-service-java:latest
```

## REST APIs
```
curl http://localhost:8080/hello/{name}
```

## Metrics
```
curl http://localhost:8080/actuator
```

## Kubernetes probes
```
curl http://localhost:8080/actuator/health
curl http://localhost:8080/hello/ready
```
