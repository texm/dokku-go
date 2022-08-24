FROM ubuntu:focal
LABEL org.opencontainers.image.description='A testing environment with dokku already set up.'

RUN apt-get update \
  && apt-get -y install wget openssh-server \
  && wget https://raw.githubusercontent.com/dokku/dokku/v0.28.1/bootstrap.sh \
  && DOKKU_TAG=v0.28.1 DOKKU_DOCKERFILE=true bash bootstrap.sh

RUN service ssh start
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]