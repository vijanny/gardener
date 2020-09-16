#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2020 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0
#
# This file was copied and modified from the kubernetes-sigs/controller-runtime project
# https://github.com/kubernetes-sigs/controller-runtime/blob/a9bd9117a77a2f84bbc546e28991136fe0000dc0/hack/setup-envtest.sh
#
# Modifications SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors

// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o nounset
set -o pipefail

# Turn colors in this script off by setting the NO_COLOR variable in your
# environment to any value:
#
# $ NO_COLOR=1 test.sh
NO_COLOR=${NO_COLOR:-""}
if [ -z "$NO_COLOR" ]; then
  header=$'\e[1;33m'
  reset=$'\e[0m'
else
  header=''
  reset=''
fi

function header_text {
  echo "$header$*$reset"
}

function setup_envtest_env {
  header_text "setting up env vars"

  # Setup env vars
  KUBEBUILDER_ASSETS=${KUBEBUILDER_ASSETS:-""}
  if [[ -z "${KUBEBUILDER_ASSETS}" ]]; then
    export KUBEBUILDER_ASSETS="$1/bin"
  fi
}

# Fetch k8s API gen tools and make it available under KUBEBUILDER_ASSETS.
#
# Skip fetching and untaring the tools by setting the SKIP_FETCH_TOOLS variable
# in your environment to any value:
#
# $ SKIP_FETCH_TOOLS=1 ./check-everything.sh
#
# If you skip fetching tools, this script will use the tools already on your
# machine.
function fetch_envtest_tools {
  SKIP_FETCH_TOOLS=${SKIP_FETCH_TOOLS:-""}
  if [ -n "$SKIP_FETCH_TOOLS" ]; then
    return 0
  fi

  tmp_root=/tmp
  envtest_root_dir=$tmp_root/envtest

  k8s_version="${ENVTEST_K8S_VERSION:-1.17.9}"
  goarch="$(go env GOARCH)"
  goos="$(go env GOOS)"

  if [[ "$goos" != "linux" && "$goos" != "darwin" ]]; then
    echo "OS '$goos' not supported. Aborting." >&2
    return 1
  fi

  local dest_dir="${1}"

  # use the pre-existing version in the temporary folder if it matches our k8s version
  if [[ -x "${dest_dir}/bin/kube-apiserver" ]]; then
    version=$("${dest_dir}"/bin/kube-apiserver --version)
    if [[ $version == *"${k8s_version}"* ]]; then
      header_text "Using cached envtest tools from ${dest_dir}"
      return 0
    fi
  fi

  header_text "fetching envtest tools@${k8s_version} (into '${dest_dir}')"
  envtest_tools_archive_name="kubebuilder-tools-$k8s_version-$goos-$goarch.tar.gz"
  envtest_tools_download_url="https://storage.googleapis.com/kubebuilder-tools/$envtest_tools_archive_name"

  envtest_tools_archive_path="$tmp_root/$envtest_tools_archive_name"
  if [ ! -f $envtest_tools_archive_path ]; then
    curl -sL ${envtest_tools_download_url} -o "$envtest_tools_archive_path"
  fi

  mkdir -p "${dest_dir}"
  tar -C "${dest_dir}" --strip-components=1 -zvxf "$envtest_tools_archive_path"
}

bin_dir="$(realpath "$(dirname $0)/../bin")"
kb_root_dir="$bin_dir/kubebuilder"

mkdir -p            "$kb_root_dir"
fetch_envtest_tools "$kb_root_dir"
setup_envtest_env   "$kb_root_dir"
