package com.lyft.pgv.grpc;

import com.lyft.pgv.ReflectiveValidatorIndex;
import com.lyft.pgv.ValidationException;
import com.lyft.pgv.Validator;
import com.lyft.pgv.ValidatorIndex;
import io.grpc.BindableService;
import io.grpc.ServerInterceptors;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import io.grpc.testing.GrpcServerRule;
import org.junit.Rule;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class ValidatingServerInterceptorTest {
    @Rule
    public GrpcServerRule serverRule = new GrpcServerRule();

    private BindableService svc = new GreeterGrpc.GreeterImplBase() {
        @Override
        public void sayHello(Hello.HelloRequest request, StreamObserver<Hello.HelloResponse> responseObserver) {
            responseObserver.onNext(Hello.HelloResponse.newBuilder().setMessage("Hello " + request.getName()).build());
            responseObserver.onCompleted();
        }
    };

    @Test
    public void InterceptorPassesValidMessages() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(ValidatorIndex.ALWAYS_VALID);

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        stub.sayHello(Hello.HelloRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorPassesValidMessagesGenerated() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(new ReflectiveValidatorIndex());

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        stub.sayHello(Hello.HelloRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorRejectsInvalidMessages() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(new ValidatorIndex() {
            @Override
            public <T> Validator<T> validatorFor(Class clazz) {
                return proto -> {
                    throw new ValidationException("one", "", "is invalid");
                };
            }
        });

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        assertThatThrownBy(() -> stub.sayHello(Hello.HelloRequest.newBuilder().setName("World").build()))
            .isInstanceOf(StatusRuntimeException.class)
            .hasMessage("INVALID_ARGUMENT: one: is invalid - Got ");
    }

    @Test
    public void InterceptorRejectsInvalidMessagesGenerated() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(new ReflectiveValidatorIndex());

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        assertThatThrownBy(() -> stub.sayHello(Hello.HelloRequest.newBuilder().setName("Bananas").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessageStartingWith("INVALID_ARGUMENT: .com.lyft.pgv.grpc.HelloRequest.name: must equal World");
    }
}
