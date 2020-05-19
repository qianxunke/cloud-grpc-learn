# grpc for java-普通RPC 开发实例

# 一. 摘要
    此篇文章将用实例介绍grpc四种服务类型中的最普通的单项 rpc。
# 二. 实践
整体项目如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589860646675-cc175bb5-ae8a-40f1-b71a-650b6d68e381.png#align=left&display=inline&height=305&margin=%5Bobject%20Object%5D&name=image.png&originHeight=610&originWidth=824&size=130357&status=done&style=none&width=412)
其中cloud-grpc-java为maven项目，cloud-grpc-protos为定义接口项目。
## 1.通过protobuf定义接口和数据类型
在cloud-grpc-protos文件夹下创建hello.proto,内容如下：
```
syntax = "proto3";

option go_package = "pbfs/hello";
option java_multiple_files = true;
option java_package = "com.cloud.grpc.hello";
option java_outer_classname = "HelloProto";
option objc_class_prefix = "HL";

package hello;

service Hello {
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloResponse {
  string message = 1;
}
```
   以上，一个 _简单 RPC_ ， 客户端使用存根发送请求到服务器并等待响应返回，就像平常的函数调用一样。定义一个SayHello rpc服务，入参：HelloRequest，返参：HelloResponse


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
说明： <protoSourceRoot>../cloud-grpc-protos</protoSourceRoot> 红色标注部分需要根据pom.xml与上面创建proto所在路径相匹配，这样protobuf 插件才会根据此目录找到定义的proto文件生成相关代码。
## 3.生成grpc,protobuf相关类
```bash
mvn protobuf:compile
mvn protobuf:compile-custom
```
  同步pom.xml,即可在target下生成如下源码：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589866036730-1ca3bb00-bc8a-42ee-92fb-e05ac86684ff.png#align=left&display=inline&height=417&margin=%5Bobject%20Object%5D&name=image.png&originHeight=834&originWidth=860&size=326988&status=done&style=none&width=430)
## 4.编写服务端，并运行
编写HelloServer.java如下
```java
package com.cloud.grpc.java.standard.server;

import com.cloud.grpc.hello.HelloGrpc;
import com.cloud.grpc.hello.HelloRequest;
import com.cloud.grpc.hello.HelloResponse;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.IOException;
import java.util.concurrent.TimeUnit;

public class HelloServer {
    
    private Server server;
    private void start() throws IOException {
        /* The port on which the server should run */
        int port = 50051;
        server = ServerBuilder.forPort(port)
                .addService(new HelloIml())  //这里可以添加多个模块
                .build()
                .start();
        System.out.println("Server started, listening on " + port);
        Runtime.getRuntime().addShutdownHook(new Thread() {
            @Override
            public void run() {
                // Use stderr here since the logger may have been reset by its JVM shutdown hook.
                System.err.println("*** shutting down gRPC server since JVM is shutting down");
                try {
                    HelloServer.this.stop();
                } catch (InterruptedException e) {
                    e.printStackTrace(System.err);
                }
                System.err.println("*** server shut down");
            }
        });
    }

    private void stop() throws InterruptedException {
        if (server != null) {
            server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
        }
    }
    
    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }
    
    public static void main(String[] args) throws IOException, InterruptedException {
        final HelloServer server = new HelloServer();
        server.start();
        server.blockUntilShutdown();
    }

    private static class HelloIml extends HelloGrpc.HelloImplBase{
        @Override
        public void sayHello(HelloRequest request, StreamObserver<HelloResponse> responseObserver) {
           // super.sayHello(request, responseObserver);
            HelloResponse helloResponse=HelloResponse.newBuilder().setMessage("Hello "+request.getName()+", I'm Java grpc Server").build();
            responseObserver.onNext(helloResponse);
            responseObserver.onCompleted();
        }
    }
}

```
通过main函数启动：结果如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589867709891-aca57a15-9f86-4c85-a561-89046d203ac0.png#align=left&display=inline&height=103&margin=%5Bobject%20Object%5D&name=image.png&originHeight=206&originWidth=1372&size=77538&status=done&style=none&width=686)


## 5.编写，运行客户端
编写HelloClient.java,如下：
```java
package com.cloud.grpc.java.standard.client;

import com.cloud.grpc.hello.HelloGrpc;
import com.cloud.grpc.hello.HelloRequest;
import com.cloud.grpc.hello.HelloResponse;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.util.concurrent.TimeUnit;

public class HelloClient {

    //远程连接管理器,管理连接的生命周期
    private final ManagedChannel channel;
    private final HelloGrpc.HelloBlockingStub blockingStub;

    public HelloClient(String host, int port) {
        //初始化连接
        channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
        //初始化远程服务Stub
        blockingStub = HelloGrpc.newBlockingStub(channel);
    }


    public void shutdown() throws InterruptedException {
        //关闭连接
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    public String sayHello(String name) {
        //构造服务调用参数对象
        HelloRequest request = HelloRequest.newBuilder().setName(name).build();
        //调用远程服务方法
        HelloResponse response = blockingStub.sayHello(request);
        //返回值
        return response.getMessage();
    }


    public static void main(String[] args) throws InterruptedException {
        HelloClient client = new HelloClient("127.0.0.1", 50051);
        //服务调用
        String content = client.sayHello("Java client");
        //打印调用结果
        System.out.println(content);
        //关闭连接
        client.shutdown();
    }

}
```
执行main函数，结果如下：
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1112927/1589867869149-20e94c61-8154-445d-8d41-c113f91886bc.png#align=left&display=inline&height=58&margin=%5Bobject%20Object%5D&name=image.png&originHeight=116&originWidth=1316&size=76680&status=done&style=none&width=658)
可见已从服务端获取响应。
