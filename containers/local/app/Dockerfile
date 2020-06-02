FROM golang:1.14

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv ./migrate.linux-amd64 /usr/bin/migrate

RUN apt-get update && apt-get -y install \
    git \
    vim-tiny \
    default-mysql-client \
        --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

