FROM golang

WORKDIR /app

COPY basket /app/basket
COPY cmd /app/cmd
COPY internal /app/internal

RUN go mod init firstProject
RUN go mod tidy
RUN go build ./...

ENTRYPOINT ["go", "run", "./..."]
