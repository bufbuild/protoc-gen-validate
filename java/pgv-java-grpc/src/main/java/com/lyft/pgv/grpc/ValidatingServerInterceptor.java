package com.lyft.pgv.grpc;

import com.lyft.pgv.ValidationException;
import com.lyft.pgv.ValidatorIndex;
import io.grpc.*;

import java.util.concurrent.atomic.AtomicBoolean;

public class ValidatingServerInterceptor implements ServerInterceptor {
    private final ValidatorIndex index;

    public ValidatingServerInterceptor(ValidatorIndex index) {
        this.index = index;
    }

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        return new ForwardingServerCallListener.SimpleForwardingServerCallListener<ReqT>(next.startCall(call, headers)) {

            // Implementations are free to block for extended periods of time. Implementations are not
            // required to be thread-safe.
            private boolean aborted = false;

            @Override
            public void onMessage(ReqT message) {
                try {
                    index.validatorFor(message.getClass()).assertValid(message);
                    super.onMessage(message);
                } catch (ValidationException ex) {
                    Status status = Status.INVALID_ARGUMENT.withDescription(ex.getMessage());
                    aborted = true;
                    call.close(status, new Metadata());
                }
            }

            @Override
            public void onHalfClose() {
                if (!aborted) {
                    super.onHalfClose();
                }
            }
        };
    }
}
