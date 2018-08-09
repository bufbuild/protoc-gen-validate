$ErrorActionPreference = "Stop";
trap { $host.SetShouldExit(1) }

$package = "github.com/lyft/protoc-gen-validate"

# protoc-gen-go parameters for properly generating the import path for PGV
$validate_import = "Mvalidate/validate.proto=$package/validate"
$go_import = "$validate_import," +`
  "Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any," +`
  "Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration," +`
  "Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct," +`
  "Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp," +`
  "Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers," +`
  "Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor," +`
  "Mgogoproto/gogo.proto=$package/gogoproto"

# protoc-gen-gogo parameters
$gogo_import = "$validate_import," +`
  "Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types," +`
  "Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types," +`
  "Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types," +`
  "Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types," +`
  "Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types," +`
  "Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types," +`
  "Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto"

#lint:
# don't need to run linter on Windows (it already runs on Linux)

#build:
# generates the PGV binary and installs it into $$GOPATH/bin
go install .
if ($LASTEXITCODE -ne 0) {
  exit $LASTEXITCODE
}

#tests:
# runs all tests against the package with race detection and coverage percentage
go test -race -cover
if ($LASTEXITCODE -ne 0) {
  exit $LASTEXITCODE
}

# tests validate proto generation
bazel "--output_base=c:\_pgv" "--bazelrc=windows\bazel.rc" build //validate:go_default_library
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}

$diff = Compare-Object (Get-Content "bazel-out\x64_windows-fastbuild\bin\validate\windows_amd64_pure_stripped\go_default_library%\github.com\lyft\protoc-gen-validate\validate\validate.pb.go") (Get-Content "validate\validate.pb.go")
if ($diff -ne $null) {
  $diff
  exit 1
}

#gogofast:
go build -o gogofast.exe vendor/github.com/gogo/protobuf/protoc-gen-gogofast/main.go
if ($LASTEXITCODE -ne 0) {
  exit $LASTEXITCODE
}
$gogofast_exe = (Get-Item gogofast.exe).FullName

#kitchensink
#generates the kitchensink test protos

Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/kitchensink/go
New-Item -Type Directory -Force tests/kitchensink/go > $nul
Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/kitchensink/gogo
New-Item -Type Directory -Force tests/kitchensink/gogo > $nul

pushd tests/kitchensink
  $protos = ""
  (get-childitem -recurse -Path . -Include *.proto) | foreach {$protos += "$_ "}
  protoc `
    -I "$PWD" `
    -I "$PWD\..\.." `
    "--go_out=$go_import`:./go" `
    "--validate_out=lang=go:./go" `
    "--plugin=protoc-gen-gogofast=$gogofast_exe" `
    "--gogofast_out=$gogo_import`:./gogo" `
    "--validate_out=lang=gogo:./gogo" `
    $protos.Split(" ")

  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

pushd tests/kitchensink/go
  go build .
  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

pushd tests/kitchensink/gogo
  go build .
  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

#testcases: gogofast
# generate the test harness case protos

Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/harness/cases/go
New-Item -Type Directory -Force tests/harness/cases/go > $nul
Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/harness/cases/other_package/go
New-Item -Type Directory -Force tests/harness/cases/other_package/go > $nul
Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/harness/cases/gogo
New-Item -Type Directory -Force tests/harness/cases/gogo > $nul
Remove-Item -Force -Recurse -ErrorAction SilentlyContinue tests/harness/cases/other_package/gogo
New-Item -Type Directory -Force tests/harness/cases/other_package/gogo > $nul

# protoc-gen-go makes us go a package at a time
pushd tests/harness/cases/other_package
  $protos = ""
  (get-item -Path "*.proto") | foreach {$protos += "$_ "}
  protoc `
    -I "$PWD" `
    -I "$PWD\..\..\..\.." `
    "--go_out=$go_import`:./go" `
    "--validate_out=lang=go:./go" `
    "--plugin=protoc-gen-gogofast=$gogofast_exe" `
    "--gogofast_out=$gogo_import`:./gogo" `
    "--validate_out=lang=gogo:./gogo" `
    $protos.Split(" ")

  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

pushd tests/harness/cases
  $protos = ""
  (get-item -Path "*.proto") | foreach {$protos += "$_ "}
  protoc `
    -I "$PWD" `
    -I "$PWD\..\..\.." `
    "--go_out=Mtests/harness/cases/other_package/embed.proto=$package/tests/harness/cases/other_package/go,$go_import`:./go" `
    "--plugin=protoc-gen-gogofast=$gogofast_exe" `
    "--gogofast_out=Mtests/harness/cases/other_package/embed.proto=$package/tests/harness/cases/other_package/gogo,$gogo_import`:./gogo" `
    "--validate_out=lang=go:./go" `
    "--validate_out=lang=gogo:./gogo" `
    $protos.Split(" ")

  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd


#tests/harness/go/harness.pb.go:
# generates the test harness protos
pushd tests/harness
  protoc -I "$PWD" "--go_out=$go_import`:./go" harness.proto
  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

#tests/harness/gogo/harness.pb.go: gogofast
# generates the test harness protos
pushd tests/harness/
  protoc -I "$PWD" "--plugin=protoc-gen-gogofast=$gogofast_exe" "--gogofast_out=$gogo_import`:./gogo" harness.proto
  $le = $LASTEXITCODE
  if ($le -ne 0) {
    popd
    exit $le
  }
popd

#tests/harness/go/main/go-harness
# generates the go-specific test harness
go build -o ./tests/harness/go/main/go-harness.exe ./tests/harness/go/main
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}

#tests/harness/gogo/main/go-harness:
#generates the gogo-specific test harness
go build -o ./tests/harness/gogo/main/go-harness.exe ./tests/harness/gogo/main
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}

#tests/harness/cc/cc-harness
bazel "--output_base=c:\_pgv" "--bazelrc=windows\bazel.rc" build //tests/harness/cc:cc-harness
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}

cp -Force bazel-bin/tests/harness/cc/cc-harness.exe tests/harness/cc/cc-harness.exe

#harness:
# runs the test harness, validating a series of test cases in all supported languages
$go_files = ""
(get-item -path "tests\harness\executor\*.go") | foreach {$go_files += "$_ "}
go run $go_files.Split(" ")
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}

#bazel-harness
# runs the test harness via bazel
bazel "--output_base=c:\_pgv" "--bazelrc=windows\bazel.rc" run //tests/harness/executor:executor
$le = $LASTEXITCODE
if ($le -ne 0) {
  exit $le
}
