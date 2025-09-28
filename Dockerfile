# 第一階段：建置環境
# 使用官方的 Go 映像檔作為建置環境
FROM golang:1.24.6-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 並下載依賴
# 這樣可以利用 Docker 的層快取，只有在依賴變更時才重新下載
COPY go.mod go.sum ./
RUN go mod download

# 複製所有原始碼
COPY . .

# 編譯應用程式，並將二進位檔輸出到 /app/bin 目錄
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN go build -o ./bin/server ./cmd/api

# 第二階段：運行環境
# 使用更小的映像檔作為運行環境
FROM alpine:3.21.2

# 設定工作目錄
WORKDIR /app

# 創建非 root 用戶
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 從 builder stage 複製已編譯的二進位檔
COPY --from=builder --chown=appuser:appgroup /app/bin/server .

# 複製任何必要的靜態檔案(例如資料庫遷移檔案)
COPY --from=builder --chown=appuser:appgroup /app/deployments/migrations /app/deployments/migrations

# 切換到非 root 用戶
USER appuser

# 設定容器啟動時執行的命令
ENTRYPOINT ["./server"]
