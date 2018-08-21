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

def _protoc_gen_validate_impl(ctx):
  """Generate protos using protoc-gen-validate plugin"""
  protos = []
  for dep in ctx.attr.deps:
    protos += [f for f in dep.proto.direct_sources]

  cc_hdrs = [p.basename[:-len(".proto")] + ".pb.validate.h" for p in protos]
  cc_hdrs += [p.basename[:-len(".proto")] + ".pb.h" for p in protos]

  cc_srcs = [p.basename[:-len(".proto")] + ".pb.validate.cc" for p in protos]
  cc_srcs += [p.basename[:-len(".proto")] + ".pb.cc" for p in protos]

  out_files = [ctx.actions.declare_file(out) for out in cc_hdrs+cc_srcs]

  dir_out = ctx.genfiles_dir.path
  if ctx.label.workspace_root:
    dir_out += ("/"+ctx.label.workspace_root)

  args = [
    "--cpp_out="+dir_out,
    "--plugin=protoc-gen-validate="+ctx.executable._plugin.path,
    "--validate_out=lang=cc:"+ dir_out,
  ]

  tds = depset([], transitive = [dep.proto.transitive_descriptor_sets for dep in ctx.attr.deps])
  descriptor_args = [ds.path for ds in tds]

  if len(descriptor_args) != 0:
    args += ["--descriptor_set_in=%s" % ctx.configuration.host_path_separator.join(descriptor_args)]

  ctx.action(
      inputs=protos + tds.to_list() + [ctx.executable._plugin],
      outputs=out_files,
      arguments=args + [_proto_path(proto) for proto in protos],
      executable=ctx.executable._protoc,
      mnemonic="ProtoGenValidateCcGenerate",
      use_default_shell_env=True,
  )

  return struct(
      files=depset(out_files)
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
    implementation = _protoc_gen_validate_impl,
)
