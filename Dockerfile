FROM cent:v2

EXPOSE 8080

ADD blog /

RUN mkdir -p /static

ADD static /static

RUN mkdir -p /uploads

ADD uploads /uploads

CMD [ "\blog"]