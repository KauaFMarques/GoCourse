# Usa a imagem oficial do Go para compilar
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Compila o binário (ajuste o caminho se seu main.go estiver em /cmd)
RUN go build -o main ./cmd/main.go

# Imagem final mais leve para rodar o app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Porta que sua API usa
EXPOSE 8080

# Comando para rodar
CMD ["./main"]
