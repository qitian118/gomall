# Qitian Shop - 简易版抖音商城

基于微服务架构的电商平台，实现抖音商城核心业务流程，包含完整的前后端技术实践方案。

## 核心功能
### 用户管理
- 注册/登录（JWT认证）
- 信息查询与维护
- 安全退出机制

### 电商功能
- 商品浏览与搜索
- 购物车操作（增/删/改）
- 订单生命周期管理
- 在线支付流程
- 物流通知系统

## 服务架构
采用七项独立微服务：
- auth：令牌生成与验证
- user：用户账户管理
- product：商品数据维护
- cart：购物车状态管理
- order：订单流程处理
- payment：支付渠道对接
- email：异步通知服务

## 技术方案
### 后端技术栈
- 通信框架：Kitex（RPC）、Hertz（HTTP）
- 数据存储：MySQL（主存）、Redis（缓存）
- 服务治理：Consul（发现）、Jaeger（追踪）、Prometheus（监控）

### 前端实现
- 基础技术：HTML5 + CSS3 + JavaScript
- 组件框架：Bootstrap 5.x
- 开发规范：
  - 响应式布局（支持移动端）
  - 模块化CSS命名
  - ES6+语法标准

## 项目规范
- 安全规范：JWT令牌验证 + HTTPS
- 部署方案：Docker容器化部署
