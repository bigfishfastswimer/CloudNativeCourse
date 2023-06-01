FROM golang:1.20.4 As builder
COPY ./app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o web main.go


FROM scratch
COPY --from=builder /app/web /app/web
ENTRYPOINT [ "./app/web" ]
