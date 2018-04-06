#DevHack
FROM php:7.1-apache
MAINTAINER DevHack

RUN pecl install redis
RUN docker-php-ext-enable redis
COPY web/ /var/www/html/