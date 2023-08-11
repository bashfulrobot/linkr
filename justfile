# Docs
# load a .env file if in the directory
set dotenv-load
# Ignore recipe lines beginning with #.
set ignore-comments
# Search justfile in parent directory if the first recipe on the command line is not found.
set fallback

# "_" hides the recipie from listings
_default:
    @just --list --list-prefix ····

# Print `just` help
help:
    @just --help

# Create Linux cfg dir
create_linux_config_dir:
	mkdir -p {{env_var('HOME')}}/.config/linkr

# Create Mac cfg dir
create_darwin_config_dir:
	mkdir -p {{env_var('HOME')}}/Library/Application Support/linkr

# Copy Linux cfg file to cfg dir
copy_sample_config_linux:
	cp sample/config.yaml {{env_var('HOME')}}/.config/linkr/config.yaml

# Copy Mac cfg file to cfg dir
copy_sample_config_darwin:
	cp sample/config.yaml {{env_var('HOME')}}/Library/Application Support/linkr/config.yaml

# Build for current platform
build:
    go build -o bin/linkr main.go

# Cross compile for Linux
build-linux:
    GOOS=linux GOARCH=amd64 go build -o bin/linkr-linux main.go

# Cross compile for MacOS
build-darwin:
    GOOS=darwin GOARCH=amd64 go build -o bin/linkr-darwin main.go

# Cross compile for all platforms
build-all:
    mkdir -p bin
    # GOOS=linux GOARCH=amd64 go build -o bin/linkr-linux main.go
    # GOOS=darwin GOARCH=amd64 go build -o bin/linkr-darwin main.go
    just build-linux
    just build-darwin

# Rebuild linux binary and copy sample config
dev-rebuild:
    just clean
    just build
    just copy_sample_config_linux

# Run go mod tidy
tidy:
    go mod tidy

# Install on linux
install-linux:
    mkdir -p /home/dustin/.local/bin/
    cp bin/linkr /home/dustin/.local/bin/
    cp static/linkr-icon.png /home/dustin/.local/share/icons/hicolor/48x48/apps/
    cp static/linkr.desktop /home/dustin/.local/share/applications/
    sed -i 's/x-scheme-handler\/http=.*/x-scheme-handler\/http=linkr.desktop/g' ~/.config/mimeapps.list
    sed -i 's/x-scheme-handler\/https=.*/x-scheme-handler\/https=linkr.desktop/g' ~/.config/mimeapps.list
    sed -i 's/text\/html=.*/text\/html=linkr.desktop/g' ~/.config/mimeapps.list

# Deploy on linux (no dev)
deploy:
    just build
    just create_linux_config_dir
    just copy_sample_config_linux
    just install-linux
    just clean

# Clean up binaries
clean:
    rm -f ./bin/*
    rmdir ./bin
    go clean
