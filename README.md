### 项目简介
字节青训营大作业， 使用Go语言实现极简版抖音后端系统，详情可看下述文档。<br>
![1683035883947](https://user-images.githubusercontent.com/81971445/235688692-398f7cd2-c2d2-45de-91f0-199abc11d8ed.png)

字节跳动青训营：[https://youthcamp.bytedance.com/](https://youthcamp.bytedance.com/)\
项目文档：[https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#6QCRJV](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#6QCRJV)\
接口文档：[https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c)

### 下载项目
```
git clone https://github.com/Elainasama/Tiktok-Project.git
```

### 实现功能
|  数据库表   | 对应接口功能  |
|  ----  | ----  |
| User  |  用户注册、用户登录、获取用户信息|
| Video  | 上传视频、获取上传列表、推送视频流 |
| Favorite  | 点赞 、 取消点赞 、 点赞列表 |
| Comment  | 评论 、 删除评论 、 评论列表 |
| Relation  | 关注 、 取消关注 、 关注列表 、 粉丝列表 、 互关列表|
| Message | 发送消息 、 聊天记录 |

### 项目结构
- common:数据库以及中间件的初始化
- controller：接受前端信息，调用service层，返回Response响应
- dao：数据库底层增删改查操作
- logger：zap日志配置
- message：前端接收的信息格式
- model：数据库表
- service：处理业务逻辑，调用dao层操作，处理完成后返回给controller层
- config：配置文件
- main：项目入口
- router：项目路由
