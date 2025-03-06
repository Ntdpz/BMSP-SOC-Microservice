# ใช้ Golang image สำหรับ build
FROM golang:1.24-alpine as builder

# กำหนด working directory
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum
COPY go.mod go.sum ./

# ดาวน์โหลด dependencies
RUN go mod tidy

# คัดลอกซอร์สโค้ดทั้งหมด
COPY . .

# สร้าง binary
RUN go build -ldflags="-w -s" -o main



FROM scratch


WORKDIR /app


COPY --from=builder /app/main /main

CMD ["/main"]

EXPOSE 8090

