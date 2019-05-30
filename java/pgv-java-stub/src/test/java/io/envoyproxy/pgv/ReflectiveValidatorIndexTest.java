package io.envoyproxy.pgv;

import io.envoyproxy.pvg.cases.TokenUse;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class ReflectiveValidatorIndexTest {
    @Test
    public void indexFindsOuterMessage() throws ValidationException {
        TokenUse token = TokenUse.newBuilder().setPayload(TokenUse.Payload.newBuilder().setToken(TokenUse.Payload.Token.newBuilder().setValue("FOO"))).build();
        ReflectiveValidatorIndex index = new ReflectiveValidatorIndex();
        Validator<TokenUse> validator = index.validatorFor(TokenUse.class);

        assertThat(validator).withFailMessage("Unexpected Validator.ALWAYS_VALID").isNotEqualTo(Validator.ALWAYS_VALID);
        validator.assertValid(token);
    }

    @Test
    public void indexFindsEmbeddedMessage() throws ValidationException {
        TokenUse.Payload payload = TokenUse.Payload.newBuilder().setToken(TokenUse.Payload.Token.newBuilder().setValue("FOO")).build();
        ReflectiveValidatorIndex index = new ReflectiveValidatorIndex();
        Validator<TokenUse.Payload> validator = index.validatorFor(TokenUse.Payload.class);

        assertThat(validator).withFailMessage("Unexpected Validator.ALWAYS_VALID").isNotEqualTo(Validator.ALWAYS_VALID);
        validator.assertValid(payload);
    }

    @Test
    public void indexFindsDoubleEmbeddedMessage() throws ValidationException {
        TokenUse.Payload.Token token = TokenUse.Payload.Token.newBuilder().setValue("FOO").build();
        ReflectiveValidatorIndex index = new ReflectiveValidatorIndex();
        Validator<TokenUse.Payload.Token> validator = index.validatorFor(TokenUse.Payload.Token.class);

        assertThat(validator).withFailMessage("Unexpected Validator.ALWAYS_VALID").isNotEqualTo(Validator.ALWAYS_VALID);
        validator.assertValid(token);
    }
}
