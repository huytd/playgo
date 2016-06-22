# Shameless copy paste from tj/luna project
# https://github.com/tj/luna/blob/master/Makefile

SRC = $(wildcard *.go)
OUT = playgo 

# Fallback to gcc if clang not available
CC = go
go = $(shell which go 2> /dev/null)

ifeq (, $(go))
	@printf "\e[91mGo not found!"
endif

$(OUT): clean $(SRC)
	@printf "\e[33mBuilding\e[90m %s\e[0m\n" $@
	@$(CC) build -o $@ $(SRC)
	@printf "\e[34mDone!\e[0m\n"

test: clean
	@printf "\e[33mTesting...\e[0m\n"
	go test ./...
	@printf "\e[34mDone!\e[0m\n"

clean:
	@rm -f $(OUT)
	@printf "\e[34mAll clear!\e[0m\n"

install: $(OUT)
	@printf "\e[33mInstalling\e[90m %s\e[0m\n" $(OUT)
	sudo rm -f /usr/local/bin/$(OUT)
	sudo ln -s $(PWD)/$(OUT) /usr/local/bin/$(OUT)
	@printf "\e[34mDone!\e[0m\n"

uninstall:
	@printf "\e[33mRemoving\e[90m %s\e[0m\n" $(OUT)
	sudo rm -f /usr/local/bin/$(OUT)
	@printf "\e[34mDone!\e[0m\n"
