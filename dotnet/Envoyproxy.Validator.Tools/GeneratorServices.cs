#region Copyright notice and license

// Copyright 2018 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#endregion

using System;
using System.IO;
using System.Text;
using Microsoft.Build.Framework;
using Microsoft.Build.Utilities;

namespace Envoyproxy.Validator
{
    // Abstract class for language-specific analysis behavior, such
    // as guessing the generated files the same way protoc does.
    internal abstract class GeneratorServices
    {
        protected readonly TaskLoggingHelper Log;
        protected GeneratorServices(TaskLoggingHelper log) { Log = log; }

        // Guess whether item's metadata suggests gRPC stub generation.
        // When "gRPCServices" is not defined, assume gRPC is not used.
        // When defined, C# uses "none" to skip gRPC, C++ uses "false", so
        // recognize both. Since the value is tightly coupled to the scripts,
        // we do not try to validate the value; scripts take care of that.
        // It is safe to assume that gRPC is requested for any other value.
        protected bool ValidatorOutputPossible(ITaskItem proto)
        {
            string vm = proto.GetMetadata(Metadata.Validator);
            return string.Equals(vm, "enabled", StringComparison.OrdinalIgnoreCase);
        }

        public abstract string[] GetPossibleOutputs(ITaskItem protoItem);
    }

    // C# generator services.
    internal class CSharpGeneratorServices : GeneratorServices
    {
        public CSharpGeneratorServices(TaskLoggingHelper log) : base(log) { }

        public override string[] GetPossibleOutputs(ITaskItem protoItem)
        {
            bool generate = ValidatorOutputPossible(protoItem);
            if (!generate)
                return new string[0];

            var outputs = new string[1];
            string proto = protoItem.ItemSpec;
            string basename = Path.GetFileNameWithoutExtension(proto);
            string outdir = protoItem.GetMetadata(Metadata.OutputDir).Replace('\\', '/');
            string filename = LowerUnderscoreToUpperCamelProtocWay(basename);
            outputs[0] = Path.Combine(outdir, filename) + "Validator.cs";

            return outputs;
        }

        // This is how the protoc codegen constructs its output filename.
        // See protobuf/compiler/csharp/csharp_helpers.cc:137.
        // Note that protoc explicitly discards non-ASCII letters.
        string LowerUnderscoreToUpperCamelProtocWay(string str)
        {
            var result = new StringBuilder(str.Length, str.Length);
            bool cap = true;
            foreach (char c in str)
            {
                char upperC = char.ToUpperInvariant(c);
                bool isAsciiLetter = 'A' <= upperC && upperC <= 'Z';
                if (isAsciiLetter || ('0' <= c && c <= '9'))
                {
                    result.Append(cap ? upperC : c);
                }
                cap = !isAsciiLetter;
            }
            return result.ToString();
        }
    }
}
