# GIN-APP-FRAME

## 背景
基于Gin框架实现了项目的基础骨架，项目的目录结构参照了golang-standards/project-layout进行实现。同时封装了包括配置读取、日志处理、gorm、验证器、接口返回数据处理等功能。

## 使用说明

### 获取代码
```sh
git clone https://github.com/SuperCoCa/gin-app-frame.git
```

### 启动
```sh
# 进入项目目录
cd gin-app-frame

# 修改配置
vi ./configs/config.yaml

# 启动服务
go run cmd/app/main.go
```