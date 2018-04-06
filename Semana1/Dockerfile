# Imagen Base
FROM gcc:4.9

# Datos sobre el propietario de la imagen
MAINTAINER <NOMBRE> <EMAIL>

# Copiar archivos necesarios para correr la aplicacion
COPY . /usr/src/myapp

# Cambiar de directorio (Por defecto se esta en /)
RUN cd /usr/src/myapp && gcc pi.c -lm -o pi &&  cp pi /bin

# Comando a ejecutar al correr el contenedor 
ENTRYPOINT ["pi"]