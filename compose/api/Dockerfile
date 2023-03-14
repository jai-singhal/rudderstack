FROM golang:1.20-alpine

WORKDIR /rudderstack-api

COPY . .

# RUN go mod download github.com/kr/text

RUN go build -o rudderstack && chmod +x rudderstack


CMD ["go", "run", "rudderstack"]
