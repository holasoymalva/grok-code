#!/bin/bash
set -e

REPO="holasoymalva/grok-code"

echo "Installing Grok Code..."

# Detect OS and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux)  OS_NAME="Linux" ;;
    Darwin) OS_NAME="Darwin" ;;
    *)      echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
    x86_64)  ARCH_NAME="x86_64" ;;
    arm64)   ARCH_NAME="arm64" ;;
    aarch64) ARCH_NAME="arm64" ;;
    *)       echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Formulate the download URL based on GitHub Releases (assumes GoReleaser format)
# In a real scenario, this would fetch the latest release tag via the GitHub API
LATEST_TAG=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo "Could not fetch latest release. Please ensure the repository is public and has releases."
    exit 1
fi

BINARY_URL="https://github.com/$REPO/releases/download/${LATEST_TAG}/grok-code_${OS_NAME}_${ARCH_NAME}.tar.gz"

echo "Downloading Grok Code $LATEST_TAG for $OS_NAME ($ARCH_NAME)..."
curl -sL "$BINARY_URL" -o /tmp/grok-code.tar.gz

echo "Extracting binary..."
tar -xzf /tmp/grok-code.tar.gz -C /tmp grok-code

echo "Installing to /usr/local/bin (may require sudo)..."
sudo mv /tmp/grok-code /usr/local/bin/grokcode
sudo chmod +x /usr/local/bin/grokcode

echo "Cleaning up..."
rm /tmp/grok-code.tar.gz

echo ""
echo "✅ Grok Code installed successfully!"
echo "Run 'grokcode chat' to get started."
