# Telepush

Vercel Telegram 消息推送机器人

## 一键部署

<a href="https://vercel.com/new/project?template=https://github.com/indes/telepush"><img src="https://vercel.com/button"></a>

### 环境变量设置
在项目设置中添加`telegram_bot_token`环境变量，如图下图所示（将xxxx换成你的telegram bot token）
![image](https://user-images.githubusercontent.com/4705478/122669874-d113ee80-d1f1-11eb-8add-77587c64053e.png)


## 使用

### GET 请求
```
curl -X GET "http://<your_url>/api/notify?u=123&m=HelloWorld"
```

### POST 请求
```
curl -X POST "http://<your_url>/api/notify?u=123" \
  -H "Content-Type: application/json" \
  -d '{"title": "标题", "body": "内容"}'
```

或在 JSON body 中指定用户 ID：
```
curl -X POST "http://<your_url>/api/notify" \
  -H "Content-Type: application/json" \
  -d '{"title": "标题", "body": "内容", "u": 123}'
```

### 路由参数说明

| url 路径   | 参数（类型, 说明）        | 说明         |
| ---------- | ---------------------- | ------------ |
| api/notify | m(消息内容)，u(telegram用户id) | 发送文字消息 |
| api/notify | JSON body: title, body, u | POST请求发送消息 |
