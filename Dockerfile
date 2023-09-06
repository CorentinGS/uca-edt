FROM golang:1.21-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache upx

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go get -d -v
RUN go build -ldflags="-s -w" -o /app/uca-edt .
RUN upx /app/uca-edt

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64
RUN chmod +x /usr/local/bin/dumb-init

FROM gcr.io/distroless/static:nonroot AS production

COPY --from=builder /usr/share/zoneinfo/Europe/Paris /usr/share/zoneinfo/Europe/Paris
ENV TZ Europe/Paris

WORKDIR /app

COPY --from=builder /app/uca-edt /app/uca-edt
COPY --from=builder /build/.env /app/.env
COPY --from=builder  /usr/local/bin/dumb-init /usr/bin/dumb-init



EXPOSE 3000

USER nonroot

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/app/uca-edt"]