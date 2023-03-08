# Gazelle Repro

Compiling results in the following error:

```
% bazel run :gazellerepro          
INFO: Invocation ID: 9eb27dbc-f716-425e-9507-71b32939c8b9
ERROR: BUILD.bazel:16:11: no such package '@org_golang_x_exp//slices': BUILD file not found in directory 'slices' of external repository @org_golang_x_exp. Add a BUILD file to a directory to mark it as a package. and referenced by '//:gazellerepro_lib'
ERROR: Analysis of target '//:gazellerepro' failed; build aborted: 
INFO: Elapsed time: 0.073s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (0 packages loaded, 0 targets configured)
ERROR: Build failed. Not running target
```

Moving the call to `gazelle_dependencies` below
`load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")`  resolves the issue.
