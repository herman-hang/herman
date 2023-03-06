FROM alpine

ENV TZ Asia/Shanghai
WORKDIR /code/herman/api
COPY ./herman /code/herman/api

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --update gcc g++ libc6-compat \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache tzdata \
    && chmod +x /code/herman/api/herman

EXPOSE 8000
CMD ["/code/herman/api/herman","server","--migrate=true"]
