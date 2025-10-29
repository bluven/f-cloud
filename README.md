一个用来学习go-zero的仓库

## 依赖

- MySQL
- Docker
- Apisix
- Redis
- Jaeger
- Consul
- Etcd
- modd

## 启动

1. MySQL没有用Docker启动，需要手动安装，并初始化数据，数据在deploy目录下；
2. 启动其他依赖：`docker compose up -d`;
3. 启动所有服务:`modd`;
4. 注册路由：`deploy/add-routes.sh`.

## 服务发现

服务发现有两点：

1. RPC服务的服务发现，这是go-zero自带的，只用了etcd；
2. API服务的注册和发现，这里使用了consul来注册服务，apisix的consul插件来自动发现服务；一个用来学习go-zero的仓库

## API测试

因为没有做前端，所以API的测试都是通过vscode rest client来进行的，所有的文件都在`http`目录下.