STEAMPIPE_INSTALL_DIR ?= ~/.steampipe

build:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/hudsonrock@latest/steampipe-plugin-hudsonrock.plugin *.go
