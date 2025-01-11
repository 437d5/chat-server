FROM postgres:17.2-alpine3.21

EXPOSE 5432

COPY ./auth.sql /docker-entrypoint-initdb.d/
# docker run -d -p 5432:5432 auth_db:1