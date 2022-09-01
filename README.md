



# 00 项目分工

| **团队成员** | **主要贡献**                                                 |
| ------------ | ------------------------------------------------------------ |
| 林斌         | 首页、文章详情页、创作者中心布局和文章管理部分、编辑个人信息、后端 |
| 杜万凯       | 创作者中心的首页部分                                         |
| 戴柏霖       | 个人信息页面                                                 |

由于杜万凯和戴柏霖他们比较忙，所以他们不太有时间，我（林斌）也是边学边做，是到处碰壁然后几乎都是自己写的，导致项目完成度不高，我也是只能大致完成一些基本功能和后端的代码。

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
    │      article.ts
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
                EditAccount.vue
                EditIndex.vue
                EditProfile.vue
                ListItem.vue
                UserArticleList.vue
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
| 首页 GET/postlist/:time 根据时间倒序展示的文章列表           | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/1首页.png) |
| 注册和登录 POST/user/login       Login POST/user/register        Register | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/2登录.png) |
| 创作者中心 [文章管理]  GET/user/bloglist  通过userID查询文章详情（这个userID不需要前端直接传，前端只要携带token访问即可，后端会根据token解析userID用于查询该用户的文章列表） | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/3创作者中心.png) |
| 写文章和发布文章 POST/post 编写文章之后可以选择相应内容的标签进行发布 | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/4发布文章.png) |
| 删除发布的文章DELETE user/delete/:id 根据 文章的 id 进行删除，需要优先通过token鉴权 | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/5删除文章.png) |
| 个人主页展示                                                 | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/6个人主页.png) |
| 修改个人资料                                                 | ![](https://cdn.jsdelivr.net/gh/Bin-lin-rgb/blog-img@main/7修改个人资料.png) |

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

- 用户之间的交互还未完成（点赞、评论等）



### 反思

​      在项目过程当中，发现分工、api设计很重要，同时发现自学能力和经验的积累十分重要，而且平时一定要养成写笔记的习惯。



### 一些个人总结

​       首先十分感谢这次字节跳动举办的青训营，课程内容很丰富，引导我们学习的方向。这次活动也督促我不荒废暑假时光，在这之前我连一个页面也写不出来，通过这次青训营的学习，起码掌握了一些技能。

​       在项目过程中，由于队友在暑假的时间不是很充裕，项目的大部分功能都是我写的，但我基本上是处于一种边学边做的状态，在这过程中四处碰壁，很难受，不过解决一个问题之后也获得很舒心的心情。

​       在这个项目的过程当中，我也更加深刻意识到实践的重要性，之前学的东西只要不去实践，10天可能就忘光了，那约等于没学。所以我觉得边学边实践才是符合前端的学习方式吧。
