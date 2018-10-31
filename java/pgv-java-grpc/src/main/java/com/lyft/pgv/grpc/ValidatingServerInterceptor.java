package com.lyft.pgv.grpc;

import com.lyft.pgv.ValidationException;
import com.lyft.pgv.ValidatorIndex;
import io.grpc.*;

public class ValidatingServerInterceptor implements ServerInterceptor {
    private final ValidatorIndex index;

    public ValidatingServerInterceptor(ValidatorIndex index) {
        this.index = index;
    }

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        return new ForwardingServerCallListener.SimpleForwardingServerCallListener<ReqT>(next.startCall(call, headers)) {
            @Override
            public void onMessage(ReqT message) {
                try {
                    index.validatorFor(message.getClass()).assertValid(message);
                    super.onMessage(message);
                } catch (ValidationException ex) {
                    Status status = Status.INVALID_ARGUMENT.withDescription(ex.getMessage());
                    call.close(status, new Metadata());
                }
            }
        };
    }
}
