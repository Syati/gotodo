ROOTDIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: migrate

migrate:
    migrate up -source ${ROOTDIR}/db/migrate