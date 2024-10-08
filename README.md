# LinkMe - 开源论坛项目(微服务)

![LinkMe](https://socialify.git.ci/GoSimplicity/LinkMe-microservices/image?description=1&font=Source%20Code%20Pro&forks=1&issues=1&language=1&logo=https%3A%2F%2Fgithub.com%2Fwangzijian2002%2FLinkMe%2Fassets%2F71474660%2F22ef2063-ab82-481f-898f-29d95fa70236&name=1&pattern=Solid&pulls=1&stargazers=1&theme=Dark)

## 项目简介
LinkMe 是一个使用 Go 语言开发的论坛项目。它旨在为用户提供一个简洁、高效、可扩展的在线交流平台。本项目使用DDD领域设计理念，采用微服务架构，使得添加新功能或进行定制化修改变得非常容易。LinkMe 支持多种数据库后端，并且可以通过 Kubernetes 进行部署。

## 功能特性
- 用户注册、登录、注销
- 发布帖子、评论、点赞
- 用户个人资料编辑
- 论坛版块管理
- 多种数据库支持
- Kubernetes 一键部署
- RESTful API 设计
- 前后端分离架构

## 技术栈
- Go 语言
- Kratos 微服务框架
- Wire 依赖注入
- Kubernetes 集群管理
- MySQL 数据库
- Redis 缓存数据库
- MongoDB 文档数据库
- Kafka 消息队列
- Prometheus 监控
- ELK 日志收集
- ElasticSearch 搜索引擎
- Docker 容器化
- 随项目进度技术栈实时更新...

## 目录结构
```
.
├── api             # proto文件
├── app             # 核心业务逻辑
├── deploy          # 集成部署
├── doc             # 项目文档
├── go.mod          # Go模块定义文件
├── go.sum          # Go模块依赖校验和
├── LICENSE         # 开源许可证书
├── README.md       # 项目自述文件
```

## 如何贡献
我们欢迎任何形式的贡献，包括但不限于：
- 提交代码（Pull Requests）
- 报告问题（Issues）
- 文档改进
- 功能建议

请确保在贡献代码之前阅读了我们的[贡献指南](#贡献指南)。

## 贡献指南
- Fork 本仓库
- 创建您的特性分支 (`git checkout -b my-new-feature`)
- 提交您的改动 (`git commit -am 'Add some feature'`)
- 将您的分支推送到 GitHub (`git push origin my-new-feature`)
- 创建一个 Pull Request

## 开始使用
### 克隆项目
```bash
git clone git@github.com:GoSimplicity/LinkMe-microservices.git
```

### 安装依赖
```bash
go mod tidy
```

### 创建数据库
你可以选择手动创建数据库，或者使用提供的 Kubernetes YAML 文件自动创建。

#### 手动创建
```bash
# 进入项目目录中的scripts目录下
cd scripts
# windows用户执行下面文件
windows_init.bat
# linux用户及mac用户执行下面文件
chmod +x linux_init.sh && ./linux_init.sh
```

#### 使用 docker-compose 启动中间件(推荐)
```bash
cd yaml/ && docker-compose up -d
```

#### 使用 Kubernetes YAML 文件 启动中间件
```bash
kubectl apply -f yaml/  # 需要有k8s环境
```

### 使用 Wire 进行依赖注入
```bash
go install github.com/google/wire/cmd/wire@latest
go generate ./...
```

### 启动项目
```bash
kratos run
```

### 项目超级管理员账号
```bash
admin/admin
```

## 许可证
本项目使用 MIT 许可证，详情请见 [LICENSE](./LICENSE) 文件。

## 联系方式
- Email: [wzijian62@gmail.com](mailto:wzijian62@gmail.com)

## 致谢
感谢所有为本项目做出贡献的人！

---
欢迎来到 LinkMe，让我们一起构建更好的论坛社区！

---
