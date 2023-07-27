FROM golang:1.21rc3-alpine3.17 AS base

WORKDIR /app

COPY /main.go /app/main.go


FROM base AS build
RUN go build -o /server

FROM build AS prod
COPY --from=base /server /server
CMD [ "/server" ]

FROM base AS dev
# ENTRYPOINT [ "go" ]
WORKDIR /app

CMD [ "go", "run", "main.go" ]
