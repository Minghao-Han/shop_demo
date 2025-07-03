# Build stage
FROM golang:1.24-alpine AS builder

# 所有代码放进构建上下文
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# RUN echo $SVC
# 复制整个工程（以便支持 shared 等模块）
COPY . .

# 切换到对应子服务目录
# WORKDIR /build/rpc/item
ARG SVC
WORKDIR /build/$SVC
RUN echo "Working directory is: $PWD"



# 构建输出放到统一路径
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.io,direct go build -o /app/server .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

ENV GIN_MODE=release
EXPOSE 8889

CMD ["./server"]
