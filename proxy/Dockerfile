FROM daocloud.io/nginx

RUN mkdir /wwwroot
COPY nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /etc/nginx/sites-conf
COPY sites/* /etc/nginx/sites-conf/



EXPOSE 80

CMD ["service", "nginx", "start"]
