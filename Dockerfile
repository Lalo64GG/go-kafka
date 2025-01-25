# Usa la imagen base de Go con Ubuntu
FROM golang:1.23.2

# Instala gcc y las herramientas necesarias para Cgo
RUN apt-get update && apt-get install -y \
    build-essential \
    libssl-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Establece la variable de entorno para habilitar CGO
ENV CGO_ENABLED=1

# Establece el directorio de trabajo
WORKDIR /go/src/myapp

# Copia el código fuente al contenedor
COPY . .

# Descarga las dependencias del proyecto
RUN go mod tidy

# Compila la aplicación Go
RUN go build -o /go-app

# Expone el puerto
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["/go-app"]
