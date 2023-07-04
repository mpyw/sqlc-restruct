FROM golang:1.20-bullseye

RUN apt-get update \
  && apt-get install -y postgresql-client \
  && go install -v github.com/rubenv/sql-migrate/sql-migrate@v1.5.1

ENTRYPOINT ["sql-migrate"]
CMD ["up"]
