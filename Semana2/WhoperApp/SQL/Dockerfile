# Whoper DataBase
# Build in postgres:9.6.8
FROM postgres:9.6.4
MAINTAINER Eduardo Arango <arangogutierrez@gmail.com>

# Copy the SQL config file
COPY whoper.sql /docker-entrypoint-initdb.d/
VOLUME /var/lib/postgresql/data
EXPOSE 5432

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres"]
