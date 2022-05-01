FROM debian:stable
LABEL org.opencontainers.image.description 'A testing environment with dokku already set up.'
RUN  apt-get update \
  && apt-get install -y wget \
  && rm -rf /var/lib/apt/lists/* \
  && wget -nv -O - https://get.docker.com/ | sh \
  && wget -nv -O - https://packagecloud.io/dokku/dokku/gpgkey | apt-key add - \
  && echo "deb https://packagecloud.io/dokku/dokku/ubuntu/ bionic main" | tee /etc/apt/sources.list.d/dokku.list \
  && apt-get update \
  && apt-get install -y dokku \
  && echo "dokku dokku/vhost_enable boolean true" | sudo debconf-set-selections \
  && echo "dokku dokku/hostname string test" | sudo debconf-set-selections \
  && echo "dokku dokku/skip_key_file boolean true" | sudo debconf-set-selections \
  && echo "dokku dokku/key_file string ''" | sudo debconf-set-selections \
  && echo "dokku dokku/nginx_enable boolean true" | sudo debconf-set-selections \
  && dokku plugin:install-dependencies --core
