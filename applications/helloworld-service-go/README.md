
# helloworld-service-go

## Docker
```
docker build . -t $DOCKER_REGISTRY_USER/helloworld-service-go:latest

docker run -ti --rm --name helloworld-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/helloworld-service-go:latest

docker run -d --name helloworld-service -p 8080:8080 -p 8090:8090 -p 8091:8091 $DOCKER_REGISTRY_USER/helloworld-service-go:latest
```

## REST APIs
```
curl http://localhost:8080/hello/{name}
```

## Metrics
```
curl http://localhost:8090/metrics
```

## Kubernetes probes
```
curl http://localhost:8091/live

curl http://localhost:8091/ready
```
