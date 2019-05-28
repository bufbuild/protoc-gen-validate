package io.envoyproxy.pgv.grpc;

import io.envoyproxy.pgv.ReflectiveValidatorIndex;
import io.envoyproxy.pgv.ValidationException;
import io.envoyproxy.pgv.Validator;
import io.envoyproxy.pgv.ValidatorIndex;
import io.envoyproxy.pgv.grpc.asubpackage.GreeterGrpc;
import io.envoyproxy.pgv.grpc.asubpackage.HelloJKRequest;
import io.envoyproxy.pgv.grpc.asubpackage.HelloResponse;
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
        public void sayHello(HelloJKRequest request, StreamObserver<HelloResponse> responseObserver) {
            responseObserver.onNext(HelloResponse.newBuilder().setMessage("Hello " + request.getName()).build());
            responseObserver.onCompleted();
        }
    };

    @Test
    public void InterceptorPassesValidMessages() {
        serverRule.getServiceRegistry().addService(svc);

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(ValidatorIndex.ALWAYS_VALID);

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        stub.sayHello(HelloJKRequest.newBuilder().setName("World").build());
    }

    @Test
    public void InterceptorPassesValidMessagesGenerated() {
        serverRule.getServiceRegistry().addService(svc);

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ReflectiveValidatorIndex());

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        stub.sayHello(HelloJKRequest.newBuilder().setName("World").build());
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
        assertThatThrownBy(() -> stub.sayHello(HelloJKRequest.newBuilder().setName("Foo").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessage("INVALID_ARGUMENT: one: is invalid - Got ");
    }

    @Test
    public void InterceptorRejectsInvalidMessagesGenerated() {
        // Don't set up server, so it will error if the call goes through

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ReflectiveValidatorIndex());

        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        assertThatThrownBy(() -> stub.sayHello(HelloJKRequest.newBuilder().setName("Foo").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessageStartingWith("INVALID_ARGUMENT: .io.envoyproxy.pgv.grpc.HelloJKRequest.name: must equal World");
    }

    // Also testing compilation of proto files with strings of uppercase characters in their filename.
    @Test
    public void InterceptorRejectsInvalidMessagesGenerated2() {
        // Don't set up server, so it will error if the call goes through

        ValidatingClientInterceptor interceptor = new ValidatingClientInterceptor(new ReflectiveValidatorIndex());

        DismisserGrpc.DismisserBlockingStub stub = DismisserGrpc.newBlockingStub(serverRule.getChannel()).withInterceptors(interceptor);
        assertThatThrownBy(() -> stub.sayGoodbye(GooDBYe.GoodbyeJKRequest.newBuilder().setName("Foo").build()))
                .isInstanceOf(StatusRuntimeException.class)
                .hasMessageStartingWith("INVALID_ARGUMENT: .io.envoyproxy.pgv.grpc.GoodbyeJKRequest.name: must equal World");
    }
}
