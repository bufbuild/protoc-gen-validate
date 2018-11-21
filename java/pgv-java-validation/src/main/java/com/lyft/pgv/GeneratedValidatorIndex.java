package com.lyft.pgv;

import java.lang.reflect.InvocationTargetException;
import java.util.Map;

public final class GeneratedValidatorIndex implements ValidatorIndex {
  private GeneratedValidatorIndex() {}

  public static final GeneratedValidatorIndex INSTANCE = new GeneratedValidatorIndex();

  private final Map<java.lang.Class, Validator> VALIDATOR_INDEX = new java.util.HashMap<>();

  private static final Validator<Object> ALWAYS_VALID =
      new Validator<Object>() {
        @Override
        public void assertValid(Object proto) throws ValidationException {
          System.err.println("always valid");
          // Do nothing. Always valid.
        }
      };

  private static void ensureValidatorLoaded(java.lang.Class clazz) {
    Class enclosingClass = clazz;
    while (enclosingClass.getEnclosingClass() != null) {
      enclosingClass = enclosingClass.getEnclosingClass();
    }
    try {
      String validatorClassName = enclosingClass.getName() + "Validator";
      Class validatorClass = clazz.getClassLoader().loadClass(validatorClassName);
      validatorClass.getDeclaredMethod("registerAllValidators").invoke(null);
    } catch (ClassNotFoundException
        | NoSuchMethodException
        | IllegalAccessException
        | InvocationTargetException e) {
      System.err.println("Couldn't load validator class for " + clazz.toString());
    }
  }

  @SuppressWarnings("unchecked")
  public <T> Validator<T> validatorFor(java.lang.Class clazz) {
    // Try to load the message validation class.
    ensureValidatorLoaded(clazz);
    /*
    System.err.println("Retrieving validator for " + clazz.toString());
    for (Class c : VALIDATOR_INDEX.keySet()) {
      System.err.println("  k: " + c.toString());
    }
    */
    return VALIDATOR_INDEX.getOrDefault(clazz, ALWAYS_VALID);
  }

  @SuppressWarnings("unchecked")
  public static <T> Validator<T> validatorFor(Object instance) {
    return INSTANCE.validatorFor(instance == null ? null : instance.getClass());
  }

  public static <T> void registerValidator(
      java.lang.Class<? extends T> clazz, Validator<T> validator) {
    INSTANCE.VALIDATOR_INDEX.put(clazz, validator);
  }
}
