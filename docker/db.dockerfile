FROM postgres:alpine

COPY /db/initdb.sql /docker-entrypoint-initdb.d/initdb.sql

EXPOSE 5432