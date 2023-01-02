FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download && \
    go build -o /out/lotoweb

FROM alpine

WORKDIR /app

COPY --from=build /out/lotoweb /app

COPY --from=build /app /app

EXPOSE 8080

CMD [ "/app/lotoweb" ]