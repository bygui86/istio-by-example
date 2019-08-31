FROM openjdk:11-stretch

COPY target/lib/* /app/lib/
COPY target/helloworld-service-0.0.1.jar /app/app.jar

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
