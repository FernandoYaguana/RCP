# Usa una imagen base con Go
FROM golang:1.20

# Define el directorio de trabajo
WORKDIR /app

# Copia el código del servidor y cliente
COPY . .

# Compila el código
RUN go mod init app && go mod tidy && go build -o app

# Expone el puerto
EXPOSE 1234

# Comando para ejecutar el servidor
CMD ["./app"]