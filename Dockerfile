FROM golang:1.18.4-alpine3.15 AS build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update make
RUN make

FROM alpine:3.14
EXPOSE 8880
COPY --from=build /app/bin/orch-rest-template /orch-rest-template
COPY .env .env
CMD [ "/orch-rest-template"]
