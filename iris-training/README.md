# iris web框架学习

## 目录结构解释

```
.
├── config              # 存放项目的配置信息
├── datamodels          # 存放数据模型
├── models              # 存放业务数据模型
├── datasource          # 可以是mock的数据，提供数据的数据源
├── repositories        # 根据数据模型操作数据组件进行增删改查
├── validators          # 验证器，验证业务数据的正确性
├── services            # 封装model对应的业务逻辑代码
├── utils               # 工具函数
├── rpc                 # rpc代表了rpc服务
├── web                 # web代表了http服务
│   ├── controllers     # 控制请求
│   ├── middleware      # 中间件
│   ├── public          # 静态文件
│   └── views           # 视图模板
├── main.go             # 项目入口文件
├── go.mod              # go module记录第三方依赖包，版本信息
├── go.sum              # go module记录第三方依赖包，版本信息
├── config.toml         # 项目配置文件
├── .env                # 环境变量文件
├── .gitignore          # git忽略文件
├── AUTHORS             # 项目参与者信息
├── CHANGELOG.md        # 项目更新日志
└── README.md           # 项目说明文件 
```
