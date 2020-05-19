# grpc for java 双向流式通信实践



# 一. 摘要
   此篇文章将用实例介绍grpc四种服务类型中的双向流式通信，模拟一个简单的聊天室功能。
# 二. 实践
整体项目如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589860646675-cc175bb5-ae8a-40f1-b71a-650b6d68e381.png#align=left&display=inline&height=305&margin=%5Bobject%20Object%5D&name=image.png&originHeight=610&originWidth=824&size=130357&status=done&style=none&width=412)
其中cloud-grpc-java为maven项目，cloud-grpc-protos为定义接口项目。
## 1.通过protobuf定义接口和数据类型
在cloud-grpc-protos文件夹下创建doubleWayStream.proto,内容如下：
```
syntax = "proto3";

option go_package = "pbfs/double_way_stream";
option java_multiple_files = true;
option java_package = "com.cloud.grpc.doubleWayStream";
option java_outer_classname = "DoubleWayStreamProto";
option objc_class_prefix = "DWS";

package double_way_stream;

//双向流式
service DoubleWayStreamService{
  rpc DoubleWayStreamFun(stream RequestMessage) returns (stream ResponseMessage){}
}

message RequestMessage{
  string req_msg = 1;
}

message ResponseMessage{
  string rsp_msg = 1;
}
```
  以上，一个 _简单 的双向流式RPC_ 


## 2.maven 配置
创建一个如上图（cloud-grpc-java)的maven项目,pom.xml加入grpc开发相关配置，如下：
```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.cloud.grpc</groupId>
    <artifactId>java</artifactId>
    <version>1.0.0.0</version>
    <name>java</name>
    <description>Demo project for grpc for java</description>

    <properties>
        <java.version>11</java.version>
        <grpc.version>1.29.0</grpc.version>
        <protobuf.version>3.11.0</protobuf.version>
    </properties>

    <dependencies>

        <dependency>
            <groupId>io.grpc</groupId>
            <artifactId>grpc-netty-shaded</artifactId>
            <version>${grpc.version}</version>
        </dependency>

        <dependency>
            <groupId>io.grpc</groupId>
            <artifactId>grpc-protobuf</artifactId>
            <version>${grpc.version}</version>
        </dependency>

        <dependency>
            <groupId>io.grpc</groupId>
            <artifactId>grpc-stub</artifactId>
            <version>${grpc.version}</version>
        </dependency>

        <dependency> <!-- necessary for Java 9+ -->
            <groupId>org.apache.tomcat</groupId>
            <artifactId>annotations-api</artifactId>
            <version>6.0.53</version>
            <scope>provided</scope>
        </dependency>

    </dependencies>

    <build>

        <extensions>
            <extension>
                <groupId>kr.motd.maven</groupId>
                <artifactId>os-maven-plugin</artifactId>
                <version>1.5.0.Final</version>
            </extension>
        </extensions>

        <plugins>
            <plugin>
                <groupId>org.xolstice.maven.plugins</groupId>
                <artifactId>protobuf-maven-plugin</artifactId>
                <version>0.5.1</version>

                <configuration>
                    <protocArtifact>com.google.protobuf:protoc:${protobuf.version}:exe:${os.detected.classifier}</protocArtifact>
                   <pluginId>grpc-java</pluginId>
                    <pluginArtifact>io.grpc:protoc-gen-grpc-java:${grpc.version}:exe:${os.detected.classifier}</pluginArtifact>
                    <protoSourceRoot>../cloud-grpc-protos</protoSourceRoot>
                </configuration>

                <executions>
                    <execution>
                        <goals>
                            <goal>compile</goal>
                            <goal>compile-custom</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>

</project>

```
说明： <protoSourceRoot>../cloud-grpc-protos</protoSourceRoot> 红色标注部分需要根据pom.xml与上面创建proto所在路径相匹配，这样protobuf 插件才会根据此目录找到定义的proto文件生成相关代码。
## 3.生成grpc,protobuf相关类
```bash
mvn protobuf:compile
mvn protobuf:compile-custom
```
  同步pom.xml,即可在target下生成如下源码：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589868549530-f00252ce-2eec-454e-bbe3-06ac65956bba.png#align=left&display=inline&height=518&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1036&originWidth=858&size=316175&status=done&style=none&width=429)
