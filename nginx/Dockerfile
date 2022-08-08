FROM nginx:1.19
RUN apt-get update && \
    apt-get install -y apt-utils \
    locales && \
    sed -i -e 's/# ja_JP.UTF-8/ja_JP.UTF-8/g' /etc/locale.gen && \
    locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_TIME C
ENV TZ Asia/Tokyo
RUN echo "${TZ}" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata
COPY conf/default.conf /etc/nginx/conf.d/default.conf
COPY ssl_keys/localhost.pem /etc/pki/tls/certs/example_com_combined.crt
COPY ssl_keys/localhost-key.pem /etc/pki/tls/private/example_com.key