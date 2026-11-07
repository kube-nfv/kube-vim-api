Code cloned from the https://github.com/kubeovn/kube-ovn/tree/master to allow standalone kube-ovn API types.

Importing the types from the kube-ovn module would pull its full dependency graph, its `replace` directives (which Go ignores in dependencies) and its Go and k8s versions. This copy builds against the k8s version in this module's go.mod. To update it, copy `pkg/apis/kubeovn`, `pkg/internal` and the two `pkg/ovsdb/ovnnb` files from the kube-ovn release tag, rewrite the `github.com/kubeovn/kube-ovn/pkg/` imports to this module, then run `make generate-client`.

/pkg/client directory contains generated client using k8s code-generator and /pkg/apis/kubeovn

All code in this foled under Apache License 2.0
