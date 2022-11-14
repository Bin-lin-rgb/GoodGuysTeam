/*
 Navicat Premium Data Transfer

 Source Server         : LocalMysql
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : localhost:3306
 Source Schema         : blogserver

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 08/09/2022 20:10:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `post_id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `status` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_author_Id`(`author_id`) USING BTREE,
  INDEX `idx_comment_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 'test', 0, 0, 0, 200, '2022-08-08 20:46:25.522', '2022-08-08 20:46:25.522', NULL);
INSERT INTO `comment` VALUES (2, 'test', 0, 0, 0, 200, '2022-08-08 20:47:06.468', '2022-08-08 20:47:06.468', NULL);

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) UNSIGNED NOT NULL,
  `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_community_id`(`community_id`) USING BTREE,
  UNIQUE INDEX `idx_community_name`(`community_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of community
-- ----------------------------
INSERT INTO `community` VALUES (1, 1, '综合', 'Golang');
INSERT INTO `community` VALUES (2, 2, '后端', '刷题刷题刷题');
INSERT INTO `community` VALUES (3, 3, '前端', 'x1HphYso8M');
INSERT INTO `community` VALUES (4, 4, 'Android', 'csEiB5b6Kv');
INSERT INTO `community` VALUES (5, 5, 'iOS', 'bNPHCppqzE');
INSERT INTO `community` VALUES (6, 6, '人工智能', 'n3nPq4zzmO');
INSERT INTO `community` VALUES (7, 7, '开发工具', 'vfXJQAdgut');
INSERT INTO `community` VALUES (8, 8, '阅读', 'T097aKHwIr');

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `author_id` bigint(20) NULL DEFAULT NULL,
  `community_id` bigint(20) NULL DEFAULT NULL,
  `status` tinyint(4) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_author_id`(`author_id`) USING BTREE,
  INDEX `idx_community_id`(`community_id`) USING BTREE,
  INDEX `idx_post_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES (2, 139209095604338689, 'test', 'content test content', 138919669049327617, 1, 0, '2022-08-18 16:38:06.930', '2022-08-18 16:38:06.930', NULL);
INSERT INTO `post` VALUES (12, 139677924822351873, 'test', 'content test content', 138919669049327617, 4, 0, '2022-08-21 22:15:30.923', '2022-08-21 22:15:30.923', NULL);
INSERT INTO `post` VALUES (13, 139677937875025921, 'test', 'content test content', 138919669049327617, 4, 0, '2022-08-21 22:15:38.710', '2022-08-21 22:15:38.710', NULL);
INSERT INTO `post` VALUES (14, 139677940727152641, 'test', 'content test content', 138919669049327617, 4, 0, '2022-08-21 22:15:40.405', '2022-08-21 22:15:40.405', NULL);
INSERT INTO `post` VALUES (15, 139677942874636289, 'test', 'content test content', 138919669049327617, 4, 0, '2022-08-21 22:15:41.685', '2022-08-21 22:15:41.685', NULL);
INSERT INTO `post` VALUES (16, 140352220511076353, '集成 Markdown 编辑器1 | 青训营笔记', '\n\n## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```vue\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n\n', 138925754648887297, 3, 0, '2022-08-26 13:54:02.475', '2022-08-26 13:54:02.475', NULL);
INSERT INTO `post` VALUES (17, 140355060138770433, '集成 Markdown 编辑器2 | 青训营笔记', '\n\n## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```vue\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n', 138925754648887297, 1, 0, '2022-08-26 14:22:15.025', '2022-08-26 14:22:15.025', NULL);
INSERT INTO `post` VALUES (18, 140355111711932417, '集成 Markdown 编辑器3 | 青训营笔记', '\n\n## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```vue\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n', 138925754648887297, 4, 0, '2022-08-26 14:22:45.768', '2022-08-26 14:22:45.768', NULL);
INSERT INTO `post` VALUES (24, 140366792479473665, '雪花算法生成id，返回给前端不一致 | 青训营笔记', '### 前言\n\n今天发现后端使用雪花算法生成的 PostId 传给前端之后，居然不一致！\n\n导致文章详情页带参数查询为空！\n\n为啥会这样呢？\n\n### 什么是雪花算法？\n\n雪花算法，Twitter的分布式自增ID算法\n\nsnowflake优缺点：\n\n优点：\n\n- （1）经测试snowflake每秒能生成26万个自增可排序的ID。\n- （2）snowflake生成的ID结果是一个64bit大小的整数，转换成字符串后长度最多19。**[注意这里的 19 位数，后面会用]**\n- （3）分布式系统内不会产生ID碰撞（datacenter和workerId作区分）并且效率高。\n\n缺点：依赖机器时钟，如果机器时钟回拨，会导致id重复。由于是部署到分布式环境，每台机器上的时钟不可能完全同步，有时候出现不是全局递增的情况。\n\n### Golang 如何实现？\n\n这里使用的是 sonyflake 库，基本实现和 snowflake 差不多，不过是配位不太一样，影响不大。\n\n地址在这里：[github.com/sony/sonyfl…](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fsony%2Fsonyflake)\n\n具体使用代码：\n\n```go\npackage snowflake\n\nimport (\n	\"fmt\"\n	\"time\"\n\n	\"github.com/sony/sonyflake\"\n)\n\nvar (\n	sonyFlake     *sonyflake.Sonyflake	// 实例\n	sonyMachineID uint16				// 机器ID\n)\n\nfunc getMachineID() (uint16, error) {  // 返回全局定义的机器ID\n	return sonyMachineID, nil\n}\n\n// 需传入当前的机器ID\nfunc Init(machineId uint16) (err error) {\n	sonyMachineID = machineId\n	t, _ := time.Parse(\"2006-01-02\", \"2020-01-01\")	// 初始化一个开始的时间\n	settings := sonyflake.Settings{		// 生成全局配置\n		StartTime: t,\n		MachineID: getMachineID,								// 指定机器ID\n	}\n	sonyFlake = sonyflake.NewSonyflake(settings)	// 用配置生成sonyflake节点\n	return\n}\n\n// GetID 返回生成的id值\nfunc GetID() (id uint64, err error) {		// 拿到sonyflake节点生成id值\n	if sonyFlake == nil {\n		err = fmt.Errorf(\"snoy flake not inited\")\n		return\n	}\n\n	id, err = sonyFlake.NextID()\n	return\n}\n\nfunc main()  { \n	if err := Init(1);err!=nil{\n		fmt.Printf(\"Init failed,err:%v\\n\",err)\n		return\n	}\n	id,_ := GetID()\n	fmt.Println(id)\n}\n\n```\n\n如果将其封装在 snowflake 包里，就这样直接调用即可：\n\n![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c6eff59fbaba4a23b82003f4968c9f1e~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp?)\n\n### 如何解决返回前端的id不一致？\n\n看一下哪里不一样：\n\n139677942874636289 ==> 139677942874636290\n\n好家伙怎么就差一位数！我当时以为是精度的问题，上网一查发现Number精度是16位（雪花ID是19位的），所以是JS的Number数据类型导致的精度丢失。\n\n解决方案：后端的ID(uint64) ==> uint64 转 String ==> 前端使用 String 类型的ID，这样精度就不会丢失了。那前端再把String类型的19位数字传回服务端的时候，再转 uint64 。\n\n### 总结\n\n提前发现 bug 也是一种好事，哈哈。\n\n解决方案：前端用String类型的ID，后端及数据库继续使用 uint64 类型不影响数据库查询执行效率。', 138925754648887297, 3, 0, '2022-08-26 16:18:48.042', '2022-08-26 16:18:48.042', NULL);
INSERT INTO `post` VALUES (25, 141056241383243777, '测试文章发布1', '# **第** **1** **章：原生** **AJAX** \n\n## **1.1 AJAX** **简介** \n\nAJAX 全称为 Asynchronous JavaScript And XML，就是异步的 JS 和 XML。 \n\n通过 AJAX 可以在浏览器中向服务器发送异步请求，最大的优势：无刷新获取数据。 \n\nAJAX 不是新的编程语言，而是一种将现有的标准组合在一起使用的新方式。 \n\n\n\n**使用步骤** \n\n```js\n//\\1) 创建 XMLHttpRequest 对象 \n\nvar xhr = new XMLHttpRequest(); \n\n//\\2) 设置请求信息 \n\nxhr.open(method, url); \n\n//可以设置请求头，一般不设置 \nxhr.setRequestHeader(\'Content-Type\', \'application/x-www-form-urlencoded\'); \n\n//\\3) 发送请求 \n\nxhr.send(body) //get 请求不传 body 参数，只有 post 请求使用 \n\n//\\4) 接收响应 \n\n//xhr.responseXML 接收 xml 格式的响应数据 \n//xhr.responseText 接收文本格式的响应数据 \n\nxhr.onreadystatechange = function (){ \n    if(xhr.readyState == 4 && xhr.status == 200){ \n        var text = xhr.responseText; \n        console.log(text); \n    } \n} \n```\n\n', 138925754648887297, 8, 0, '2022-08-31 10:27:51.618', '2022-08-31 10:27:51.618', '2022-09-01 17:03:16.902');
INSERT INTO `post` VALUES (26, 141202893561659393, '集成 Markdown 编辑器4 | 青训营笔记', '## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```js\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n\n', 138925754648887297, 4, 0, '2022-09-01 10:44:43.125', '2022-09-01 10:44:43.125', '2022-09-01 17:03:15.889');
INSERT INTO `post` VALUES (27, 141203017394290689, '集成 Markdown 编辑器5 | 青训营笔记', '## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```js\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n\n', 138925754648887297, 1, 0, '2022-09-01 10:45:56.934', '2022-09-01 10:45:56.934', '2022-09-01 17:03:14.481');
INSERT INTO `post` VALUES (28, 141241054480302081, '集成 Markdown 编辑器6 | 青训营笔记', '\n\n## 集成 Markdown 编辑器\n\n因为项目使用的技术栈是 vite + vue3，需要使用一个Markdown编辑器来实现发布文章的功能，所以要选择一种。\n\n这里选用的是md-editor-v3，它是Markdown编辑器Vue3版本，使用jsx和 typescript 语法开发，支持切换主题、prettier美化文本等。\n\n\n\n下面介绍一下它的简单使用。\n\n### 安装\n\n```\nyarn add md-editor-v3\n```\n\n\n\n### 使用\n\n新建一个view下的文件\n\n```vue\n<template>\n  <CommHeader />\n  <div class=\"md\"><md-editor v-model=\"text\" /></div>\n</template>\n\n<script setup lang=\"ts\">\nimport CommHeader from \"@/components/common/CommHeader.vue\";\nimport { ref } from \"vue\";\n// md-editor-v3\nimport MdEditor from \"md-editor-v3\";\nimport \"md-editor-v3/lib/style.css\";\nconst text = ref(\"Hello Editor!\");\nconsole.log(text.value);\n</script>\n\n<style scoped>\n.md {\n  height: 100%;\n  width: 100%;\n  position: fixed;\n  top: 70px;\n  left: 0;\n}\n</style>\n\n```\n\n编辑器的内容是使用 vue 的 v-model 来实现数据绑定在 text 中，通过 ref 实现响应式`const text = ref(\"Hello Editor!\");`. 通过 value 属性可以获取具体值。\n\n\n\n在 `<md-editor v-model=\"text\" />` 标签里面，可以选择许多属性。具体查看文档：https://imzbf.github.io/md-editor-v3/docs#%F0%9F%A4%AF%20Props%20%E8%AF%B4%E6%98%8E\n\n\n\n\n\n### 获取html代码\n\n\n\nonHtmlChanged\n\n- **类型**：`(h: string) => void`\n\n  html 变化回调事件，用于获取预览 html 代码。\n\n\n\n编辑器考虑到了可能后端不存储`md`格式的文本，而是`html`内容，所以提供了`onHtmlChanged`方法，用于编辑内容变化后，`marked`编译了内容的回调，入参即是`html`内容。\n\n```js\n<template>\n  <md-editor\n    v-model=\"text\"\n    @onHtmlChanged=\"saveHtml\"\n  />\n</template>\n\n<script lang=\"ts\">\nimport { defineComponent } from \'vue\';\nimport MdEditor from \'md-editor-v3\';\nimport \'md-editor-v3/lib/style.css\';\n\nexport default defineComponent({\n  components: { MdEditor },\n  data() { return { text: \'\' }; },\n  methods: { saveHtml(h: string) { console.log(h) }}\n});\n</script>\n```\n\n\n\n### 总结\n\n文章只是介绍一下这个 Markdown 编辑器，安利一波，具体的实现风格按照 docs 改一下可以了。\n\n更多用法请查看官网：[md-editor-v3](https://zhuanlan.zhihu.com/write)\n\n', 138925754648887297, 3, 0, '2022-09-01 17:03:48.807', '2022-09-01 17:03:48.807', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `company` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司',
  `position` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职位',
  `user_page` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人主页',
  `user_introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人介绍',
  `gender` tinyint(4) NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (47, 'abcd', '$2a$10$Gc8g44IM/V/EDDn6kd5En.25l96dEatECv2.CCIvSmpmq/9Hz.xmi', '', NULL, '', NULL, NULL, 0, '2022-08-15 22:07:09.460', '2022-08-15 22:07:09.460', NULL, 138807352517197825);
INSERT INTO `user` VALUES (48, 'ddff', '$2a$10$F09P5EZpEADzsmf/djWODevj5hwhw/QChuHIhPOKdUVN88XaD5UgO', '', '公司11', '程序员11', '这是个人主页11', '这是个人介绍11', 0, '2022-08-16 16:42:55.323', '2022-08-31 21:38:49.248', NULL, 138919669049327617);
INSERT INTO `user` VALUES (49, '沉默的网友', '$2a$10$imYj9LjGpakv28RdA6Xmx.KbbBwkrTLFpfUblO7K/qD6FmxLIdjG6', '', '公司222', '程序员11', '我，就是太阳', '太阳落山了！！！', 0, '2022-08-16 17:43:22.623', '2022-09-01 16:58:34.346', NULL, 138925754648887297);
INSERT INTO `user` VALUES (50, '456', '$2a$10$wog4Ci81Dn6yr84GNgD92.6PEkQ4vost7ugqxX5QjjvtSDSxkgxw6', '', '', '', '', '', 0, '2022-09-01 14:52:23.923', '2022-09-01 14:52:23.923', NULL, 141227825729372161);

SET FOREIGN_KEY_CHECKS = 1;
