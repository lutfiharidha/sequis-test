FROM golang:1.18.0-bullseye as builder

WORKDIR /app
COPY . .
RUN make build

FROM redhat/ubi8-micro:8.5-437

WORKDIR /app
COPY --from=builder /app/bin/app ./
COPY --from=builder /app/.env ./
ENV GODEBUG madvdontneed=1

RUN mv ./app /usr/local/bin/app
CMD ["/usr/local/bin/app"]