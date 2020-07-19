using System;
using System.Linq;
using Google.Protobuf;
using Google.Protobuf.Reflection;
using Google.Protobuf.WellKnownTypes;
using Envoyproxy.Validator;

namespace Tests.Harness
{
    class Program
    {
        static object Unpack(Any message)
        {
            // TODO: Is there a smarter way to load the correct message definition?
            var descriptor = AppDomain.CurrentDomain.GetAssemblies()
                .Where(x => x.GetName().Name == "Harness")
                .SelectMany(x => x.GetTypes())
                .Select(x => x.GetProperty("Descriptor")?.GetValue(null) as MessageDescriptor)
                .First(x => x != null && message.Is(x));

            return descriptor.Parser.ParseFrom(message.Value);
        }

        static void Main(string[] args)
        {
            var stdin = Console.OpenStandardInput();
            var stdout = Console.OpenStandardOutput();

            var testResult = new TestResult();

            try
            {
                var testCase = TestCase.Parser.ParseFrom(stdin);

                var message = Unpack(testCase.Message);
                if (message is IValidateable validateable)
                {
                    try
                    {
                        validateable.Validate();

                        testResult.Valid = true;
                    }
                    catch (ValidationException ex)
                    {
                        testResult.Reason = ex.Message;
                    }
                }
                else
                {
                    testResult.Error = true;
                    testResult.Reason = "Message is not IValidateable";
                }
            }
            catch (Exception ex)
            {
                testResult.Error = true;
                testResult.Reason = ex.Message + Environment.NewLine + ex.StackTrace.ToString();
            }

            testResult.WriteTo(stdout);
        }
    }
}
