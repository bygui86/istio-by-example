
# Istio by example

## 1. Clone
```
git clone https://github.com/saturnism/istio-by-example.git
cd istio-by-example/
```

## 2. Build & Docker
```
export DOCKER_REGISTRY_PASSWORD=...
export DOCKER_REGISTRY_USER=...
echo "$DOCKER_REGISTRY_PASSWORD" | docker login -u "$DOCKER_REGISTRY_USER" --password-stdin
```
```
cd applications/guestbook-service/
./mvnw clean package
docker build . -t $DOCKER_REGISTRY_USER/guestbook-service:latest
docker push $DOCKER_REGISTRY_USER/guestbook-service:latest
cd -
```
```
cd applications/guestbook-ui/
./mvnw clean package
docker build . -t $DOCKER_REGISTRY_USER/guestbook-ui:1.0
docker push $DOCKER_REGISTRY_USER/guestbook-ui:1.0
cd -
```
`Change background color in 'guestbook-ui/src/main/resources/templates/index.html'`
```
cd applications/guestbook-ui/
./mvnw clean package
docker build . -t $DOCKER_REGISTRY_USER/guestbook-ui:2.0
docker push $DOCKER_REGISTRY_USER/guestbook-ui:2.0
cd -
```
```
cd applications/helloworld-service-java/
./mvnw clean package
docker build . -t $DOCKER_REGISTRY_USER/helloworld-service-java:latest
docker push $DOCKER_REGISTRY_USER/helloworld-service-java:latest
cd -
```
```
cd applications/helloworld-service-go/
docker build . -t $DOCKER_REGISTRY_USER/helloworld-service-go:latest
docker push $DOCKER_REGISTRY_USER/helloworld-service-go:latest
cd -
```

## 3. Kubernetes
### v1
```
kubectl apply -f kubernetes-v1/permissive/
```
### v2
```
kubectl apply -f kubernetes-v2/permissive/
```
