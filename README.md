

# 01 项目实现

### 1.1 技术选型与相关开发文档

前端：前端三件套 + Vue3 +  Vite +  Arco Design Vue + Element-plus + TypeScript 。

后端：Golang + Gin + Gorm + mysql

### 1.2 项目代码介绍

#### 1.2.1 前端项目结构

一开始使用 element plus 组件库，后来发现 Arco Design Vue 很不错，然后就两个都使用了。

打包工具选用 vite ，发现启动速度真的很快。

```Shell
└─src
    │  App.vue
    │  main.ts
    │  
    ├─assets
    │          
    ├─components
    │  │  testData.vue
    │  │  
    │  └─common
    │          ArticleItem.vue   
    │          CommHeader.vue    // 公共导航栏
    │          SimpleHeader.vue    // 创作者中心导航
    │          SubHeader.vue    // 首页子导航栏
    │          
    ├─router    // 路由
    │      index.ts
    │      
    ├─stores    // pinia 状态管理
    │      index.ts
    │      user.ts
    │      
    ├─utils
    │  │  aes.ts
    │  │  
    │  └─api    // 请求服务端数据
    │          article.ts
    │          login.ts
    │          request.ts    // axios 二次封装（拦截器）
    │          
    └─views
        ├─ArticleDetails    // 文章详情
        │      ArticleIndex.vue
        │      
        ├─Creator          // 创作者中心
        │      CreatorIndex.vue
        │      ListItem.vue
        │      RightIndex.vue
        │      ShowArticle.vue
        │      SuccessPage.vue
        │      WriteArticle.vue
        │      
        ├─Home         // 首页
        │      HomeIndex.vue
        │      
        └─UserInfo     // 用户相关
                edit.vue
                UserIndex.vue
                UserLogin.vue
```

#### 1.2.2 后端项目结构

我当时想的是侧重点在前端，所以 控制层 、逻辑层 写在一起了，也没考虑使用到太多框架技术。

```Shell
├─Config
├─Controller
├─Dao
├─Logger
├─middleware
├─Model
├─pkg
│  ├─jwt
│  └─snowflake
├─Router
├─server
│  └─img
├─Utils
├─config.json
└─main.go
```

# 02 测试结果

| 功能（接口）                                                 | 图片                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| 首页 GET/postlist/:time 根据时间倒序展示的文章列表           | ![img](https://ksr7oe3m3x.feishu.cn/space/api/box/stream/download/asynccode/?code=NTJiMWYyMzc1OTkyYjE3ZDQ1MmFmYTNkOTVlODE1NWZfMXNlRDV4UDVtTU83N0xHNHhJRFByOGkydDdRTndXSWVfVG9rZW46Ym94Y25peGtzUkNiOFM5QTRocm9BS0pCZ2toXzE2NjE1MDgwMTE6MTY2MTUxMTYxMV9WNA) |
| 注册和登录 POST/user/login       Login POST/user/register        Register | ![img](https://ksr7oe3m3x.feishu.cn/space/api/box/stream/download/asynccode/?code=MjczY2YzYTFiNmFiMGMwMzIxYjBmN2Y2ZjdlMTFkY2FfaWozSkNLNll5WEtYTWlFdk1DMXdmOFhiOG5NMEcwUWRfVG9rZW46Ym94Y255ejV5SW9UTVg0eXFBdU5yOURRbENlXzE2NjE1MDgwMTE6MTY2MTUxMTYxMV9WNA) |
| 创作者中心 [文章管理]  GET/user/bloglist  通过userID查询文章详情（这个userID不需要前端直接传，前端只要携带token访问即可，后端会根据token解析userID用于查询该用户的文章列表） | ![img](https://ksr7oe3m3x.feishu.cn/space/api/box/stream/download/asynccode/?code=MjQzNDNkYTk2ZmExNzU0YTZiMzk3NTNkMjJjZWUwYjFfeW9pMTdRTUZsMmczOVRvczhBTlFVZ1Fqd0doYWs3N2FfVG9rZW46Ym94Y25UbmFiWFJyMk5kcU51Qnhya0MzQmhiXzE2NjE1MDgwMTE6MTY2MTUxMTYxMV9WNA) |
| 写文章和发布文章 POST/post 编写文章之后可以选择相应内容的标签进行发布 | ![img](https://ksr7oe3m3x.feishu.cn/space/api/box/stream/download/asynccode/?code=ZGI5MGZiYmI0NzY1OTkzNzAzZWUzYjcyNmFkNmU1ZjJfZWQ0TExGdjhoNm1rTlQ0N0lrWG50WGk4bkpvT2I1MWJfVG9rZW46Ym94Y25kMURRc0xhMm5ZeG9TTTdSV3ExbXRmXzE2NjE1MDgwMTE6MTY2MTUxMTYxMV9WNA) |
| 删除发布的文章DELETE user/delete/:id 根据 文章的 id 进行删除，需要优先通过token鉴权 | ![img](https://ksr7oe3m3x.feishu.cn/space/api/box/stream/download/asynccode/?code=NzMxYTU3ZDBkNWMzYzkwZDBlMjk5ZTEyZjBjZTEzMWJfU0pUdWppNkdtZ0o3RlhKbzZ0bExaTjc5bENRZk1ObUpfVG9rZW46Ym94Y25mZ3lUZEQ1R0NoZ054VkJ6bTE1ZjZkXzE2NjE1MDgwMTE6MTY2MTUxMTYxMV9WNA) |

# 03 Demo 演示视频

视频地址：[仿掘金项目.mp4](https://ksr7oe3m3x.feishu.cn/file/boxcnDhlfp6ohLSFPfpMWqzwDvg) 

# 04 项目总结与反思

> 1. 目前仍存在的问题
>
> 1. 已识别出的优化项
>
> 1. 架构演进的可能性
>
> 1. 项目过程中的反思与总结

### 仍然存在的问题 & 优化项

- 没有做移动端适配

- 组件拆分有些不太合理

- 有些数据可以直接组件之间进行通信，不必再次请求服务端

- 文章详情的目录暂时没有做分级处理

### 反思

​      在项目过程当中，发现分工、api设计很重要，同时发现自学能力和经验的积累十分重要，平时一定要养成写笔记的习惯
