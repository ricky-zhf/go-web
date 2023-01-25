# go_web

# 1、OverView
搭建一个分布式Go Web项目。重点是用上所学的东西，所以业务逻辑很简单，主要是使用更多的技术。

# 2、常用命令
````
CREATE TABLE `blog_tab` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(64) NOT NULL DEFAULT '',
  `title` varchar(64) NOT NULL DEFAULT '',
  `content` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

docker run -itd --name db -p 3306:3306 --net blog_network -e MYSQL_ROOT_PASSWORD=1234 mysql

docker run -itd --name blog_server1 --net blog_network blog_server

docker run -itd --name user_server1  --net blog_network user_server

docker run -itd --name router_server1 --net blog_network router_server

docker network inspect blog_network

etcdctl --endpoints=etcd_server:2379 get --prefix ''
etcdctl --endpoints=etcd_server:2379 put Server_Register:blog.UserService:192.11.11.11:9099 "1"

/Users/zhouhuaifeng/docker/nginx/html

docker run -itd --name nginx -p 8080:80 -v /Users/zhouhuaifeng/docker/nginx/html:/usr/share/nginx/html nginx

docker run -itd --name nginx -p 8080:80 --net blog_network \
-v /Users/zhouhuaifeng/docker/nginx/html:/usr/share/nginx/html \
-v /Users/zhouhuaifeng/docker/nginx/conf:/etc/nginx \
nginx

````
# 2、技术选型
由上往下：
- NG
- Gin实现的网关
- grpc实现微服务间调用
- MySQL、Redis
- docker、etcd

# 3、网关
- HTTP转GRPC
- 流量转发，根据请求接口的地址

