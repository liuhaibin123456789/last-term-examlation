# last-term-examlation
网校第二学期考核

# api

## 所有api基本返回参数

## 下棋时的消息格式

```json
{"qi_zi":11,"from_x":1,"from_y":0,"to_x":1,"to_y":0}//表示qi_zi从（from_x,from_Y）移动到（to_x,to_y）
```

**上面的`qi_zi`的取值如下**

![image-20220612072652307](https://raw.githubusercontent.com/liuhaibin123456789/img-for-cold-bin-blog/master/img/%E8%B1%A1%E6%A3%8B%E5%B8%83%E5%B1%80.png)

其他字段顾名思义

看不到棋盘，只能在终端上看。因为客户端只有postman....

## 一、用户登录

| 返回参数               | 说明                                                  |
| ---------------------- | ----------------------------------------------------- |
| code                 | 状态码                                                |
| msg                   | status为1时，error为空字符串；0时，返回服务端报错信息 |
| data | 会自定义标识某项资源，详情如以下api                   |

### 1. 注册api

- 访问方法

```http
localhost:8085/user/register
无需携带token
```

- 请求参数

| 请求参数  | 类型`Content-Type`          | 说明                                    |
| --------- | --------------------------- | --------------------------------------- |
| user_name | `multipart/form-data`，可选 | 昵称，格式要求英文大小写字母数字8到16位 |
| password  | `multipart/form-data`，必选 | 密码，格式要求英文大小写字母数字8到16位 |
| phone     | `multipart/form-data`，必选 | 手机号作为登录账号，格式要求11位数字    |
| answer    | `multipart/form-data`，可选 | 忘记密码时，需要的问题                  |
| question  | `multipart/form-data`，可选 | 忘记密码时，问题对应的答案              |

- 其他返回参数

| 其他返回参数 | 说明                                                      |
| ------------ | --------------------------------------------------------- |
| token        | 请求成功返回token字符串，12小时过期；请求失败返回空字符串 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjEyMzQ1Njc4OTAxIiwiZXhwIjoxNjQ1MDYxNzY4LCJpYXQiOjE2NDUwMTg1NjgsImlzcyI6ImNvbGQgYmluIFx1MDAyNiB0YW8gcnVpIn0.tz6R3mjK9pwOM_4_WmRX51JrUSOkObBvy_rHmfDA3_k"
  }
  ```

### 2. 密码登录api

- 访问方法

```http
localhost:8085/user/login
不需要token
```

- 请求参数

| 请求参数 | 类型                                | 说明                                    |
| -------- | ----------------------------------- | --------------------------------------- |
| phone    | `multipart/form-data`or `json`,必选 | 注册时的手机号，11位数字即可            |
| password | `multipart/form-data`or `json`,必选 | 密码，格式要求英文大小写字母数字8到16位 |

- 其他返回参数

| 返回参数      | 说明                                                      |
| ------------- | --------------------------------------------------------- |
| refresh_token | 刷新token，30天过期                                       |
| access_token  | 请求成功返回token字符串，12小时过期；请求失败返回空字符串 |

- 返回实例

  ```json
  {
      "code": 1000,
      "msg": "请求成功",
      "data": {
          "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc1Njc3NDcsImlzcyI6ImNvbGQgYmluIn0.L_J90qBhegvfQ6o28bVvepfYKSKbWOkoQjWiBAb3tsM",
          "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjA0NzY0Mjc3MTEyODMyIiwiZXhwIjoxNjU1MDExNzQ3LCJpYXQiOjE2NTQ5NzU3NDcsImlzcyI6ImNvbGQgYmluIn0.awNxWmYrvulH_dMRu-7EnsLEOsLKtGuUxBfzuWAosKA"
      }
  }
  ```

## 二、下棋

### 1. 创建房间

- 访问方法

```http
ws://localhost:8085/room/create
websocket连接
```

- 请求参数

| 请求参数 | 类型                         | 说明 |
| -------- | ---------------------------- | ---- |
| Cookie   | Header里放的access_token必选 |      |

- 其他返回参数

### 2. 搜索房间

- 访问方法

```http
localhost:8085/room/search
```

- 请求参数

| 请求参数 | 类型       | 说明           |
| -------- | ---------- | -------------- |
| phone    | query,必选 | 手机号11位数字 |

- 其他返回参数

### 3. 进入房间

- 访问方法

```http
 ws://localhost:8085/room/enter
 websocket连接
```

- 请求参数

| 请求参数 | 类型                                                         | 说明               |
| -------- | ------------------------------------------------------------ | ------------------ |
| Cookie   | Header里放的access_token必选                                 |                    |
| room_id  | query,必选                                                   | 房间id号           |
| message  | 第一次连接上，需要发送一次”已准备“，不能多发；等待两个客户端都发送已准备时，两个客户端可以进行通信。通信时，发送的消息格式变化如下：`{"qi_zi":11,"from_x":1,"from_y":0,"to_x":1,"to_y":0}` | webocket发送的消息 |

- 其他返回参数
