FROM golang:1.15.2-alpine AS build
#
ADD ./ /opt/build/golang
WORKDIR /opt/build/golang
RUN go install .

RUN go build -o /docker_backend


CMD [ "/docker_backend" ]
#--docker run -d --name url_shortener -p 5432:5432 url_shortener_db

#FROM postgres AS db
#ENV POSTGRES_PASSWORD postgres
#ENV POSTGRES_DB postgres
#COPY /db/db.sql /docker-entrypoint-initdb.d/
#
