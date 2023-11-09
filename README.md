## File Santizer

**Still Work in Progress**

### Install

```
go install github.com/anfernee/sanitizer@latest
```

### dockerfile sanitizer

The dockerfile sanitizer tries to add digest to any reference to the docker image in `FROM`
command, to make the image reference more secure.

```
cat testdata/Dockerfile.sample | ./sanitizer dockerfile
FROM alpine:3.9.5@sha256:ab3fe83c0696e3f565c9b4a734ec309ae9bd0d74c192de4590fd6dc2ef717815

RUN apk add --no-cache \
    bash \
    curl \
    git \
    jq \
    zip

CMD ["/bin/bash"]
```
