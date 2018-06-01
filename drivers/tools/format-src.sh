#!/usr/bin/env bash

# 1 - enable, 0 - disable
debug=0

# Output error messages to stderr.
function error () {
    >&2 echo "$@"
}

# Output debug messages only if debug=1
function debug () {
    if [[ "${debug}" -eq 1 ]]; then
        echo "$@"
    fi
}

# Script require single argument - path to directory of file.
if [ "$#" -ne 1 ]; then
    error "Please pass single parameter: file name or directory"
    exit 1
fi

# Function format format single file passed as first (and single) argument.
# It returns error 3 if file is of unknown type (extension).
# Otherwise it returns result of called formatter (go fmt or clang-format).
function format () {
    ext="${1##*.}"
    case "${ext}" in
        go)
            debug "go fmt \"${1}\""
            go fmt "${1}"
            res=$?
            if [[ ${res} -gt 0 ]]; then
                error "go fmt \"${1}\" returns ${res}"
            fi
            return ${res}
            ;;
        h|c|cpp|m)
            debug "clang-format -i -style=file \"${1}\""
            clang-format -i -style=file "${1}"
            res=$?
            if [[ ${res} -gt 0 ]]; then
                error "clang-format \"${1}\" returns ${res}"
            fi
            return ${res}
            ;;
        *)
            error "File \"${1}\" is of unknown type - unable format."
            return 3
            ;;
    esac
}

# If passed argument is a directory then find all files in this directory (recursively) with known extensions.
# Each found file will be passed to format function.
if [ -d "${1}" ]; then
    err=0
    while read file ; do
        format "${file}"
        newErr=$?
        if [[ ${err} -eq 0 && ${newErr} -gt 0 ]]; then
            err=${newErr}
        fi
    done < <(find "${1}" -type f \( -name "*.go" -or -name "*.h" -or -name "*.c" -or -name "*.cpp" -or -name "*.m" \))
    exit ${err}
fi

# If passed argument is not a directory then check if file with that name exists.
if [ ! -f "${1}" ]; then
    error "File \"${1}\" does not exists."
    exit 2
fi

# Call format function for single file.
format "${1}"
exit $?
