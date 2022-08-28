
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
PROJ_DIR := ${ROOT_DIR}proj
JSON_DIR := ${PROJ_DIR}/jsonFiles
OSS_DIR := ${PROJ_DIR}/oss
PROP_DIR := ${PROJ_DIR}/prop
EXE_NAME := sep
BIN := bin

ifeq ($(OS),Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
    RM_RF_CMD = ${RM_F_CMD} -Recurse
	exe =${BIN}/${EXE_NAME}.exe
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
else
	SHELL := bash
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	exe =${BIN}/${EXE_NAME}
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
endif

all: clean build test
.DEFAULT_GOAL := help
.PHONY: clean build test all git help

all: $(PROJ_DIR) ## performs clean build and test
clean: $@ ## Move back the files to original state in test folder
build: $@ ## Generate the windows and linux builds for sep
test: $@ ## Separates the test folder
git: $@ ## commits and push the changes if commit msg m is given without spaces ex m=added_files

build:
	echo "Compiling for every OS and Platform"
	set GOOS=windows
	set GOARCH=arm64
	go build -o ${BIN}/${EXE_NAME}.exe ${EXE_NAME}.go
	set GOOS=linux
	set GOARCH=arm64
	go build -o ${BIN}/${EXE_NAME} ${EXE_NAME}.go


test:
	echo "===========Testing==============="
	${exe} ${PROJ_DIR}

del:
	${RM_RF_CMD} ${JSON_DIR}
	${RM_RF_CMD} ${OSS_DIR}
	${RM_RF_CMD} ${PROP_DIR}
	${RM_RF_CMD} bin/*

mv:
	mv ${JSON_DIR}/* ${PROJ_DIR}/
	mv ${OSS_DIR}/* ${PROJ_DIR}/
	mv ${PROP_DIR}/* ${PROJ_DIR}/

clean: mv del


git:
	git status
	git add .
	git status
	git commit -m ${m}
	git push

help: ## Show this help
	@${HELP_CMD}
