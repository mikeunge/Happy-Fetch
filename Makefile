CC         = go
BUILD_PATH = ./bin
SRC        = main.go
TARGET     = happy-msg
BINS       = $(BUILD_PATH)/$(TARGET)
INST       = /usr/local/bin

.PHONY: all clean build run install

all: run

clean:
	rm -rf $(BUILD_PATH)

build: clean
	mkdir -p $(BUILD_PATH)
	$(CC) build -o $(BINS) $(SRC)

run: build
	$(BINS) --help

install: build 
	sudo cp -v $(BINS) $(INST)

