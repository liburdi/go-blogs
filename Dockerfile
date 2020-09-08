FROM cent:v2

EXPOSE 8080


RUN
if [ ! -x "blog.exe"]; then \
   ADD blog / \
fi

RUN
if [ ! -x "blog.exe"]; then \
   ADD blog.exe / \
fi


RUN mkdir -p /static

ADD static /static

RUN mkdir -p /uploads

ADD uploads /uploads

CMD [ "blog"]