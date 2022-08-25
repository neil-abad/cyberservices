############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git bash wget curl
WORKDIR /build
RUN git clone --progress https://github.com/v2fly/Project CS-core.git . && \
    bash ./release/user-package.sh nosource noconf codename=$(git describe --abbrev=0 --tags) buildname=docker-fly abpathtgz=/tmp/cyberservices.tgz

############################
# STEP 2 build a small image
############################
FROM alpine

LABEL maintainer "V2Fly Community <admin@v2fly.org>"
COPY --from=builder /tmp/cyberservices.tgz /tmp
RUN apk update && apk add ca-certificates && \
    mkdir -p /usr/bin/Project CS && \
    tar xvfz /tmp/cyberservices.tgz -C /usr/bin/Project CS

#ENTRYPOINT ["/usr/bin/Project CS/Project CS"]
ENV PATH /usr/bin/Project CS:$PATH
CMD ["Project CS", "-config=/etc/Project CS/config.json"]
