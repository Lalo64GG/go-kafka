# Imagen base de Go
FROM golang:1.20-alpine

# Crea y configura el directorio de trabajo
WORKDIR /app

# Copia el archivo go.mod y go.sum, y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila la aplicación
RUN go build -o app .

# Expone el puerto de tu API
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./app"]
