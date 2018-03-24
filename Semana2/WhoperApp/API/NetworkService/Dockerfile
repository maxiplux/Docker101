# NetworkAPI RestFul alpine
# Build in Go 1.8.3
FROM alpine:3.6
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the GO-Binary from host
COPY bin/NetworkAPI-linux /bin/
# Forward request and error logs to docker log collector
RUN mkdir -p /var/log/networkapi/ && \
  echo "NetworkAPI RESTful API logs" >> /var/log/networkapi/api.log && \
  ln -sf /dev/stdout /var/log/networkapi/api.log
# Expose the 3003 port for  API calls
EXPOSE 3003

ENTRYPOINT ["NetworkAPI-linux","2>","/var/log/networkapi/api.log"]
