# Docker file for DEPROY
FROM golang:onbuild

RUN go build
EXPOSE 8080
CMD regisys
