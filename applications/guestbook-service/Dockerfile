FROM openjdk:11-stretch

COPY target/guestbook-service.jar /app/app.jar
COPY health-check.sh /app/health-check.sh
RUN chmod +x /app/health-check.sh

EXPOSE 8080

USER 1001

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
