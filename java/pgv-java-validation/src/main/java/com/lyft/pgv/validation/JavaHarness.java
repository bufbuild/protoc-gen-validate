package com.lyft.pgv.validation;

import com.google.common.base.Throwables;
import com.google.protobuf.Message;
import com.lyft.pgv.GeneratedValidatorIndex;
import com.lyft.pgv.ValidationException;
import tests.harness.Harness;
import tests.harness.cases.*;
import tests.harness.cases.other_package.EmbedOuterClass;

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

            Harness.TestCase testCase = Harness.TestCase.parseFrom(System.in);
            Message message = typeMap.unpackAny(testCase.getMessage());
            GeneratedValidatorIndex.validatorFor(message).assertValid(message);

            writeResult(Harness.TestResult.newBuilder().setValid(true).build());
        } catch (ValidationException ex) {
            writeResult(Harness.TestResult.newBuilder().setValid(false).setReason(ex.getMessage()).build());
        } catch (Throwable ex) {
            writeResult(Harness.TestResult.newBuilder().setValid(false).setError(true).setReason(Throwables.getStackTraceAsString(ex)).build());
        }

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
