FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o api  main.go


FROM scratch 

COPY --from=build /app/api .

EXPOSE 5600

CMD ["./api"]
