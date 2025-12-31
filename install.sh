#!/bin/bash

# Mozi Installation Script
# This script downloads and installs the latest Mozi binary

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture
case $ARCH in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  *)
    echo -e "${RED}Error: Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

# Map OS
case $OS in
  linux)
    BINARY_NAME="mozi-linux-${ARCH}"
    ;;
  darwin)
    BINARY_NAME="mozi-darwin-${ARCH}"
    ;;
  *)
    echo -e "${RED}Error: Unsupported OS: $OS${NC}"
    echo "Please install manually from: https://github.com/Edison-A-N/mozi/releases"
    exit 1
    ;;
esac

# Get latest version
echo -e "${YELLOW}Fetching latest version...${NC}"
LATEST_VERSION=$(curl -s https://api.github.com/repos/Edison-A-N/mozi/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' || echo "")

if [ -z "$LATEST_VERSION" ]; then
  echo -e "${YELLOW}Warning: Could not fetch latest version, using 'latest'${NC}"
  DOWNLOAD_URL="https://github.com/Edison-A-N/mozi/releases/latest/download/${BINARY_NAME}"
else
  VERSION="$LATEST_VERSION"
  echo -e "${GREEN}Latest version: $VERSION${NC}"
  DOWNLOAD_URL="https://github.com/Edison-A-N/mozi/releases/download/${VERSION}/${BINARY_NAME}"
fi

# Temporary file
TMP_FILE=$(mktemp)
INSTALL_DIR="/usr/local/bin"
BINARY_PATH="${INSTALL_DIR}/mozi"

echo -e "${YELLOW}Downloading Mozi...${NC}"
echo "  URL: $DOWNLOAD_URL"
echo "  Target: $BINARY_PATH"

# Download binary
if ! curl -fsSL "$DOWNLOAD_URL" -o "$TMP_FILE"; then
  # If specific version failed and we have a version, try latest as fallback
  if [ -n "$VERSION" ] && [ "$VERSION" != "latest" ]; then
    echo -e "${YELLOW}Warning: Failed to download version $VERSION, trying latest...${NC}"
    FALLBACK_URL="https://github.com/Edison-A-N/mozi/releases/latest/download/${BINARY_NAME}"
    echo "  Fallback URL: $FALLBACK_URL"
    if ! curl -fsSL "$FALLBACK_URL" -o "$TMP_FILE"; then
      echo -e "${RED}Error: Failed to download binary${NC}"
      echo "Please check your internet connection and try again."
      rm -f "$TMP_FILE"
      exit 1
    fi
  else
    echo -e "${RED}Error: Failed to download binary${NC}"
    echo "Please check your internet connection and try again."
    rm -f "$TMP_FILE"
    exit 1
  fi
fi

# Make executable
chmod +x "$TMP_FILE"

# Check if mozi is already installed
if [ -f "$BINARY_PATH" ]; then
  echo -e "${YELLOW}Mozi is already installed at $BINARY_PATH${NC}"
  read -p "Do you want to overwrite it? (y/N): " -n 1 -r
  echo
  if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Installation cancelled."
    rm -f "$TMP_FILE"
    exit 0
  fi
fi

# Install binary (requires sudo)
echo -e "${YELLOW}Installing to $INSTALL_DIR...${NC}"
if sudo mv "$TMP_FILE" "$BINARY_PATH"; then
  echo -e "${GREEN}✓ Successfully installed Mozi!${NC}"
  echo ""
  echo "You can now use 'mozi' command:"
  echo "  mozi --help"
else
  echo -e "${RED}Error: Failed to install binary (permission denied)${NC}"
  echo "You may need to run this script with sudo or install manually:"
  echo "  sudo mv $TMP_FILE $BINARY_PATH"
  exit 1
fi

# Verify installation
if command -v mozi &> /dev/null; then
  echo ""
  echo -e "${GREEN}Verifying installation...${NC}"
  mozi --help | head -n 5
  echo ""
  echo -e "${GREEN}Installation complete!${NC}"
else
  echo -e "${YELLOW}Warning: mozi command not found in PATH${NC}"
  echo "Please make sure $INSTALL_DIR is in your PATH"
fi
