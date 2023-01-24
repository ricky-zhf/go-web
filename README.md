# go_web

# 1、OverView
搭建一个分布式Go Web项目。重点是用上所学的东西，所以业务逻辑很简单，主要是使用更多的技术。

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
