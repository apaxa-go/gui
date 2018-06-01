#!/usr/bin/env bash

# Output error messages to stderr.
function error () {
    >&2 echo "$@"
}

# Get absolute path of cocoa driver.
pkgDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Get absolute path of GoLang sources directory
srcDir="${GOPATH}/src/"

# Compute package path as relation between pkgDir and srcDir.
if [[ "${pkgDir}" = "${srcDir}"* ]] ; then
	pkgName=${pkgDir:${#srcDir}}
else
	error "Current package \"${pkgDir}\" is outside of GoLang's source dir \"${srcDir}\"."
	exit 1
fi

../tools/analyze-objc.sh "${pkgName}"