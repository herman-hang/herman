FROM alpine

ENV TZ Asia/Shanghai

COPY ./herman /herman

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --update gcc g++ libc6-compat \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache tzdata \
    && chmod +x /herman

EXPOSE 8000
CMD ["/herman","server","--migrate=true"]
