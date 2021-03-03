FROM centos

EXPOSE 8083

ADD golangschool /

RUN mkdir -p /static

ADD static /static

RUN mkdir -p /uploads

ADD uploads /uploads

RUN mkdir -p /templates

ADD templates /templates

CMD [ "./golangschool"]