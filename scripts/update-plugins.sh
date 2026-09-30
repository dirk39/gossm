#!/usr/bin/env bash
set -euo pipefail

# Directory of this script and repository root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
ASSETS_DIR="${REPO_ROOT}/internal/assets/plugin"

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

echo "==> Downloading and updating AWS session-manager-plugin binaries..."

# 1. macOS ARM64 (Apple Silicon)
echo "--> Fetching macOS arm64..."
mkdir -p "${TMP_DIR}/mac_arm64" "${ASSETS_DIR}/darwin_arm64"
curl -sSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/mac_arm64/sessionmanager-bundle.zip" -o "${TMP_DIR}/mac_arm64.zip"
unzip -q "${TMP_DIR}/mac_arm64.zip" -d "${TMP_DIR}/mac_arm64"
cp "${TMP_DIR}/mac_arm64/sessionmanager-bundle/bin/session-manager-plugin" "${ASSETS_DIR}/darwin_arm64/session-manager-plugin"
chmod 0755 "${ASSETS_DIR}/darwin_arm64/session-manager-plugin"

# 2. macOS AMD64 (Intel)
echo "--> Fetching macOS amd64..."
mkdir -p "${TMP_DIR}/mac_amd64" "${ASSETS_DIR}/darwin_amd64"
curl -sSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/mac/sessionmanager-bundle.zip" -o "${TMP_DIR}/mac_amd64.zip"
unzip -q "${TMP_DIR}/mac_amd64.zip" -d "${TMP_DIR}/mac_amd64"
cp "${TMP_DIR}/mac_amd64/sessionmanager-bundle/bin/session-manager-plugin" "${ASSETS_DIR}/darwin_amd64/session-manager-plugin"
chmod 0755 "${ASSETS_DIR}/darwin_amd64/session-manager-plugin"

# 3. Linux AMD64
echo "--> Fetching Linux amd64..."
mkdir -p "${TMP_DIR}/linux_amd64" "${ASSETS_DIR}/linux_amd64"
curl -sSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_64bit/session-manager-plugin.deb" -o "${TMP_DIR}/linux_amd64.deb"
tar -xf "${TMP_DIR}/linux_amd64.deb" -C "${TMP_DIR}/linux_amd64" data.tar.gz
tar -xf "${TMP_DIR}/linux_amd64/data.tar.gz" -C "${TMP_DIR}/linux_amd64" ./usr/local/sessionmanagerplugin/bin/session-manager-plugin
cp "${TMP_DIR}/linux_amd64/usr/local/sessionmanagerplugin/bin/session-manager-plugin" "${ASSETS_DIR}/linux_amd64/session-manager-plugin"
chmod 0755 "${ASSETS_DIR}/linux_amd64/session-manager-plugin"

# 4. Linux ARM64
echo "--> Fetching Linux arm64..."
mkdir -p "${TMP_DIR}/linux_arm64" "${ASSETS_DIR}/linux_arm64"
curl -sSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_arm64/session-manager-plugin.deb" -o "${TMP_DIR}/linux_arm64.deb"
tar -xf "${TMP_DIR}/linux_arm64.deb" -C "${TMP_DIR}/linux_arm64" data.tar.gz
tar -xf "${TMP_DIR}/linux_arm64/data.tar.gz" -C "${TMP_DIR}/linux_arm64" ./usr/local/sessionmanagerplugin/bin/session-manager-plugin
cp "${TMP_DIR}/linux_arm64/usr/local/sessionmanagerplugin/bin/session-manager-plugin" "${ASSETS_DIR}/linux_arm64/session-manager-plugin"
chmod 0755 "${ASSETS_DIR}/linux_arm64/session-manager-plugin"

# 5. Windows AMD64
echo "--> Fetching Windows amd64..."
mkdir -p "${TMP_DIR}/windows_amd64" "${ASSETS_DIR}/windows_amd64"
curl -sSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/windows/SessionManagerPlugin.zip" -o "${TMP_DIR}/windows_amd64.zip"
unzip -q "${TMP_DIR}/windows_amd64.zip" -d "${TMP_DIR}/windows_amd64"
unzip -q "${TMP_DIR}/windows_amd64/package.zip" -d "${TMP_DIR}/windows_amd64/package"
cp "${TMP_DIR}/windows_amd64/package/bin/session-manager-plugin.exe" "${ASSETS_DIR}/windows_amd64/session-manager-plugin.exe"
chmod 0755 "${ASSETS_DIR}/windows_amd64/session-manager-plugin.exe"

# Generate SHA256 checksums
echo "==> Generating checksums..."
cd "${ASSETS_DIR}"
shasum -a 256 */session-manager-plugin* > "${ASSETS_DIR}/checksums.sha256"

echo "==> Verifying binary architectures:"
file */session-manager-plugin*

echo "==> Done!"
