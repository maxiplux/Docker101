#!/bin/bash
# URL calls for Kong initial setup

# setup the AuthServer-linux API
curl -i -X POST --url http://localhost:8001/apis/ \
  --data 'name=AuthServer-linux' \
  --data 'uris=/auth' \
  --data 'upstream_url=http://auth:5001' \
  --data 'strip_uri=true'

# setup the UsersAPI-linux
curl -i -X POST --url http://localhost:8001/apis/ \
    --data 'name=UsersAPI-linux' \
    --data 'uris=/usersapi' \
    --data 'upstream_url=http://users:3001' \
    --data 'strip_uri=true'

# setup the PostsAPI-linux
curl -i -X POST --url http://localhost:8001/apis/ \
    --data 'name=PostsAPI-linux' \
    --data 'uris=/postsapi' \
    --data 'upstream_url=http://posts:3002' \
    --data 'strip_uri=true'

# setup the NwtAPI-linux
curl -i -X POST --url http://localhost:8001/apis/ \
    --data 'name=NetworkAPI-linux' \
    --data 'uris=/ntwapi' \
    --data 'upstream_url=http://ntwrs:3003' \
    --data 'strip_uri=true'

# setup the ImagesAPI-linux
curl -i -X POST --url http://localhost:8001/apis/ \
    --data 'name=ImagesAPI-linux' \
    --data 'uris=/imagesapi' \
    --data 'upstream_url=http://images:3004' \
    --data 'strip_uri=true'

# Secure the UsersAPI-linux with APIKey
curl -X POST http://localhost:8001/apis/UsersAPI-linux/plugins \
    --data "name=key-auth" \
    --data "config.hide_credentials=true"

# Secure the POstsAPI-linux with APIKey
curl -X POST http://localhost:8001/apis/PostsAPI-linux/plugins \
    --data "name=key-auth" \
    --data "config.hide_credentials=true"

# Secure the UsersAPI-linux with APIKey
curl -X POST http://localhost:8001/apis/NetworkAPI-linux/plugins \
    --data "name=key-auth" \
    --data "config.hide_credentials=true"

# Secure the UsersAPI-linux with APIKey
curl -X POST http://localhost:8001/apis/ImagesAPI-linux/plugins \
    --data "name=key-auth" \
    --data "config.hide_credentials=true"
# EOF!
