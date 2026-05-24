#!/usr/bin/env bash
set -e

echo "==> Initializing workspace..."

# Initialize Go module if not already present
if [ ! -f /workspace/go.mod ]; then
  echo "==> Creating Go module (edit module path as needed)"
  cd /workspace && go mod init app
fi

# Install npm dependencies if package.json exists
if [ -f /workspace/package.json ]; then
  echo "==> Installing npm dependencies..."
  cd /workspace && npm install
fi

echo "==> Done. Happy coding!"
