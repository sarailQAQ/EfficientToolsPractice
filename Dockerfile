FROM alpine
LABEL maintainer="sarail@qq.com"
WORKDIR /app
COPY . /app
CMD /app/main
