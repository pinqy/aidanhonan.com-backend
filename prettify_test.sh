#!/bin/bash

MOD_NAME='aidanhonan.com/backend'

RED='\033[0;31m'
GREEN='\033[0;32m'
BOLDRED='\033[1;31m'
BOLDGREEN='\033[1;32m'
BOLDYELLOW='\033[1;33m'
NC='\033[0m' # No Color

TEST_CLASS_BUFFER=""

STDOUT_STARTED=false
STDOUT_BUFFER=""

TESTS_PASSED=0
TESTS_FAILED=0
SUITES_PASSED=0
SUITES_FAILED=0

while IFS= read -r line; do
    # Trim leading whitespace
    line="${line#"${line%%[![:space:]]*}"}"
    
    # "=== RUN" lines mark the start of a single test execution
    if [[ "${line:0:7}" = "=== RUN" ]]; then
        STDOUT_STARTED=true

    # "--- PASS" or "--- FAIL" mark the end of a single test execution
    # Single test passed line
    elif [[ "${line:0:8}" = "--- PASS" ]]; then
        TEST_CLASS_BUFFER+="${GREEN}${line}${NC}\\n"
        STDOUT_STARTED=false
        STDOUT_BUFFER=""
        TESTS_PASSED=$(($TESTS_PASSED+1))
    # Single test failed line
    elif [[ "${line:0:8}" = "--- FAIL" ]]; then
        TEST_CLASS_BUFFER+="${RED}${line}${NC}\\n"
        if [[ "${#STDOUT_BUFFER}" -gt 0 ]]; then
            TEST_CLASS_BUFFER+="${STDOUT_BUFFER}"
        fi
        STDOUT_STARTED=false
        STDOUT_BUFFER=""
        TESTS_FAILED=$(($TESTS_FAILED+1))

    # If test STDOUT has started, store the value but don't print.
    # Will be outputted below the test case only if it fails.
    elif [[ "$STDOUT_STARTED" = true ]]; then
        STDOUT_BUFFER+="\\t${line}\\n"

    # Extraneous lines that can be ignored
    elif [[ "$line" = "FAIL" || "$line" = "PASS" || "${line:0:9}" = "coverage:" ]]; then
        : # Pass

    # End of test class, print all test cases + stdout below the module name
    elif [[ "${line:0:4}" = "FAIL" || "${line:0:2}" == "ok" || "${line:0:${#MOD_NAME}}" == "$MOD_NAME" ]]; then
        # Determine module header color and trim prefix
        MOD_COLOR=$BOLDYELLOW
        if [[ "${line:0:4}" = "FAIL" ]]; then
            line="${line:4:${#line}}"
            MOD_COLOR=$BOLDRED
            SUITES_FAILED=$(($SUITES_FAILED+1))
        elif [[ "${line:0:2}" == "ok" ]]; then
            line="${line:2:${#line}}"
            MOD_COLOR=$BOLDGREEN
            SUITES_PASSED=$(($SUITES_PASSED+1))
        fi

        line="${line#"${line%%[![:space:]]*}"}"

        echo -e "${MOD_COLOR}${line}${NC}"
        if [[ ${#TEST_CLASS_BUFFER} -gt 0 ]]; then
            echo -e "${TEST_CLASS_BUFFER}"
        else
            echo
        fi
        TEST_CLASS_BUFFER=""

    else
        echo -e "${NC}${line}"
    fi
done

echo "Summary:"
if [[ $SUITES_FAILED -gt 0 ]]; then
    echo -e "| Modules: ${GREEN}${SUITES_PASSED} passed, ${RED}${SUITES_FAILED} failed${NC}"
else
    echo -e "| Modules: ${GREEN}${SUITES_PASSED} passed, ${NC}${SUITES_FAILED} failed"
fi
if [[ $TESTS_FAILED -gt 0 ]]; then
    echo -e "| Tests: ${GREEN}${TESTS_PASSED} passed, ${RED}${TESTS_FAILED} failed${NC}"
    echo
    echo -e "${BOLDRED}BUILD FAILED${NC}"
else
    echo -e "| Tests: ${GREEN}${TESTS_PASSED} passed, ${NC}${TESTS_FAILED} failed"
    echo
    echo -e "${BOLDGREEN}BUILD SUCCEEDED${NC}"
fi

echo

# If tests passed run coverage, else fail out
if [[ $TESTS_FAILED = 0 ]]; then
    exit 0
else
    exit 1
fi