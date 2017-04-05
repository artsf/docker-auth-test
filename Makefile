.PHONY: all setupmac clean binary rootfs create enable
FILE_NAME=docker-auth-test
PLUGIN_SHORT_NAME=docker-auth-test
PLUGIN_NAME=test/docker-auth-test
PLUGIN_TAG=next
PLUGIN_FOLDER=./plugin

all: clean binary rootfs create enable

setupmac:
	@echo
	@echo
	@echo ================================================
	@echo  Setup...
	@echo ================================================
	@echo
	brew reinstall go --with-cc-all
	brew install glide

clean:
	@echo
	@echo
	@echo ================================================
	@echo  Cleanup...
	@echo ================================================
	@echo
	rm -f ${FILE_NAME}
	chmod -R +w ${PLUGIN_FOLDER} || true
	rm -rf ${PLUGIN_FOLDER}

binary:
	@echo
	@echo
	@echo ================================================
	@echo  Building binary...
	@echo ================================================
	@echo
	glide install
	go build -o ${FILE_NAME} --ldflags '-extldflags "-static"'

rootfs:
	@echo
	@echo
	@echo ================================================
	@echo  Preparing files...
	@echo ================================================
	@echo
	@echo "### create rootfs directory in ${PLUGIN_FOLDER}/rootfs"
	mkdir -p ${PLUGIN_FOLDER}/rootfs
	@echo "### copy config.json to .${PLUGIN_FOLDER}/"
	cp config.json ${PLUGIN_FOLDER}/
	cp ${FILE_NAME} ${PLUGIN_FOLDER}/rootfs/

create:
	@echo
	@echo
	@echo ================================================
	@echo  Building docker plugin from file system...
	@echo ================================================
	@echo
	@echo "### disable plugin ${PLUGIN_NAME}:${PLUGIN_TAG} if exists"
	docker plugin disable ${PLUGIN_NAME}:${PLUGIN_TAG} || true
	@echo "### remove existing plugin ${PLUGIN_NAME}:${PLUGIN_TAG} if exists"
	docker plugin rm -f ${PLUGIN_NAME}:${PLUGIN_TAG} || true
	@echo "### create new plugin ${PLUGIN_NAME}:${PLUGIN_TAG} from ${PLUGIN_FOLDER}"
	docker plugin create ${PLUGIN_NAME}:${PLUGIN_TAG} ${PLUGIN_FOLDER}

enable:
	@echo
	@echo
	@echo ================================================
	@echo  Enabling plugin...
	@echo ================================================
	@echo
	@echo "### enable plugin ${PLUGIN_NAME}:${PLUGIN_TAG}"
	docker plugin enable ${PLUGIN_NAME}:${PLUGIN_TAG}
	@echo
	@echo
	@echo ================================================
	@echo  SUCCESS! Plugin is deployed and enabled.
	@echo ================================================
