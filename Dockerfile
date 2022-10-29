FROM alpine:3.13.1
RUN echo "http://mirrors.aliyun.com/alpine/v3.13/main/" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.13/community/" >> /etc/apk/repositories
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone \
  && apk del tzdata

WORKDIR /gmqttd
COPY build/ .

RUN ls
EXPOSE 1883 8883 8082 8083 8084
ENV PATH=$PATH:/gmqttd
RUN chmod +x gmqttd

ENTRYPOINT ["gmqttd","start","-c","gmqtt.yml"]
