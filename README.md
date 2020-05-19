# Grpc 引进

# 什么是gRPC？
gRPC是一个支持请求/响应和流式处理（非持久化）用例的传输机制。
它是一个模式优先的RPC框架，协议在**Protobuf服务描述符**（_protobuf service descriptor_）中声明，请求和响应将通过 HTTP/2 连接流式的传输。
如下图：
![image.png](https://user-gold-cdn.xitu.io/2020/5/19/1722bd49bdfb5114?w=552&h=327&f=png&s=28284)
从上图和文档中我们可以了解到，用gRPC来进行远程服务调用就仅仅需要gRPC Stub（Client）用Proto Request向远方的gRPC Server发起服务调用，然后远方的gRPC Server通过Proto Response(s)将调用结果返回给gRPC Stub。


上面这段逻辑的背后，gRPC做了什么：
![](https://user-gold-cdn.xitu.io/2020/5/19/1722bd49be69588c?w=482&h=434&f=jpeg&s=21018)
一个gRPC从开始发起请求到返回总共要经历过序列化，编解码，以及网络传输这些内容。这些东西在我们使用gRPC框架做远程服务调用的时候完全感知不到。


**它有几个优点**：

- 模式优先设计倾向于定义良好且分享的服务接口，而不是脆弱的自组织方案；
- 基于Protobuf的wire协议是高效的、众所周知的，并且允许兼容的模式演化；
- 基于 HTTP/2 ，这样它允许在单个连接上复用多个数据流；
- 流式处理的请求和响应是第一类的；
- 许多语言都有可用的工具，不同语言编写的客户端与服务之间可无缝互操作。

**这使得gRPC非常适合**：

- 内部服务之间的连接
- 连接到公开的gRPC API外部服务（甚至是用其它语言编写的服务）
- 给Web或移动设备前端提供数据



# Grpc 四种服务类型
gRPC 允许你定义四类服务方法：
## 单项 RPC
即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
```
rpc SayHello(HelloRequest) returns (HelloResponse){}
```


使用场景：普通的远程调用，像http一样，即请求-响应。实现内容可见以下内容
## 服务端流式 RPC
即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。
```
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}
```


使用场景：一次请求，建立连接后，服务端多次数据返回。比如请求某个股票接口，需要源源不断获取实时的股票信息。
## 客户端流式 RPC
即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。
```
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {}
```


使用场景：一次请求，建立连接后，客户端在此连接上多次向服务端发送消息，待客户端发送完毕，服务端再返回响应。比如以下实例：
做一个加法的服务，服务端接收客户端传过来的一系列的数字（int型），然后进行所有数字的和、数字数量和平均值的统计，将最终统计结果返回给调用者。
## 双向流式 RPC
即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。此模式可以媲美Websocket。
```
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){}
```


使用场景：一次请求，建立连接后，客户端与服务端可向对方发送消息，比如机器人聊天程序，以下是实现一个极其简易的聊天室。

[grpc for java - 双向流式实践](https://github.com/qianxunke/cloud-grpc-learn/blob/master/cloud-grpc-java/REAME.md)