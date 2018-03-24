# Whoper AuthServer alpine
# Build in Go 1.8.3
FROM alpine:3.6
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the GO-Binary from host
COPY bin/AuthServer-linux /bin/
# Copy the SignKeys from host
COPY keys/app.rsa /etc/keys/
# Forward request and error logs to docker log collector
RUN mkdir -p /var/log/AuthServer/ && \
  echo "Whoper AuthServer-linux logs" >> /var/log/AuthServer/AuthServer.log && \
  ln -sf /dev/stdout /var/log/AuthServer/AuthServer.log
# Expose the 5001 port for  API calls
EXPOSE 5001

ENTRYPOINT ["AuthServer-linux","2>","/var/log/AuthServer/AuthServer.log"]
