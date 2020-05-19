package com.cloud.grpc.android;

import androidx.appcompat.app.AppCompatActivity;

import android.app.Activity;
import android.os.AsyncTask;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.cloud.grpc.hello.HelloGrpc;
import com.cloud.grpc.hello.HelloRequest;
import com.cloud.grpc.hello.HelloResponse;

import java.lang.ref.WeakReference;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class MainActivity extends AppCompatActivity {

   static TextView messageTxView;
   static Button btnToGo;
   static Button btnToJava;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        messageTxView=findViewById(R.id.tx_message);
        btnToGo=findViewById(R.id.btn_go);
        btnToJava=findViewById(R.id.btn_java);
        btnToGo.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
              send(8488);
            }
        });
        btnToJava.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
               send(2000);
            }
        });
    }


    private void send(int port){
        btnToGo.setEnabled(false);
        btnToGo.setEnabled(false);
        new GrpcTask(this).execute(port);
    }


    public static class GrpcTask extends AsyncTask<Integer,Void,String>{

        private final WeakReference<Activity> activityWeakReference;
        private ManagedChannel channel;

        public GrpcTask(Activity activity){
            this.activityWeakReference=new WeakReference<>(activity);
        }


        @Override
        protected String doInBackground(Integer... integers) {
            try {
                channel = ManagedChannelBuilder.forAddress("192.168.2.102", integers[0]).usePlaintext().build();
                HelloGrpc.HelloBlockingStub stub = HelloGrpc.newBlockingStub(channel);
                HelloRequest request = HelloRequest.newBuilder().setName("Android").build();
                HelloResponse reply = stub.sayHello(request);
                return reply.getMessage();
            }catch (Exception e){
                e.printStackTrace();
                return e.getMessage();
            }

        }

        @Override
        protected void onPostExecute(String s) {
            super.onPostExecute(s);
            try {
                channel.shutdown().awaitTermination(1, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
            Activity activity = activityWeakReference.get();
            if (activity == null) {
                return;
            }
            messageTxView .setText(s);
            btnToGo.setEnabled(true);
            btnToGo.setEnabled(true);
        }
    }



}
