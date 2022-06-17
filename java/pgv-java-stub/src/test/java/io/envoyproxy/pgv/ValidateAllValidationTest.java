package io.envoyproxy.pgv;

import com.google.protobuf.Any;
import com.google.protobuf.Descriptors;
import com.google.protobuf.TypeRegistry;
import com.google.protobuf.util.JsonFormat;
import io.envoyproxy.pgv.validate.Validate;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.Reader;
import java.util.List;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.*;
import tests.harness.cases.KitchenSink;
import tests.harness.cases.KitchenSink.KitchenSinkMessage;

public class ValidateAllValidationTest {

    @Test
    public void validateAllInterceptorWorks() throws ValidationException {
        ValidateAllExceptionInterceptor interceptor = new ValidateAllExceptionInterceptor();
        ReflectiveValidatorIndex index = new ReflectiveValidatorIndex();
        ValidatorContext context = new ValidatorContext(index);
        List<Descriptors.Descriptor> descriptors = KitchenSink.getDescriptor().getMessageTypes();
        TypeRegistry registry = TypeRegistry.newBuilder().add(descriptors).build();
        
        
        KitchenSinkMessage kitchenSink = loadFileFromJson("/bad-kitchen-sink.json",registry);
        
        assertFalse(context.isValid(kitchenSink,interceptor));
        List<ValidationException> resultList = interceptor.getAllValidationExceptions();
        resultList.forEach(exc -> {
            System.out.println(exc.getMessage());
        });
        assertEquals("Interceptor should have collected some exceptions", 22,interceptor.getAllValidationExceptions().size());
    }

    private KitchenSinkMessage loadFileFromJson(String filename,TypeRegistry typeRegistry) {
        final JsonFormat.Parser jsonParser = JsonFormat.parser().usingTypeRegistry(typeRegistry);
        final KitchenSinkMessage.Builder sinkBuilder = KitchenSinkMessage.newBuilder();
        final InputStream stream = getClass().getResourceAsStream(filename);
        assertNotNull("Stream failed to load", stream);
        final Reader resourceReader = new InputStreamReader(stream);

        try {
            jsonParser.merge(resourceReader, sinkBuilder);
        } catch (IOException ex) {
            fail(ex.getMessage());
        }
        return sinkBuilder.build();
    }

}
