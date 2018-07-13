#!/usr/bin/env bash

set -e

bootstrap_folder="`pwd`/bootstrap"
deps_folder="${bootstrap_folder}/dependencies"
templates_folder="${bootstrap_folder}/templates"

default_project_folder="skeleton"
project_folder=""
project_path=""

if [ "`which go 2>/dev/null`" = "" ]; then
  echo "Missing Go, please install it first!"
fi

echo "GOPATH=`pwd`" > .env
echo "GOBIN=`pwd`/bin" >> .env

export `cat .env | xargs`

mkdir -p "$GOBIN"

if ! [ -f "$GOBIN/dep" ]; then
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi

read -p "Project src folder [default: ${default_project_folder}]: " project_folder
if [ "${project_folder}" = "" ]; then
  project_folder=${default_project_folder}
fi
project_path="${GOPATH}/src/${project_folder}"

mkdir -p "${project_path}"
if ! [ -f "${project_path}/main.go" ]; then
  cp "${templates_folder}/main.go" "${project_path}/main.go"
  cp "${templates_folder}/main_test.go" "${project_path}/main_test.go"
fi

cd "${project_path}"

if ! [ -f "Gopkg.toml" ]; then
  "$GOBIN/dep" init
fi

for file in `ls ${deps_folder}`; do
  file_path="${deps_folder}/${file}"
  while read dependency; do
    if [ "`echo ${dependency} | grep "#"`" = "" ]; then
      go get ${dependency}
      "${GOBIN}/dep" ensure -add ${dependency}
    fi
  done <${file_path}
done

# "$GOBIN/dep"
