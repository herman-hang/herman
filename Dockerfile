FROM alpine

ENV TZ Asia/Shanghai
WORKDIR /www/herman-api
COPY . /www/herman-api

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --update gcc g++ libc6-compat \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache tzdata \
    && chmod +x /www/herman-api/herman

EXPOSE 8000
CMD ["/www/herman-api/herman","server","--host=0.0.0.0","--port=8000","--migrate=true"]
