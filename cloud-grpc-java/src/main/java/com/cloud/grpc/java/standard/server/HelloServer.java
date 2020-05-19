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
