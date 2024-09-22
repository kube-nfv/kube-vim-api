set -o errexit
set -o nounset
set -o pipefail
set -x

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=$GOPATH/src/k8s.io/code-generator


source "${CODEGEN_PKG}/kube_codegen.sh"


kube::codegen::gen_client \
    --output-dir "${SCRIPT_ROOT}/pkg/client" \
    --output-pkg github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client \
    --boilerplate /dev/null \
    --with-watch \
    --with-applyconfig \
    "${SCRIPT_ROOT}/pkg/apis"
    

