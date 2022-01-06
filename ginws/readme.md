1、拉取nginx镜像

```shell
docker pull nginx:1.18
```

2、准备数据卷目录

```shell
# mkdir -p /data/docker/nginx/config/vhost
# mkdir -p /data/docker/nginx/logs
# mkdir -p /data/docker/nginx/html
# mkdir -p /data/docker/nginx/ssl
```

3、配置nginx.conf

```shell
# vi /data/docker/nginx/config/nginx.config

user root;
worker_processes 4; 
worker_cpu_affinity 0001 0010 0100 1000;
worker_rlimit_core 768m;
worker_rlimit_nofile 65536;
 
events {
    worker_connections 65535;
    use epoll;
    epoll_events 1024;
}
 
http {
    include mime.types;
    default_type application/octet-stream;
 
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
    '$status $body_bytes_sent "$http_referer" '
    '"$http_user_agent" "$http_x_forwarded_for" $request_time '
    '"$host" "$upstream_addr" "$upstream_status" "$upstream_response_time" '
 
    access_log off;
 
    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    server_names_hash_bucket_size 128;
    client_max_body_size 100m;
    client_body_buffer_size 1024k;
    client_header_timeout 250;
    max_ranges 10;
    send_timeout 450;
    keepalive_timeout 750;
    server_name_in_redirect off;
    server_tokens off;
 
 
    gzip on;
    gzip_buffers 4 16k;
    gzip_comp_level 9;
    gzip_http_version 1.0;
    gzip_min_length 800;
    gzip_proxied any;
    gzip_types text/plain application/x-javascript text/css text/javascript application/x-httpd-php image/jpeg image/gif image/png image/jpg;
    gzip_vary on;
 
    proxy_set_header Connection Keep-Alive;
    proxy_set_header Host $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
 
 
    include /etc/nginx/config.d/*.config;
}
```

4、运行nginx容器

```shell
docker run -d --name nginx-server --rm  -p 80:80 \
-v /data/docker/nginx/config/vhost:/etc/nginx/config.d:rw \
-v /data/docker/nginx/logs:/var/log/nginx:rw \
-v /data/docker/nginx/config/nginx.config:/etc/nginx/nginx.config:rw \
-v /data/docker/nginx/html:/etc/nginx/html:rw \
nginx:1.18
```

原文链接：https://blog.csdn.net/bowei026/article/details/107251334

5、输入命令`sudo docker ps`，查看`nginx`是否已经启动，若已经启动则将前端页面`client.html`文件拷贝至`/data/docker/nginx/html`目录下，并启动后端go程序，分别在浏览器得两个窗口运行http://10.21.53.183/client.html?uid=1&to_uid=2；http://10.21.53.183/client.html?uid=2&to_uid=1，就可以开启聊天模式啦！

