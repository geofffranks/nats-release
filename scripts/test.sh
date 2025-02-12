#!/bin/bash

specificied_package="${1}"

set -e -u

go version # so we see the version tested in CI

SCRIPT_PATH="$(cd "$(dirname "${0}")" && pwd)"

cd "${SCRIPT_PATH}/.."

declare -a packages=(
  "src/code.cloudfoundry.org/nats-v2-migrate"
)

install_ginkgo() {
    echo "Installing Ginkgo V2"
    go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
    ginkgo version
}

test_package() {
  local package=$1
  if [ -z "${package}" ]; then
    return 0
  fi
  shift
  pushd "${package}" &>/dev/null
  ginkgo -r --race -randomize-all -randomize-suites -fail-fast \
      -ldflags="extldflags=-WL,--allow-multiple-definition" \
       "${@}";
  rc=$?
  popd &>/dev/null
  return "${rc}"
}

install_ginkgo

for dir in "${packages[@]}"; do
    test_package "${dir}"
done

