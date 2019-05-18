# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

deps:
	go get -d -t ./...

fmt:
	go fmt $$(go list ./... )

vet:
	go vet $$(go list ./...)

test:
	bash scripts/test.sh
