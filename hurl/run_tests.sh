#!/bin/bash

# Hurl API Test Runner
# This script runs all Hurl test files for the payment gateway API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Base URL for the API. Override with BASE_URL when the gateway is exposed
# on another host/port (for example, http://localhost:5000).
BASE_URL="${BASE_URL:-http://localhost:5000}"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if hurl is installed
check_hurl() {
    if ! command -v hurl &> /dev/null; then
        print_error "Hurl is not installed. Please install it first:"
        echo "Visit: https://hurl.dev/docs/installation.html"
        exit 1
    fi
}

# Check if API is running
check_api() {
    print_status "Checking if API Gateway is running on $BASE_URL..."
    if curl -fsS -o /dev/null "$BASE_URL/health"; then
        print_success "API Gateway is running!"
    else
        print_error "API Gateway is not responding on $BASE_URL"
        print_warning "Please start the API Gateway first"
        exit 1
    fi
}

# Run a single test file. Existing fixtures contain both :5000 and :8080
# literals, so normalize them into a temporary copy instead of requiring every
# fixture to be edited when the gateway port changes.
run_test_file() {
    local file=$1
    local rendered
    local escaped_base_url
    rendered=$(mktemp "${TMPDIR:-/tmp}/hurl.XXXXXX.hurl")
    escaped_base_url=$(printf '%s' "$BASE_URL" | sed 's/[&|\\]/\\&/g')
    sed -E "s|https?://localhost:[0-9]+|${escaped_base_url}|g" "$file" > "$rendered"

    print_status "Running tests in $file..."
    if grep -Fq '{{auth_token}}' "$rendered" && [[ -z "${AUTH_TOKEN:-}" ]]; then
        print_error "AUTH_TOKEN is required by fixture $file"
        rm -f "$rendered"
        return 1
    fi
    if grep -Fq '{{refresh_token}}' "$rendered" && [[ -z "${REFRESH_TOKEN:-}" ]]; then
        print_error "REFRESH_TOKEN is required by fixture $file"
        rm -f "$rendered"
        return 1
    fi
    if grep -Fq '{{reset_token}}' "$rendered" && [[ -z "${RESET_TOKEN:-}" ]]; then
        print_error "RESET_TOKEN is required by fixture $file"
        rm -f "$rendered"
        return 1
    fi
    if hurl \
        --variable BASE_URL="$BASE_URL" \
        --variable baseUrl="$BASE_URL" \
        --variable auth_token="${AUTH_TOKEN:-}" \
        --variable refresh_token="${REFRESH_TOKEN:-}" \
        --variable reset_token="${RESET_TOKEN:-}" \
        --variable accessToken="${AUTH_TOKEN#Bearer }" \
        --variable refreshToken="${REFRESH_TOKEN:-}" \
        --variable userId="${USER_ID:-1}" \
        --variable cardNumber="${CARD_NUMBER:-4532123456789012}" \
        --variable merchantId="${MERCHANT_ID:-1}" \
        --variable saldoId="${SALDO_ID:-1}" \
        --variable apiKey="${API_KEY:-}" \
        --variable testEmail="${TEST_EMAIL:-hurl.$$.${RANDOM}@example.com}" \
        "$rendered"; then
        print_success "All tests in $file passed!"
        rm -f "$rendered"
        return 0
    else
        print_error "Some tests in $file failed!"
        rm -f "$rendered"
        return 1
    fi
}

# Main execution
main() {
    echo "========================================"
    echo "  Payment Gateway API Test Runner"
    echo "========================================"
    echo

    # Check prerequisites
    check_hurl
    check_api

    echo
    print_status "Starting API tests..."
    echo

    # Run the maintained canonical collection only. The root hurl/ files are
    # legacy placeholders with obsolete routes and must not be mixed into this
    # suite; doing so creates false green 200 responses and state collisions.
    mapfile -t test_files < <(find tests/hurl -type f -name '*.hurl' -print | sort)
    failed_tests=()
    passed_tests=()
    runs="${HURL_RUNS:-1}"

    for ((run=1; run<=runs; run++)); do
        print_status "Starting Hurl pass $run/$runs..."
        for file in "${test_files[@]}"; do
            if [[ -f "$file" ]]; then
                echo "----------------------------------------"
                if run_test_file "$file"; then
                    passed_tests+=("$file (run $run)")
                else
                    failed_tests+=("$file (run $run)")
                fi
                echo
            fi
        done
    done

    # Summary
    echo "========================================"
    echo "  Test Summary"
    echo "========================================"
    echo

    if [[ ${#passed_tests[@]} -gt 0 ]]; then
        print_success "Passed tests (${#passed_tests[@]}):"
        for file in "${passed_tests[@]}"; do
            echo "  ✓ $file"
        done
        echo
    fi

    if [[ ${#failed_tests[@]} -gt 0 ]]; then
        print_error "Failed tests (${#failed_tests[@]}):"
        for file in "${failed_tests[@]}"; do
            echo "  ✗ $file"
        done
        echo
    fi

    # Exit with appropriate code
    if [[ ${#failed_tests[@]} -eq 0 ]]; then
        print_success "All tests passed! 🎉"
        exit 0
    else
        print_error "Some tests failed. Please check the output above."
        exit 1
    fi
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [options]"
        echo
        echo "Options:"
        echo "  --help, -h     Show this help message"
        echo "  --check         Only check prerequisites (hurl installation and API status)"
        echo
        echo "Examples:"
        echo "  $0              # Run all tests"
        echo "  $0 --check      # Check prerequisites only"
        exit 0
        ;;
    --check)
        check_hurl
        check_api
        print_success "All prerequisites met!"
        exit 0
        ;;
    "")
        main
        ;;
    *)
        print_error "Unknown option: $1"
        echo "Use --help for usage information"
        exit 1
        ;;
esac