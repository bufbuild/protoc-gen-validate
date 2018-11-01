package com.lyft.pgv.demo;

import com.lyft.pgv.GeneratedValidatorIndex;
import com.lyft.pgv.grpc.ValidatingServerInterceptor;
import greeter.GreeterGrpc;
import greeter.Hello;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

public class HelloServer extends GreeterGrpc.GreeterImplBase {
    public static void main(String... args) throws Exception {
        System.out.println("Listening...");
        ServerBuilder.forPort(9999)
                .intercept(new ValidatingServerInterceptor(GeneratedValidatorIndex.INSTANCE))
                .addService(new HelloServer())
                .build()
                .start()
                .awaitTermination();
    }

    @Override
    public void sayHello(Hello.HelloRequest request, StreamObserver<Hello.HelloResponse> responseObserver) {
        System.out.println("Requested " + request.getTimes() + " responses for " + request.getName());

        for (int i = 0; i < request.getTimes(); i++) {
            responseObserver.onNext(Hello.HelloResponse.newBuilder().setMessage("Hello " + request.getName()).build());
        }
        responseObserver.onCompleted();
    }
}
