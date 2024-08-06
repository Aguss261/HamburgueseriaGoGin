# Usa una imagen base de Go
FROM golang:1.22

# Configura el directorio de trabajo
WORKDIR /app

# Copia los archivos del módulo Go
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el código fuente
COPY src/ .

# Construye la aplicación
RUN go build -o main .

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
