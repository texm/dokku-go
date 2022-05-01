FROM debian:stable-slim
LABEL org.opencontainers.image.description 'A testing environment with dokku already set up.'
RUN  apt-get update \
  && apt-get install -y wget \
  && rm -rf /var/lib/apt/lists/*
RUN wget https://raw.githubusercontent.com/dokku/dokku/v0.27.1/bootstrap.sh && sudo DOKKU_TAG=v0.27.1 bash bootstrap.sh