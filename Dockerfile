# ใช้ Go official image
FROM golang:1.20-alpine as build

# ตั้ง working directory
WORKDIR /app

# คัดลอก go.mod และ go.sum และทำการ install dependencies
COPY go.mod go.sum ./
RUN go mod download

# คัดลอกไฟล์ .env ลงใน image
COPY .env .env

# คัดลอกโค้ดทั้งหมด
COPY . .

# สร้างแอปพลิเคชัน
RUN go build -o main ./cmd/main.go

# ใช้ image ที่เบากว่าในการรัน
FROM alpine:latest  

# คัดลอกไฟล์จาก build image
WORKDIR /root/
COPY --from=build /app/main .

# คัดลอกไฟล์ .env ไปที่ container
COPY --from=build /app/.env .env

# เปิดพอร์ตที่แอปพลิเคชันจะรัน
EXPOSE 8080

# รันแอปพลิเคชัน
CMD ["./main"]