## 4.编写服务端，并运行
服务端源码如图：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589868748683-0c92a295-dc5b-4db1-b8e1-a7a783d2cf7f.png#align=left&display=inline&height=208&margin=%5Bobject%20Object%5D&name=image.png&originHeight=416&originWidth=902&size=115756&status=done&style=none&width=451)
首先实现proto文件中定义的rpc DoubleWayStream 的服务端流式响应接口，DoubleWayStreamIml.java如下
```java
package com.cloud.grpc.java.doubleWayStream.server;

import com.cloud.grpc.doubleWayStream.DoubleWayStreamServiceGrpc;
import com.cloud.grpc.doubleWayStream.RequestMessage;
import com.cloud.grpc.doubleWayStream.ResponseMessage;
import io.grpc.stub.StreamObserver;

public class DoubleWayStreamIml extends DoubleWayStreamServiceGrpc.DoubleWayStreamServiceImplBase {
   
    //声明此服务端流响应，为了后面通过控制台向后端发送消息
    private StreamObserver<com.cloud.grpc.doubleWayStream.ResponseMessage> responseOb;
    
    @Override
    public StreamObserver<com.cloud.grpc.doubleWayStream.RequestMessage> doubleWayStreamFun(StreamObserver<com.cloud.grpc.doubleWayStream.ResponseMessage> responseObserver) {
        this.responseOb=responseObserver;
       return new StreamObserver<RequestMessage>() {
            @Override
            public void onNext(RequestMessage requestMessage) {
                System.out.println("[收到客户端消息]: " + requestMessage.getReqMsg());
                responseObserver.onNext(ResponseMessage.newBuilder().setRspMsg("hello client ,I'm Java grpc Server,your message '" + requestMessage.getReqMsg() + "'").build());
            }

            @Override
            public void onError(Throwable throwable) {
                throwable.fillInStackTrace();
            }

            @Override
            public void onCompleted() {
                responseObserver.onCompleted();
            }
        };

    }
    public StreamObserver<ResponseMessage> getResponseOb() {
        return responseOb;
    }
}

```
如上图， private StreamObserver<com.cloud.grpc.doubleWayStream.ResponseMessage> responseOb 声明此服务端流响应，为了后面通过控制台向后端发送消息。
编码服务端启动类DoubleWayStreamServer.java，如下
```java
package com.cloud.grpc.java.doubleWayStream.server;

import com.cloud.grpc.doubleWayStream.ResponseMessage;
import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;
import java.util.Scanner;

public class DoubleWayStreamServer {

    public static void main(String[] args) {
        ServerBuilder<?> serverBuilder = ServerBuilder.forPort(8899);
        DoubleWayStreamIml doubleWayStreamIml=new DoubleWayStreamIml();
        serverBuilder.addService(doubleWayStreamIml);
        Server server = serverBuilder.build();
        try {
            server.start();
            //开启线程向客户端输入
            new Thread(new Runnable() {
                @Override
                public void run() {
                    Scanner scanner=new Scanner(System.in);
                    for (;true;){
                        String str=scanner.nextLine();
                        if(str.equals("EOF")){
                            break;
                        }
                        try {
                            doubleWayStreamIml.getResponseOb().onNext(ResponseMessage.newBuilder().setRspMsg(str).build());
                        }catch (Exception e){
                            System.out.println("【异常】：没有客户端连接...");
                            //一般客户端链接失败就会断开
                            e.printStackTrace();
                        }
                  }
                }
            }).start();
            server.awaitTermination();

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }

    }
}

```
如上图所示，开启一个新的线程，为了向客户端推送消息。
通过main函数启动，此时我们向控制台输入一条消息，结果如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589869604060-787af773-a95e-49b8-b97d-37dac742f672.png#align=left&display=inline&height=146&margin=%5Bobject%20Object%5D&name=image.png&originHeight=292&originWidth=1494&size=121399&status=done&style=none&width=747)
可以看到，由于我们没有客户端与其建立连接，所以会在控制台打印：没有客户连接


## 5.编写，运行客户端
编写DoubleWayStreamClient.java,如下：
```java
package com.cloud.grpc.java.doubleWayStream.client;

import com.cloud.grpc.doubleWayStream.DoubleWayStreamServiceGrpc;
import com.cloud.grpc.doubleWayStream.RequestMessage;
import com.cloud.grpc.doubleWayStream.ResponseMessage;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.util.Scanner;

public class DoubleWayStreamClient {
    public static void main(String[] args) {
        //使用usePlaintext，否则使用加密连接
        ManagedChannelBuilder<?> channelBuilder = ManagedChannelBuilder.forAddress("localhost", 8899).usePlaintext();

        ManagedChannel channel = channelBuilder.build();

        StreamObserver<RequestMessage> requestObserver = DoubleWayStreamServiceGrpc.newStub(channel).doubleWayStreamFun(new StreamObserver<ResponseMessage>() {
            @Override
            public void onNext(ResponseMessage value) {
                System.out.println("[收到服务端发来] : " + value.getRspMsg());
            }
            
            @Override
            public void onError(Throwable t) {

            }
            @Override
            public void onCompleted() {
            }
        });

        Scanner scanner = new Scanner(System.in);
        for (; true; ) {
         String str=   scanner.nextLine();
            if(str.equals("EOF")){
                requestObserver.onCompleted();
                break;
            }
            try {
                requestObserver.onNext(RequestMessage.newBuilder().setReqMsg(scanner.nextLine()).build());
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}

```
在上面源码中，我们实现了proto文件中rpc方法接口中的客户端流，在next方法中打印服务端的信息，执行main函数，在控制台输入"你好，我是客户端，我上线了",
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589870103268-eb8f4d51-5ef6-47b7-9c4b-d3307e51709b.png#align=left&display=inline&height=122&margin=%5Bobject%20Object%5D&name=image.png&originHeight=244&originWidth=2280&size=184910&status=done&style=none&width=1140)
可见，服务端作出了响应，现在我们切换到服务端，让服务端想客户端发送消息，如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589870365871-38c48077-cb5c-4a6b-adc0-c5cc70c5f6fd.png#align=left&display=inline&height=131&margin=%5Bobject%20Object%5D&name=image.png&originHeight=262&originWidth=1368&size=116178&status=done&style=none&width=684)
再看看客户端：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589870408451-b485c9ea-fb9b-4765-89c5-1be5d301cda0.png#align=left&display=inline&height=209&margin=%5Bobject%20Object%5D&name=image.png&originHeight=418&originWidth=2554&size=280925&status=done&style=none&width=1277)
可见服务端可以向客户端发送消息。这样便实现了一个极其简单的聊天室。
