FROM debian:stable-20230522

RUN apt-get update && apt-get install -y nginx

COPY  nginx.conf /etc/nginx/nginx.conf

# EXPOSE 8089

CMD ["nginx", "-g", "daemon off;"]