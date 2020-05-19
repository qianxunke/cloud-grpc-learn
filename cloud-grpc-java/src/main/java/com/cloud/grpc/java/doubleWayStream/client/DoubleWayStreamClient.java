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
         String str= scanner.nextLine();
            if(str.equals("EOF")){
                requestObserver.onCompleted();
                break;
            }
            try {
                requestObserver.onNext(RequestMessage.newBuilder().setReqMsg(str).build());
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}
