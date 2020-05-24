FROM ubuntu:xenial


# apt packages
ENV INSTALL_DEPS \
  build-essential \
  ca-certificates \
  git \
  make \
  software-properties-common \
  unzip \
  wget \
  maven \
  patch \
  python
RUN apt-get update \
  && apt-get install -y -q --no-install-recommends curl openjdk-8-jdk \
  && apt-get update \
  && apt-get install -y -q --no-install-recommends ${INSTALL_DEPS} \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# protoc
ENV PROTOC_VER=3.12.0
ENV PROTOC_REL=protoc-"${PROTOC_VER}"-linux-x86_64.zip
RUN wget https://github.com/google/protobuf/releases/download/v"${PROTOC_VER}/${PROTOC_REL}" \
  && unzip ${PROTOC_REL} -d protoc \
  && mv protoc /usr/local \
  && ln -s /usr/local/protoc/bin/protoc /usr/local/bin


# go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GORELEASE go1.14.3.linux-amd64.tar.gz
RUN wget -q https://dl.google.com/go/$GORELEASE \
    && tar -C $(dirname $GOROOT) -xzf $GORELEASE \
    && rm $GORELEASE \
    && mkdir -p $GOPATH/{src,bin,pkg}

# protoc-gen-go
ENV PGG_PKG "github.com/golang/protobuf/protoc-gen-go"
ENV PGG_PATH "${GOPATH}/src/${PGG_PKG}"
ENV PGG_VER=v1.4.2
RUN go get -d ${PGG_PKG} \
  && cd ${PGG_PATH} \
  && git checkout ${PGG_VER} \
  && go install \
  && cd - \
  && rm -rf ${PGG_PATH}

# deps
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# buildozer
RUN go get github.com/bazelbuild/buildtools/buildozer
RUN go get github.com/bazelbuild/bazelisk && mv $(go env GOPATH)/bin/bazelisk $(go env GOPATH)/bin/bazel && export PATH=$PATH:$(go env GOPATH)/bin

WORKDIR ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
COPY . .

RUN make build

ENTRYPOINT ["make"]
CMD ["build"]
