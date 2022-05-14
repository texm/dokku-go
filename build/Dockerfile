FROM debian:stable
LABEL org.opencontainers.image.description 'A testing environment with dokku already set up.'

RUN apt-get update \
  && apt-get -y install wget \
  && wget https://raw.githubusercontent.com/dokku/dokku/v0.27.1/bootstrap.sh \
  && DOKKU_TAG=v0.27.1 DOKKU_DOCKERFILE=true bash bootstrap.sh
