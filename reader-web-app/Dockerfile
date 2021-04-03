FROM golang:1.16 as builder
ENV USER=appuser \
    UID=10001 \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR /app
COPY cmd /app/cmd/
COPY pkg /app/pkg/
COPY go.* /app/
RUN cd /app \
    && go build -a -o reader $PWD/cmd/reader

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
WORKDIR /app
USER ${USER}:${USER}
COPY --from=builder /app/reader .
ENV USER=appuser \
    APP_PORT=8080
EXPOSE ${APP_PORT}
CMD ["/app/reader"]
