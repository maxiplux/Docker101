# ImagesAPI RestFul alpine
# Build in Go 1.8.3
FROM alpine:3.6
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the GO-Binary from host
COPY bin/ImagesAPI-linux /bin/
# Forward request and error logs to docker log collector
RUN mkdir -p /var/log/imagesapi/ && \
  echo "ImagesAPI RESTful API logs" >> /var/log/imagesapi/api.log && \
  ln -sf /dev/stdout /var/log/imagesapi/api.log
# Expose the 3004 port for  API calls
EXPOSE 3004

ENTRYPOINT ["ImagesAPI-linux","2>","/var/log/imagesapi/api.log"]
