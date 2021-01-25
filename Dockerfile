FROM cent:v2

EXPOSE 8080

ADD golangschool /

RUN mkdir -p /static

ADD static /static

RUN mkdir -p /uploads

ADD uploads /uploads

CMD [ "./golangschool"]