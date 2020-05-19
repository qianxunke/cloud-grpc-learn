package com.cloud.grpc.java.doubleWayStream.server;

import com.cloud.grpc.doubleWayStream.DoubleWayStreamServiceGrpc;
import com.cloud.grpc.doubleWayStream.RequestMessage;
import com.cloud.grpc.doubleWayStream.ResponseMessage;
import io.grpc.stub.StreamObserver;

public class DoubleWayStreamIml extends DoubleWayStreamServiceGrpc.DoubleWayStreamServiceImplBase {
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
