package com.lyft.pgv.validation;

import com.google.common.base.Throwables;
import com.google.protobuf.ExtensionRegistry;
import com.google.protobuf.Message;
import com.lyft.pgv.ReflectiveValidatorIndex;
import com.lyft.pgv.UnimplementedException;
import com.lyft.pgv.ValidationException;
import tests.harness.Harness;
import tests.harness.cases.*;
import tests.harness.cases.other_package.EmbedOuterClass;
import com.lyft.pgv.validate.Validate;

import java.io.IOException;
import java.util.Arrays;

@SuppressWarnings("unchecked")
public class JavaHarness {
    public static void main(String[] args) {
        try {
            ProtoTypeMap typeMap = ProtoTypeMap.of(Arrays.asList(
                    Bool.getDescriptor().toProto(),
                    Bytes.getDescriptor().toProto(),
                    Enums.getDescriptor().toProto(),
                    Maps.getDescriptor().toProto(),
                    Messages.getDescriptor().toProto(),
                    Numbers.getDescriptor().toProto(),
                    Oneofs.getDescriptor().toProto(),
                    Repeated.getDescriptor().toProto(),
                    Strings.getDescriptor().toProto(),
                    WktAny.getDescriptor().toProto(),
                    WktDuration.getDescriptor().toProto(),
                    WktTimestamp.getDescriptor().toProto(),
                    WktWrappers.getDescriptor().toProto(),
                    EmbedOuterClass.getDescriptor().toProto()
            ));

            ExtensionRegistry registry = ExtensionRegistry.newInstance();
            Validate.registerAllExtensions(registry);

            Harness.TestCase testCase = Harness.TestCase.parseFrom(System.in, registry);
            Message message = typeMap.unpackAny(testCase.getMessage());
            ReflectiveValidatorIndex.validatorFor(message).assertValid(message);

            writeResult(Harness.TestResult.newBuilder().setValid(true).build());
        } catch (UnimplementedException ex) {
            writeResult(Harness.TestResult.newBuilder().setValid(false).setAllowFailure(true).setReason(ex.getMessage()).build());
        } catch (ValidationException ex) {
            writeResult(Harness.TestResult.newBuilder().setValid(false).setReason(ex.getMessage()).build());
        } catch (Throwable ex) {
            writeResult(Harness.TestResult.newBuilder().setValid(false).setError(true).setReason(Throwables.getStackTraceAsString(ex)).build());
        }

        System.out.flush();
        System.exit(0);
    }

    private static void writeResult(Harness.TestResult result) {
        try {
            result.writeTo(System.out);
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}
