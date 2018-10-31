package com.lyft.pgv.grpc;

import com.lyft.pgv.ValidationException;
import com.lyft.pgv.ValidatorIndex;
import io.grpc.*;

public class ValidatingClientInterceptor implements ClientInterceptor {
    private final ValidatorIndex index;

    public ValidatingClientInterceptor(ValidatorIndex index) {
        this.index = index;
    }

    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
        return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(next.newCall(method, callOptions)) {
            @Override
            public void sendMessage(ReqT message) {
                try {
                    index.validatorFor(message.getClass()).assertValid(message);
                    super.sendMessage(message);
                } catch (ValidationException ex) {
                    Status status = Status.INVALID_ARGUMENT.withDescription(ex.getMessage());
                    throw new StatusRuntimeException(status);
                }
            }
        };
    }
}
