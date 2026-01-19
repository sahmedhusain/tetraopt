#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
PASSED=0
FAILED=0

# Function to count empty spaces (dots) in output
count_dots() {
    echo "$1" | tr -cd '.' | wc -c | tr -d ' '
}

# Function to run test
run_test() {
    local test_name="$1"
    local test_file="$2"
    local expected_result="$3"
    local expected_dots="$4"

    echo "=========================================="
    echo "Testing: $test_name"
    echo "File: $test_file"
    echo "------------------------------------------"

    if [ ! -f "$test_file" ]; then
        echo -e "${RED}✗ FAIL: Test file not found${NC}"
        ((FAILED++))
        return
    fi

    # Run the program and capture output
    output=$(go run . "$test_file" 2>&1)
    exit_code=$?

    # Check if expecting ERROR
    if [ "$expected_result" == "ERROR" ]; then
        if echo "$output" | grep -q "ERROR"; then
            echo -e "${GREEN}✓ PASS: Program correctly outputs ERROR${NC}"
            ((PASSED++))
        else
            echo -e "${RED}✗ FAIL: Expected ERROR but got:${NC}"
            echo "$output"
            ((FAILED++))
        fi
    else
        # Check for successful execution
        if [ $exit_code -eq 0 ]; then
            dots=$(count_dots "$output")
            echo "Output:"
            echo "$output"
            echo "Empty spaces (dots): $dots"
            
            if [ "$expected_dots" != "" ]; then
                if [ "$dots" -eq "$expected_dots" ]; then
                    echo -e "${GREEN}✓ PASS: Correct number of empty spaces ($expected_dots)${NC}"
                    ((PASSED++))
                else
                    echo -e "${RED}✗ FAIL: Expected $expected_dots empty spaces but got $dots${NC}"
                    ((FAILED++))
                fi
            else
                echo -e "${GREEN}✓ PASS: Program executed successfully${NC}"
                ((PASSED++))
            fi
        else
            echo -e "${RED}✗ FAIL: Program failed with exit code $exit_code${NC}"
            echo "$output"
            ((FAILED++))
        fi
    fi
    echo ""
}

# Main test execution
echo "=========================================="
echo "TETRIS OPTIMIZER - AUDIT TEST SUITE"
echo "=========================================="
echo ""

# Build the program first
echo "Building program..."
go build -o tetris-optimizer . 2>&1
if [ $? -ne 0 ]; then
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi
echo -e "${GREEN}Build successful!${NC}"
echo ""

# Bad Examples (should all output ERROR)
echo "=========================================="
echo "TESTING BAD EXAMPLES (Should output ERROR)"
echo "=========================================="
echo ""

run_test "Bad Example 00" "testdata/invalid/bad_example_00.txt" "ERROR"
run_test "Bad Example 01" "testdata/invalid/bad_example_01.txt" "ERROR"
run_test "Bad Example 02" "testdata/invalid/bad_example_02.txt" "ERROR"
run_test "Bad Example 03" "testdata/invalid/bad_example_03.txt" "ERROR"
run_test "Bad Example 04" "testdata/invalid/bad_example_04.txt" "ERROR"
run_test "Bad Format" "testdata/invalid/bad_format.txt" "ERROR"

# Good Examples (should all succeed with specific dot counts)
echo "=========================================="
echo "TESTING GOOD EXAMPLES"
echo "=========================================="
echo ""

run_test "Good Example 00" "testdata/valid/good_example_00.txt" "SUCCESS" "0"
run_test "Good Example 01" "testdata/valid/good_example_01.txt" "SUCCESS" "9"
run_test "Good Example 02" "testdata/valid/good_example_02.txt" "SUCCESS" "4"
run_test "Good Example 03" "testdata/valid/good_example_03.txt" "SUCCESS" "5"

# Hard Example
echo "=========================================="
echo "TESTING HARD EXAMPLE"
echo "=========================================="
echo ""

run_test "Hard Example" "testdata/valid/hard_example.txt" "SUCCESS" "1"

# Summary
echo "=========================================="
echo "TEST SUMMARY"
echo "=========================================="
echo -e "Total Tests: $((PASSED + FAILED))"
echo -e "${GREEN}Passed: $PASSED${NC}"
echo -e "${RED}Failed: $FAILED${NC}"
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}All tests passed! ✓${NC}"
    exit 0
else
    echo -e "${RED}Some tests failed! ✗${NC}"
    exit 1
fi
