### 相关信息
应用前后端分离实现
后端：golang
前端：vue
框架：go-zero、Mint UI
功能清单：
- 用户注册、登录
- 用户搜索、添加好友、删除好友
- 发送/接收文本消息（支持离线消息的接收）

### 几张功能时序图
1、添加好友
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/lTlVekwXJAkGZp5T7MGwH4EEWtJxWPfw6Eycldi5.png)
> 1、用户A请求api接口添加用户B为好友
> 2、api接口处理成功后向websocket服务发送添加消息
> 3、websocket服务收到消息后转发个用户B（在线 ? 发送 : 入离线消息队列）
> 4、用户B收到好友添加消息，点击同意或拒绝请求api。
> 5、api处理成功后向websocket服务发送消息
> 6、websocket服务转发消息到用户A 完成好友添加（在线 ? 发送 : 入离线消息队列）

2、消息发送
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/pl37G5DW9tKqnwa6SlAecBjOlHp8O9sTmOoNXCcd.png)

3、离线消息处理
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/dtUExEnOx6wYUTf68taFj4Nmh0nRTy9OiWSuYdGo.png)

> Tip：
> 1、所有的消息转发都会验证（在线 ? 发送 : 入离线消息队列）。待用户下次登录上线时再发送。
> 2、所有消息都是本地缓存，数据库有保存，客户端没有调用接口获取。

### 部分效果图
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/voaxiYfa11BUA4Bhbsc477iyvgjz6Mf4Fg7EOvZk.png)
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/sCsuL3zhGUM7iHtFcpzOH96UsDE1WeDTWt0oLPEg.png)
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/yKgYSpEIJIymy3my8kUHTHbS23naT2dys8vPkPD9.png)

![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/ymXMZEcgrfGfAM5H9LrLSJ7nvra2vxgclDZcxKvi.png)
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/okFn4M18whXcwimmLUE7DDPqWocrhMot6rfsBuGQ.png)
![image.png](http://blogapi.wuyan94zl.cn/storage/blogs/eKWYxd3hy0eLGmCwCXuiHG0j3fSXlzLuooeFTE0T.png)

最后附上体验地址：http://chat.wuyan94zl.cn
源码待优化后在开源。