package io.envoyproxy.pgv.grpc;

import io.envoyproxy.pgv.ReflectiveValidatorIndex;
import io.envoyproxy.pgv.ValidationException;
import io.envoyproxy.pgv.Validator;
import io.envoyproxy.pgv.ValidatorIndex;
import io.envoyproxy.pgv.grpc.asubpackage.GreeterGrpc;
import io.envoyproxy.pgv.grpc.asubpackage.HelloJKRequest;
import io.envoyproxy.pgv.grpc.asubpackage.HelloResponse;
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
        public void sayHello(HelloJKRequest request, StreamObserver<HelloResponse> responseObserver) {
            responseObserver.onNext(HelloResponse.newBuilder().setMessage("Hello " + request.getName()).build());
            responseObserver.onCompleted();
        }
    };

    @Test
    public void InterceptorPassesValidMessages() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(ValidatorIndex.ALWAYS_VALID);

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        stub.sayHello(HelloJKRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorPassesValidMessagesGenerated() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(new ReflectiveValidatorIndex());

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        stub.sayHello(HelloJKRequest.newBuilder().setName("World").build());
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
        assertThatThrownBy(() -> stub.sayHello(HelloJKRequest.newBuilder().setName("World").build()))
            .isInstanceOf(StatusRuntimeException.class)
            .hasMessage("INVALID_ARGUMENT: one: is invalid - Got ");
    }

    @Test
    public void InterceptorRejectsInvalidMessagesGenerated() {
        ValidatingServerInterceptor interceptor = new ValidatingServerInterceptor(new ReflectiveValidatorIndex());

        serverRule.getServiceRegistry().addService(ServerInterceptors.intercept(svc, interceptor));

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel());
        assertThatThrownBy(() -> stub.sayHello(HelloJKRequest.newBuilder().setName("Bananas").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessageStartingWith("INVALID_ARGUMENT: .io.envoyproxy.pgv.grpc.HelloJKRequest.name: must equal World");
    }
}
