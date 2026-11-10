CC ?= cc
CPPFLAGS += $(shell pkg-config --cflags json-c sqlite3)
CFLAGS += -std=c11 -Wall -Wextra -Werror -pedantic
LDLIBS += $(shell pkg-config --libs json-c sqlite3)

.PHONY: setup
setup: build/setup
	./build/setup

build/setup: setup.c
	mkdir -p build
	$(CC) $(CPPFLAGS) $(CFLAGS) $< -o $@ $(LDLIBS)
