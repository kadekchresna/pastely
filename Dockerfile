FROM golang:1.23.4-alpine3.21 as build

WORKDIR /build

RUN apk add --update make

COPY . .

RUN go mod download

RUN make build


FROM golang:1.22-alpine

COPY --from=build /build/build/pastely /usr/local/bin/pastely
COPY --from=build /build/.env /usr/local/bin/.env

RUN chmod +x /usr/local/bin/pastely

ENTRYPOINT ["pastely"]

CMD [ "web" ]
