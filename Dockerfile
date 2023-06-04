FROM golang:1.19.10-buster AS gobuilder

# use static link build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/local/src/repo
COPY . /usr/local/src/repo
RUN go build ./cmd/db-setup
RUN go build ./cmd/excel-importer

FROM gcr.io/distroless/static-debian11:latest

COPY --from=gobuilder /usr/local/src/repo/db-setup /usr/bin/ 
