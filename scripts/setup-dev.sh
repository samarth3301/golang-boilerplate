#!/bin/bash

# Development setup script for golang-boilerplate
# This script installs development tools and dependencies

echo "🚀 Setting up development environment for golang-boilerplate..."

# Install reflex for live reloading (like nodemon)
echo "📦 Installing reflex for live reloading..."
go install github.com/cespare/reflex@latest

# Install golangci-lint for code linting
echo "📦 Installing golangci-lint..."
# Check if brew is available (macOS)
if command -v brew &> /dev/null; then
    brew install golangci-lint
else
    # Fallback to go install
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Install air as alternative (if Go version supports it)
echo "📦 Installing air (alternative live reloader)..."
go install github.com/air-verse/air@latest 2>/dev/null || echo "⚠️  Air requires Go 1.25+, skipping..."

echo "✅ Development environment setup complete!"
echo ""
echo "Available commands:"
echo "  make watch     - Start with live reloading (reflex)"
echo "  make run       - Run without live reloading"
echo "  make test      - Run tests"
echo "  make lint      - Run linter"
echo "  make build     - Build application"