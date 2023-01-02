FROM golang:1.19-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/lotoweb

FROM alpine

WORKDIR /app

COPY --from=build /out/lotoweb /app

COPY --from=build /src/assets /app/assets

EXPOSE 8080

CMD [ "/app/lotoweb" ]