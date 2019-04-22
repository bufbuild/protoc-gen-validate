package io.envoyproxy.pgv.grpc;

import io.envoyproxy.pgv.ReflectiveValidatorIndex;
import io.envoyproxy.pgv.ValidationException;
import io.envoyproxy.pgv.Validator;
import io.envoyproxy.pgv.ValidatorIndex;
import io.grpc.BindableService;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import io.grpc.testing.GrpcServerRule;
import org.junit.Rule;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class ValidatingClientInterceptorTest {
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
        serverRule.getServiceRegistry().addService(svc);

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(ValidatorIndex.ALWAYS_VALID);

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        stub.sayHello(Hello.HelloRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorPassesValidMessagesGenerated() {
        serverRule.getServiceRegistry().addService(svc);

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ReflectiveValidatorIndex());

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        stub.sayHello(Hello.HelloRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorRejectsInvalidMessages() {
        // Don't set up server, so it will error if the call goes through

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ValidatorIndex() {
            @Override
            public <T> Validator<T> validatorFor(Class clazz) {
                return proto -> {
                    throw new ValidationException("one", "", "is invalid");
                };
            }
        });

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        assertThatThrownBy(() -> stub.sayHello(Hello.HelloRequest.newBuilder().setName("Foo").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessage("INVALID_ARGUMENT: one: is invalid - Got ");
    }

    @Test
    public void InterceptorRejectsInvalidMessagesGenerated() {
        // Don't set up server, so it will error if the call goes through

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ReflectiveValidatorIndex());

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        assertThatThrownBy(() -> stub.sayHello(Hello.HelloRequest.newBuilder().setName("Foo").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessageStartingWith("INVALID_ARGUMENT: .io.envoyproxy.pgv.grpc.HelloRequest.name: must equal World");
    }
}
