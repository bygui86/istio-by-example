
# Istio by example

## 1. Clone
```
git clone https://github.com/saturnism/istio-by-example.git
cd istio-by-example/
```

## 2. Build
```
./mvnw clean package
```

## 3. Docker
```
echo "$DOCKER_REGISTRY_PASSWORD" | docker login -u "$DOCKER_REGISTRY_USER" --password-stdin
```
```
cd guestbook-service/
docker build . -t bygui86/guestbook-service:latest
docker push bygui86/guestbook-service:latest
cd -
```
```
cd guestbook-ui/
docker build . -t bygui86/guestbook-ui:1.0
docker push bygui86/guestbook-ui:1.0
cd -
```
`Change background color in 'guestbook-ui/src/main/resources/templates/index.html'`
```
cd guestbook-ui/
mvn clean package
docker build . -t bygui86/guestbook-ui:2.0
docker push bygui86/guestbook-ui:2.0
cd -
```
```
cd helloworld-service-java/
docker build . -t bygui86/helloworld-service-java:latest
docker push bygui86/helloworld-service-java:latest
cd -
```
```
cd helloworld-service-go/
docker build . -t bygui86/helloworld-service-go:latest
docker push bygui86/helloworld-service-go:latest
cd -
```

## 4. Kubernetes
### v1
```
kubectl apply -f kubernetes-v1/
```
### v2
```
kubectl apply -f kubernetes-v2/
```
