# wavs-go

[WAVS](https://wavs.xyz) go-lang bindings for [components](https://github.com/Lay3rLabs/wavs-foundry-template).

## Install Wit Bindgen for Go

```bash
go install go.bytecodealliance.org/cmd/wit-bindgen-go@ecfa620df5beee882fb7be0740959e5dfce9ae26

wit-bindgen-go --version
```

## System Setup

```bash
# https://component-model.bytecodealliance.org/language-support/go.html

# https://tinygo.org/getting-started/install/

# macOS
brew tap tinygo-org/tools
brew install tinygo

# Arch (btw)
sudo pacman -Sy tinygo

# Ubuntu / WSL:
# TODO: .


# -------

# verify installs
tinygo version
wkg --version

# move into the golang oracle directory

# download the WAVS package bindings
export WAVS_PACKAGE=wavs:worker@0.3.0-rc1
# TODO: this is currently broken on this release, requires:
# TODO: https://github.com/Lay3rLabs/WAVS/pull/403 to fix `failed to resolve import `wasi:cli/environment@0.2.0::get-environment`

cp /home/reece/Desktop/Programming/Rust/wavs/sdk/wavs:worker@0.3.0-rc1.wasm .
# wkg get $WAVS_PACKAGE --overwrite --format wasm --output ${WAVS_PACKAGE}.wasm

# generate the Go/ bindings
# if `error: error executing wasm-tools: module closed with exit_code(1)`, set WAVS_PACKAGE
wit-bindgen-go generate -o . ${WAVS_PACKAGE}.wasm

go mod tidy
```
