# Usa una imagen base de Go
FROM golang:1.22

# Configura el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de definición de módulos
COPY src/go.mod src/go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el código fuente al contenedor
COPY src/ .

# Compila el proyecto
RUN go build -o main .

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8080

# Define el comando por defecto para ejecutar la aplicación
CMD ["./main"]
