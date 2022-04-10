REGISTRY=ghcr.io/thelonkarrde/commentci
TAG=v0.3.0

.PHONY: docker
docker:
	docker build -t ${REGISTRY}:${TAG} .