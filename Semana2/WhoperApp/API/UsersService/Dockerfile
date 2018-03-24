# UsersAPI RestFul alpine
# Build in Go 1.8.3
FROM alpine:3.6
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the GO-Binary from host
COPY bin/UsersAPI-linux /bin/
# Forward request and error logs to docker log collector
RUN mkdir -p /var/log/usersapi/ && \
  echo "UsersAPI RESTful API logs" >> /var/log/usersapi/api.log && \
  ln -sf /dev/stdout /var/log/usersapi/api.log
# Expose the 3001 port for  API calls
EXPOSE 3001

ENTRYPOINT ["UsersAPI-linux","2>","/var/log/usersapi/api.log"]
