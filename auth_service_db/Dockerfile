FROM postgres:17.2-alpine3.21

ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=auth

EXPOSE 5432

COPY ./auth.sql /docker-entrypoint-initdb.d/
# docker run -d -p 5432:5432 4c11c1786a65