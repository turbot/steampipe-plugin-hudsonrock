STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

build:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/hudsonrock@latest/steampipe-plugin-hudsonrock.plugin -tags "${BUILD_TAGS}" *.go

