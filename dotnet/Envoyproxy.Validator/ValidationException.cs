using System;

namespace Envoyproxy.Validator
{
    public class ValidationException : Exception
    {
        public ValidationException()
        {
        }

        public ValidationException(string message) : base(message)
        {
        }

        public ValidationException(string message, Exception innerException) : base(message, innerException)
        {
        }

        public static ValidationException New(string field, object value, string message)
        {
            var ex = new ValidationException(message);

            ex.Data["field"] = field;
            ex.Data["value"] = value;

            return ex;
        }
    }
}
