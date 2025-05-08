#!/bin/sh

OS="$(uname -s)"
ARCH="$(uname -m)"
REPO="zaidfadhil/xwag"
INSTALL_DIR="/usr/local/bin"
LOCAL_BINARY="bin/xwag"

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

case "$OS" in
    Linux)
        case "$ARCH" in
            x86_64)
                SUFFIX="linux_amd64"
                ;;
            aarch64)
                SUFFIX="linux_arm64"
                ;;
            *)
                echo "Unsupported architecture: $ARCH"
                exit 1
                ;;
        esac
        ;;
    Darwin)
        case "$ARCH" in
            x86_64)
                SUFFIX="darwin_amd64"
                ;;
            arm64)
                SUFFIX="darwin_arm64"
                ;;
            *)
                echo "Unsupported architecture: $ARCH"
                exit 1
                ;;
        esac
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
      ;;
esac

if [ -f "$LOCAL_BINARY" ]; then
  echo "Local binary found. Using it for installation."
  BINARY_PATH="$LOCAL_BINARY"
else
  RELEASE_URL="https://api.github.com/repos/$REPO/releases/latest"
  echo "Fetching the latest release tag..."
  RELEASE_TAG=$(curl -s $RELEASE_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

  if [ -z "$RELEASE_TAG" ]; then
    echo "Failed to fetch the release tag."
    exit 1
  fi

  DOWNLOAD_URL="https://github.com/$REPO/releases/download/$RELEASE_TAG/xwag_${RELEASE_TAG#v}_${SUFFIX}.tar.gz"

  echo "Downloading the latest release..."
  curl -L -o xwag.tar.gz "$DOWNLOAD_URL"

  TMP_DIR=$(mktemp -d)
  tar -xzf xwag.tar.gz -C "$TMP_DIR"
  BINARY_PATH="$TMP_DIR/xwag"
  chmod +x "$BINARY_PATH"
fi

echo "Installing the application..."
sudo mv "$BINARY_PATH" "$INSTALL_DIR" || exit 1

if command_exists xwag; then
  echo "xwag was successfully installed!"
else
  echo "Failed to install xwag."
  exit 1
fi

if [ -f "xwag.tar.gz" ]; then
  rm xwag.tar.gz
fi
if [ -f "xwag" ]; then
  rm xwag
fi

echo "Installation complete."
