# 5z7Game
实现go的五子棋游戏

## 背景
- 如何使用go实现游戏项目开发
- 要求扩展性强
- 要求有高并发处理能力


## 概念说明
五子棋游戏的元素包含用户，棋盘，棋局
- 棋盘：空间容器
- 棋局：棋子的状态
- 用户：在棋盘上操控棋子对棋局进行变更

## 系统设计
游戏就是为了输赢，所以我们发现白方胜或者黑方胜，游戏就算结束
### 流程设计
![流程图](http://processon.com/chart_image/5e990887e0b34d6feab3f412.png)
核心功能点就是创建房间和判断输赢。

### 功能模块设计
- 用户模块
    - 用户登录与注册设计
    实现用户的登录与注册
    - 用户分数与段位设计
    实现用户计分与段位变更
- 棋局模块
    - 棋盘开局
    黑白双方就位，棋局开始
    
    - 棋局记录
    记录用户的落子记录，可以支持悔棋。
    
    - 棋局计输赢
    计算哪方胜出，胜者得分，负者扣分。
    
    - AI对战(有兴趣的做)
   
### 数据库设计
参考`db.sql`

### 缓存设计
暂无
## 项目分层
```
├─.idea
├─config
├─http
│  ├─handler
│  ├─response
│  └─route
├─pkg
│  ├─app
│  ├─constant
│  ├─dto
│  │  ├─request
│  │  └─response
│  ├─enum
│  ├─model
│  │  ├─dao
│  │  ├─data
│  │  └─entity
│  ├─service
│  └─utils
└─task
└─main.go
```

- main.go 项目入口
- http 接口模块
    - handler 接口处理模块
    - response 响应
    - router 路由
- pkg 业务模块
    - app 核心服务(logger,database,redis等)
    - constant 常量
    - dto 数据传输对象
        - request 请求对象
        - response 响应对象
    - enum 枚举
    - model 数据模型
        - dao 数据操作层(操作数据库，增删改查业务实现)
        - data 数据结构体(超过3个参数即可用data定义结构体)
        - entity 模型实体(表)
    - service 服务模块(对外提供服务,仅handler或者task能调用,服务模块调用model的方法)
    - utils 工具
- task 定时任务

## 开发规范
- 模式
采用gitflow的协同合作模式

- 响应
json输出

## 测试

- websocket
下载测试工具
```
npm install -g wscat
```

连接
```
wscat -c ws://localhost:8080/v1/ws
```
npm config set prefix "E:\node-v14.2.0-win-x64\node_global"
npm config set cache "E:\node-v14.2.0-win-x64\node_cache"
输入:`hello`


