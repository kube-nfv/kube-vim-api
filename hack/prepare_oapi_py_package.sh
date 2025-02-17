#!/usr/bin/env bash

PYTHON_PACKAGE_DIR="$1"
rm -rf ${PYTHON_PACKAGE_DIR}/.gitlab-ci.yml
rm -rf ${PYTHON_PACKAGE_DIR}/git_push.sh
rm -rf ${PYTHON_PACKAGE_DIR}/.openapi-generator
