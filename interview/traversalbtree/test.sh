docker run -it -v $(PWD):/go/src/traversalbtree --rm golang:1.12 bash -c "cd /go/src/traversalbtree && go test ./... --race"