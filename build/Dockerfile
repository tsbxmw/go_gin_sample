FROM ubuntu:18.04

RUN mkdir /config

RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y tzdata && rm -f /etc/localtime && rm -rf /usr/share/zoneinfo/UTC &&  echo "Asia/Shanghai" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata
COPY ./dist/main /bin
RUN mkdir log   && ls /bin/    && chmod +x /bin/main
EXPOSE 9060

ENTRYPOINT ["/bin/main"]