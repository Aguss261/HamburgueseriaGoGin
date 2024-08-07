# Usa una imagen base de Go
FROM golang:1.22

# Configura el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de definición de módulos
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el código fuente al contenedor
COPY src/ ./src

WORKDIR /app/src

# Compila el proyecto
RUN go build -o main .


# Define el comando por defecto para ejecutar la aplicación
CMD ["./main"]
