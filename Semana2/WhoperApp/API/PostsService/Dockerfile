# PostsAPI RestFul alpine
# Build in Go 1.8.3
FROM alpine:3.6
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the GO-Binary from host
COPY bin/PostsAPI-linux /bin/
# Forward request and error logs to docker log collector
RUN mkdir -p /var/log/postsapi/ && \
  echo "PostsAPI RESTful API logs" >> /var/log/postsapi/api.log && \
  ln -sf /dev/stdout /var/log/postsapi/api.log
# Expose the 3002 port for  API calls
EXPOSE 3002

ENTRYPOINT ["PostsAPI-linux","2>","/var/log/postsapi/api.log"]
