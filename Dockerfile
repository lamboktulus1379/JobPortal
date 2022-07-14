# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
ENV GOOS=linux
RUN mkdir -p /Jobportal
WORKDIR /src
RUN export GOPATH=/Jobportal
ADD /Jobportal/ /Jobportal
WORKDIR /Jobportal
RUN go build -o main .
CMD ["/Jobportal/main"]