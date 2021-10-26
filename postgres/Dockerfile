FROM postgres:13-buster

COPY ./internal/persistence/schema/ /docker-entrypoint-initdb.d/

EXPOSE 5432