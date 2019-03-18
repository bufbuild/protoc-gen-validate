def _proto_path(proto):
    """
    The proto path is not really a file path
    It's the path to the proto that was seen when the descriptor file was generated.
    """
    path = proto.path
    root = proto.root.path
    ws = proto.owner.workspace_root
    if path.startswith(root):
        path = path[len(root):]
    if path.startswith("/"):
        path = path[1:]
    if path.startswith(ws):
        path = path[len(ws):]
    if path.startswith("/"):
        path = path[1:]
    return path

def _protoc_cc_output_files(proto_file_sources):
    cc_hdrs = []
    cc_srcs = []

    for p in proto_file_sources:
        basename = p.basename[:-len(".proto")]

        cc_hdrs.append(basename + ".pb.h")
        cc_hdrs.append(basename + ".pb.validate.h")

        cc_srcs.append(basename + ".pb.cc")
        cc_srcs.append(basename + ".pb.validate.cc")

    return cc_hdrs + cc_srcs

def _proto_sources(ctx):
    protos = []
    for dep in ctx.attr.deps:
        protos += [f for f in dep.proto.direct_sources]

    return protos

def _output_dir(ctx):
    dir_out = ctx.genfiles_dir.path
    if ctx.label.workspace_root:
        dir_out += "/" + ctx.label.workspace_root
    return dir_out

def _protoc_gen_validate_cc_impl(ctx):
    """Generate C++ protos using protoc-gen-validate plugin"""
    protos = _proto_sources(ctx)

    cc_files = _protoc_cc_output_files(protos)
    out_files = [ctx.actions.declare_file(out) for out in cc_files]

    dir_out = _output_dir(ctx)

    args = [
        "--cpp_out=" + dir_out,
        "--validate_out=lang=cc:" + dir_out,
    ]

    return _protoc_gen_validate_impl(
        ctx = ctx,
        lang = "cc",
        protos = protos,
        out_files = out_files,
        protoc_args = args,
        package_command = "true",
    )

def _protoc_gen_validate_java_impl(ctx):
    """Generate Java protos using protoc-gen-validate plugin"""
    protos = _proto_sources(ctx)

    out_file = ctx.actions.declare_file(ctx.label.name + ".validate.srcjar")

    dir_out = _output_dir(ctx)

    args = [
        "--java_out=" + dir_out,
        "--validate_out=lang=java:" + dir_out,
    ]

    jar_path = out_file.path[len(dir_out) + 1:]

    return _protoc_gen_validate_impl(
        ctx = ctx,
        lang = "java",
        protos = protos,
        out_files = [out_file],
        protoc_args = args,
        package_command = "(cd " + dir_out + " && find . -name \*.java | xargs jar cf " + jar_path + ")",
    )

def _protoc_python_output_files(proto_file_sources):
    python_srcs = []

    for p in proto_file_sources:
        basename = p.basename[:-len(".proto")]

        python_srcs.append(basename + "_pb.py")
        python_srcs.append(basename + "_pb_validate.py")

def _protoc_gen_validate_python_impl(ctx):
    """Generate Python protos using protoc-gen-validate plugin"""
    protos = _proto_sources(ctx)

    python_files = _protoc_python_output_files(protos)
    out_files = [ctx.actions.declare_file(out) for out in python_files]

    dir_out = _output_dir(ctx)

    args = [
        "--python_out=" + dir_out,
        "--validate_out=lang=python:" + dir_out,
    ]

    return _protoc_gen_validate_impl(
        ctx = ctx,
        lang = "python",
        protos = protos,
        out_files = [out_files],
        protoc_args = args,
        package_command = "true",
    )

def _protoc_gen_validate_impl(ctx, lang, protos, out_files, protoc_args, package_command):
    protoc_args.append("--plugin=protoc-gen-validate=" + ctx.executable._plugin.path)

    dir_out = ctx.genfiles_dir.path
    if ctx.label.workspace_root:
        dir_out += "/" + ctx.label.workspace_root

    tds = depset([], transitive = [dep.proto.transitive_descriptor_sets for dep in ctx.attr.deps])
    descriptor_args = [ds.path for ds in tds]

    if len(descriptor_args) != 0:
        protoc_args += ["--descriptor_set_in=%s" % ctx.configuration.host_path_separator.join(descriptor_args)]

    package_command = package_command.format(dir_out = dir_out)

    ctx.actions.run_shell(
        outputs = out_files,
        inputs = protos + tds.to_list(),
        tools = [ctx.executable._plugin, ctx.executable._protoc],
        command = " && ".join([
            ctx.executable._protoc.path + " $@",
            package_command,
        ]),
        arguments = protoc_args + [_proto_path(proto) for proto in protos],
        mnemonic = "ProtoGenValidate" + lang.capitalize() + "Generate",
        use_default_shell_env = True,
    )

    return struct(
        files = depset(out_files),
    )

cc_proto_gen_validate = rule(
    attrs = {
        "deps": attr.label_list(
            mandatory = True,
            providers = ["proto"],
        ),
        "_protoc": attr.label(
            cfg = "host",
            default = Label("@com_google_protobuf//:protoc"),
            executable = True,
            single_file = True,
        ),
        "_plugin": attr.label(
            cfg = "host",
            default = Label("@com_lyft_protoc_gen_validate//:protoc-gen-validate"),
            allow_files = True,
            executable = True,
        ),
    },
    output_to_genfiles = True,
    implementation = _protoc_gen_validate_cc_impl,
)

java_proto_gen_validate = rule(
    attrs = {
        "deps": attr.label_list(
            mandatory = True,
            providers = ["proto"],
        ),
        "_protoc": attr.label(
            cfg = "host",
            default = Label("@com_google_protobuf//:protoc"),
            executable = True,
            single_file = True,
        ),
        "_plugin": attr.label(
            cfg = "host",
            default = Label("@com_lyft_protoc_gen_validate//:protoc-gen-validate"),
            allow_files = True,
            executable = True,
        ),
    },
    output_to_genfiles = True,
    implementation = _protoc_gen_validate_java_impl,
)

python_proto_gen_validate = rule(
    attrs = {
        "deps": attr.label_list(
            mandatory = True,
            providers = ["proto"],
        ),
        "_protoc": attr.label(
            cfg = "host",
            default = Label("@com_google_protobuf//:protoc"),
            executable = True,
            single_file = True,
        ),
        "_plugin": attr.label(
            cfg = "host",
            default = Label("@com_lyft_protoc_gen_validate//:protoc-gen-validate"),
            allow_files = True,
            executable = True,
        ),
    },
    output_to_genfiles = True,
    implementation = _protoc_gen_validate_python_impl,
)
