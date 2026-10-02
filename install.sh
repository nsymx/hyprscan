#!/usr/bin/env bash

BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

logger_info() {
	echo -e "${BLUE}[INFO]${NC} $1" >&2
}

logger_success() {
	echo -e "${GREEN}[SUCCESS]${NC} $1" >&2
}

logger_warn() {
	echo -e "${YELLOW}[WARN]${NC} $1" >&2
}

logger_error() {
	echo -e "${RED}[ERR]${NC} $1" >&2
}

detect_os() {
	local os
	os=$(uname -s | tr '[:upper:]' '[:lower:]')
	case $os in
	linux)
		echo "linux"
		;;
	*)
		logger_error "Unsupported OS: $os"
		logger_error "This installer currently supports Linux only"
		exit 1
		;;
	esac
}

detect_arch() {
	local arch
	arch=$(uname -m)
	case $arch in
	x86_64)
		echo "amd64"
		;;
	aarch64 | arm64)
		echo "arm64"
		;;
	*)
		logger_error "Unsupported architecture: $arch"
		logger_error "Supported architecture: x86_64 (amd64), aarch64 (arm64)"
		exit 1
		;;
	esac
}

detect_install_dir() {
	if [[ -z "$GOBIN" ]]; then
		if [[ -d "$HOME/.local/bin" ]]; then
			install_dir="${HOME}/.local/bin"
			logger_warn "\$GOBIN is not set, falling back to ${install_dir}"
		else
			logger_error "Please make sure you have \$GOBIN or ${HOME}/.local/bin set in your env/path"
			exit 1
		fi
	else
		install_dir="${GOBIN}"
	fi
}

download_and_verify() {
	binary_name="hyprscan"

	binary_name_full="${binary_name}-${project_version}-${os}-${arch}"
	checksum_file="sha512sum.txt"

	release_url="https://github.com/${project_name}/releases/download/${project_version}/${binary_name_full}"
	checksum_url="https://github.com/${project_name}/releases/download/${project_version}/${checksum_file}"

	logger_info "   Downloading binary        ${BLUE}[...]${NC}"

	if ! curl -sfL -o "${binary_name_full}" "${release_url}"; then
		logger_error "Failed to download binary"
		exit 1
	fi

	logger_info "   Downloading checksum file ${BLUE}[...]${NC}"

	if ! curl -sfL -o "${checksum_file}" "${checksum_url}"; then
		logger_warn "Could not download checksum file, skipping verification"
		return 0
	fi

	echo ""
	logger_info "   Verifying checksum        ${BLUE}[...]${NC}"

	if ! grep "  ${binary_name_full}$" "${checksum_file}" | sha512sum -c - --status; then
		logger_error "Checksum verification failed"
		rm -f "${binary_name_full}" "${checksum_file}"
		exit 1
	else
		logger_success "Checksum passed           ${GREEN}[ ✓ ]${NC}"
	fi

	return 0
}

install_binary() {
	if [[ ! -f "${binary_name_full}" ]]; then
		logger_error "Binary not found: ${binary_name_full}"
		exit 1
	fi

	chmod +x "${binary_name_full}"

	if [[ -w "${install_dir}" ]]; then
		if mv "${binary_name_full}" "${install_dir}/${binary_name}"; then
			if [[ -f "sha512sum.txt" ]]; then
				rm -f "sha512sum.txt"
			fi
		else
			logger_error "Failed to move binary to ${install_dir}"
			exit 1
		fi
	else
		logger_error "${install_dir}: permission denied... "
		exit 1
	fi
}

# Script entry
packages=(
	"curl"
	"echo"
	"go"
	"sha512sum"
)

for depends in "${packages[@]}"; do
	if ! command -v "${depends}"; then
		logger_error "Missing dependency: ${depends}"
		exit 1
	fi
done

main() {
	clear || exit

	os=$(detect_os)
	arch=$(detect_arch)

	project_name="nsymx/hyprscan"
	project_version=$(basename "$(curl -s -o /dev/null -w '%{redirect_url}' "https://github.com/${project_name}/releases/latest")")

	echo -e "${BLUE}"
	echo -e "::::::::::::::::::::::::::::::::::::::"
	echo -e "::::::::: Archutil Installer :::::::::"
	echo -e "::::::::::::::::::::::::::::::::::::::"
	echo -e "${NC}"

	if [[ -z "${project_version}" ]]; then
		logger_error "Failed to detect latest version"
		exit 1
	fi

	logger_info "   System                    :: ${os}"
	logger_info "   Architecture              :: ${arch}"
	echo ""

	logger_info "   Project Repository        :: ${project_name}"
	logger_info "   Project Version           :: ${project_version}"

	detect_install_dir
	echo ""

	download_and_verify
	install_binary

	echo ""
	logger_success "Installation complete     ${GREEN}[ ✓ ]${NC}"
	logger_success "No errors were reported   ${GREEN}[ ✓ ]${NC}"
}
main
