# Go Kafka Project

Este es un proyecto en Go que utiliza Kafka para producir y consumir mensajes, además de realizar operaciones sobre una base de datos MySQL. Este ejemplo utiliza el paquete `confluent-kafka-go` para interactuar con Kafka y maneja la inserción de datos en una base de datos MySQL.

## Requisitos

Para ejecutar este proyecto, necesitas tener las siguientes herramientas instaladas:

- **Go 1.123+**: Asegúrate de tener Go instalado en tu sistema. Puedes verificar si lo tienes ejecutando `go version`.
- **Docker**: Necesitarás Docker para ejecutar Zookeeper y Kafka en contenedores.
- **MySQL**: Este proyecto interactúa con una base de datos MySQL.

## Instalación

1. Clona este repositorio en tu máquina local:

    ```bash
    git clone https://github.com/Lalo64GG/go-kafka.git
    cd go-kafka
    ```

2. Asegúrate de tener Docker instalado y ejecuta Kafka y Zookeeper en contenedores Docker utilizando `docker-compose`:

    ```bash
    docker-compose up -d
    ```

   Esto levantará dos contenedores: uno para **Zookeeper** y otro para **Kafka**.

3. Ejecuta las migraciones para crear la tabla `temperature` en MySQL:

    ```bash
    go run ./src/config/migrates/migrates.go
    ```

    Esto creará la tabla `temperature` en la base de datos MySQL. Si la tabla ya existe, será ignorada.

4. Instala las dependencias de Go:

    ```bash
    go mod tidy
    ```

5. Configura las variables de entorno en tu `.env` o directamente en el código para que coincidan con la configuración de tu Kafka y MySQL.

## Estructura del Proyecto

```text
go-kafka/
├── src/
│   ├── config/
│   │   ├── migrates/
│   │   └── config.go
│   ├── controllers/
│   ├── middlewares/
│   ├── models/
│   ├── routers/
│   └── service/
├── docker-compose.yml
├── README.md
└── go.mod
