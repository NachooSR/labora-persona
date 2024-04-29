FROM golang:1.22 as builder

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar el módulo Go y los archivos de suma
COPY go.mod ./
COPY go.sum ./

# Descargar las dependencias del módulo Go
RUN go mod download

# Copiar el código fuente del proyecto en el contenedor
COPY . .

# Compilar la aplicación para un ejecutable
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Usar una imagen alpine para la etapa de ejecución por su tamaño reducido
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el ejecutable de la etapa de construcción al contenedor final
COPY --from=builder /app/main .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]