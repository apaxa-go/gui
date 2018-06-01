#!/usr/bin/env bash

# Static analyzing Objective-C sources.
# Primary made for AGUI cocoa driver on MacOS.
# Requires "scan-build" from LLVM ("brew install llvm").

sb=/usr/local/opt/llvm/bin/scan-build
sbOpts="-analyze-headers -no-failure-reports -analyzer-config stable-report-filename=true"
sbCheckers=(
	nullability.NullableDereferenced
	nullability.NullableReturnedFromNonnull
	security.insecureAPI.rand
	)
goOpts="-a"

# Output error messages to stderr.
function error () {
    >&2 echo "$@"
}

# Construct result options.
function makeCheckerOpts() {
	for checker in "${sbCheckers[@]}"; do
		echo -n " -enable-checker ${checker}"
	done
}

# Script requires single argument - path to directory of file.
if [ "$#" -ne 1 ]; then
    error "Please pass single parameter: GoLang package name (github.com/some-user/.../some-package)"
    exit 1
fi

sbFullOpts="${sbOpts} `makeCheckerOpts`"
echo Analyzing \"$1\"...
echo "scan-build options: ${sbFullOpts}"
echo "        go options: ${goOpts}"

"${sb}" ${sbFullOpts} "go build ${goOpts} \"$1\""
