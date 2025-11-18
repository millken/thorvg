# Makefile for building and installing ThorVG with C API bindings, and building the Go project

.PHONY: all thorvg install clean go-build deps

FAKEROOT ?= $(CURDIR)/fakeroot

all: thorvg install go-build

deps:
	if [ "$$(uname -s)" = "Darwin" ]; then \
		brew install meson ninja libomp || true; \
	elif [ -f /etc/debian_version ]; then \
		sudo apt-get update; \
		sudo apt-get install -y meson ninja-build git build-essential pkg-config libomp-dev; \
	fi

thorvg:
# 	git clone https://github.com/thorvg/thorvg.git thorvg-src || true
	cd thorvg-src && meson setup build --prefix=$(FAKEROOT) -Ddefault_library=static -Dbindings=capi -Dsavers=all -Dengines=sw,gl --buildtype release --reconfigure || true
	cd thorvg-src && ninja -C build

install:
	cd thorvg-src && ninja -C build install

clean:
	rm -rf thorvg-src $(FAKEROOT)

# Build the Go project with static linking to ThorVG in fakeroot

go-build:
	CGO_CFLAGS="-I$(FAKEROOT)/include" CGO_LDFLAGS="-L$(FAKEROOT)/lib -lthorvg -lomp" go build