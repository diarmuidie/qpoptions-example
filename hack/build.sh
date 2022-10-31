# Tell Ko to push to the local registry unless told otherwise
: ${KO_DOCKER_REPO:="ko.local"}
export KO_DOCKER_REPO

ko build --preserve-import-paths ./cmd/queue
