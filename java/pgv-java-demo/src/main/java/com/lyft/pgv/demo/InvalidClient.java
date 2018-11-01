package com.lyft.pgv.demo;

import greeter.GreeterGrpc;
import greeter.Hello;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class InvalidClient {
    public static void main(String... args) {
        Hello.HelloRequest request = Hello.HelloRequest.newBuilder()
                .setName("Me")
                .setTimes(2)
                .build();

        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 9999).usePlaintext().build();
        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(channel);

        stub.sayHello(request).forEachRemaining(response -> System.out.println(response.getMessage()));
    }
}
