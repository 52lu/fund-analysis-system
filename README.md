## 1. 项目介绍
使用`Go`开发《基金数据分析系统》，项目系统框架是基于[gin-api-template(https://github.com/52lu/gin-api-template)](https://github.com/52lu/gin-api-template)的基础上做二次开发,该系统计划完成功能，主要有: 基金数据(概括、持仓、历史业绩、财务报表)爬取、基金相似度对比、用户管理等功能。项目源码地址: https://github.com/52lu/fund-analysis-system

## 2. 系统架构图

![](https://gitee.com/QingHui/picGo-img-bed/raw/master/img/20210826184628.png)

## 3. 目录介绍

| 目录\|文件  | 说明                                     |
| ----------- | ---------------------------------------- |
| api         | 控制器                                   |
| config      | 配置相关的结构体                         |
| global      | 全局变量、常量                           |
| initialize  | 初始化相关操作，如:连接Mysql、redis等    |
| internal    | 内部实现（不对外暴露）的代码模块         |
| logs        | 存放日志目录                             |
| middleware  | 中间件                                   |
| model       | 实体类(和表结构对应)、请求入参、返回出参 |
| router      | 所有路由在这个目录下                     |
| service     | 业务实现的相关代码                       |
| test        | 单元测试目录                             |
| utils       | 工具函数目录                             |
| config.yaml | 项目配置文件                             |
| main.go     | 主程序入口                               |
| router.go   | 路由注册管理                             |
| server.go   | 主程序服务代码                           |

## 4. 启动流程

![](https://gitee.com/QingHui/picGo-img-bed/raw/master/img/20210826225841.png)
