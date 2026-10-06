#!/bin/bash

# Docker Build Script for Payment Gateway Services
# This script builds images for all services using the available
# container engine: `docker` if installed, otherwise `podman` (native).

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect the container engine: prefer docker, fall back to podman (native).
detect_engine() {
    if command -v docker &> /dev/null; then
        echo "docker"
    elif command -v podman &> /dev/null; then
        echo "podman"
    else
        echo ""
    fi
}

ENGINE="${ENGINE:-$(detect_engine)}"

# Check that the container engine is installed and running
check_engine() {
    if [ -z "$ENGINE" ]; then
        print_error "Neither Docker nor Podman is installed. Please install one of them first."
        exit 1
    fi

    print_status "Using container engine: ${ENGINE}"

    if ! ${ENGINE} info &> /dev/null; then
        print_error "${ENGINE} is not running or not accessible."
        exit 1
    fi
    print_status "${ENGINE} is installed and running"
}

# Build image for a service
build_service_image() {
    local service="$1"

    if [ -z "$service" ]; then
        print_error "Service name is required"
        return 1
    fi

    local service_dir="service/${service}"
    local dockerfile="${service_dir}/Dockerfile"
    local image="monolith-payment-gateway-grpc/${service}:latest"

    print_status "Building ${service} service image..."

    if [ ! -d "$service_dir" ]; then
        print_error "Service directory not found: $service_dir"
        return 1
    fi

    if [ ! -f "$dockerfile" ]; then
        print_error "Dockerfile not found: $dockerfile"
        return 1
    fi

    if [ ! -f "${service_dir}/go.mod" ]; then
        print_error "go.mod not found in service: $service_dir"
        return 1
    fi

    if ${ENGINE} build \
        --progress=plain \
        -f "$dockerfile" \
        -t "$image" \
        .; then

        print_status "Successfully built ${image}"
        return 0
    else
        print_error "Failed to build ${image}"
        return 1
    fi
}

# Build all service images
build_all_images() {
    print_status "Building all service images..."
    
    # List of services to build
    services=("auth" "user" "card" "merchant" "role" "saldo" "transaction" "topup" "transfer" "withdraw" "apigateway" "email" "migrate")
    
    local failed_builds=0
    
    for service in "${services[@]}"; do
        if ! build_service_image "$service"; then
            ((failed_builds++))
        fi
    done
    
    if [ $failed_builds -eq 0 ]; then
        print_status "🎉 All images built successfully!"
    else
        print_warning "⚠️  $failed_builds images failed to build"
        return 1
    fi
}

# Show built images
show_built_images() {
    print_status "Built images:"
    echo ""
    ${ENGINE} images | grep -E "(auth|user|card|merchant|role|saldo|transaction|topup|transfer|withdraw|apigateway)" | head -20
    echo ""
}

# Cleanup function
cleanup() {
    print_status "Build process completed"
}

# Main execution
main() {
    # Set trap for cleanup
    trap cleanup EXIT
    
    # Run checks
    check_engine
    
    # Build all images
    build_all_images
    
    # Show built images
    show_built_images
    
    print_status "Build process completed! 🎉"
}

# Handle script interruption
trap 'print_error "Build interrupted"; exit 1' INT

# Run main function
main "$@"
