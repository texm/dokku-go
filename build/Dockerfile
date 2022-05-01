FROM debian:stable-slim
LABEL org.opencontainers.image.description 'A testing environment with dokku already set up.'
RUN  apt-get update \
  && apt-get install -y wget \
  && rm -rf /var/lib/apt/lists/* \
  && wget -nv -O - https://get.docker.com/ | sh \
  && wget -nv -O - https://packagecloud.io/dokku/dokku/gpgkey | apt-key add - \
  && export SOURCE="https://packagecloud.io/dokku/dokku/ubuntu/" \
  && export OS_ID="$(lsb_release -cs 2>/dev/null || echo "bionic")" \
  && echo "bionic focal" | grep -q "$OS_ID" || OS_ID="bionic" \
  && echo "deb $SOURCE $OS_ID main" | tee /etc/apt/sources.list.d/dokku.list \
  && apt-get update \
  && apt-get install dokku \
  && dokku plugin:install-dependencies --core
