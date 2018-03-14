# Imagen Base
FROM gcc:4.9

# Datos sobre el propietario de la imagen
MAINTAINER <NOMBRE> <EMAIL>

# Copiar archivos necesarios para correr la aplicacion
COPY . /usr/src/myapp

# Cambiar de directorio (Por defecto se esta en /)
WORKDIR /usr/src/myapp

# Se Compila el binario de la aplicacion
RUN gcc pi.c -lm -o pi
RUN cp pi /bin

# Comando a ejecutar al correr el contenedor 
ENTRYPOINT ["pi"]