FROM ubuntu:focal

# apt packages
ENV INSTALL_DEPS \
  ca-certificates \
  git \
  make \
  zip \
  unzip \
  g++ \
  wget \
  maven \
  patch \
  python3 \
  python3-distutils \
  python3-setuptools \
  apt-transport-https \
  curl \
  openjdk-8-jdk \
  gnupg 

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update \
  && apt install -y -q --no-install-recommends ${INSTALL_DEPS} \
  && apt clean

RUN wget -O bazel https://github.com/bazelbuild/bazel/releases/download/5.3.0/bazel-5.3.0-linux-arm64 && chmod +x bazel
RUN mv bazel usr/local/bin/bazel

# protoc
ENV PROTOC_VER=21.5
ENV PROTOC_REL=protoc-"${PROTOC_VER}"-linux-aarch_64.zip
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v"${PROTOC_VER}/${PROTOC_REL}" \
  && unzip ${PROTOC_REL} -d protoc \
  && mv protoc /usr/local \
  && ln -s /usr/local/protoc/bin/protoc /usr/local/bin

# go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GORELEASE go1.17.linux-arm64.tar.gz
RUN wget -q https://dl.google.com/go/$GORELEASE \
  && tar -C $(dirname $GOROOT) -xzf $GORELEASE \
  && rm $GORELEASE \
  && mkdir -p $GOPATH/{src,bin,pkg}

# protoc-gen-go
ENV PGG_VER=v1.28.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PGG_VER}

# buildozer
ENV BDR_VER=5.1.0
RUN go install github.com/bazelbuild/buildtools/buildozer@${BDR_VER}

WORKDIR ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate

COPY . .

RUN go get ./...

# python must be on PATH for the execution of py_binary bazel targets, but
# the distribution we installed doesn't provide this alias
RUN ln -s /usr/bin/python3.8 /usr/bin/python

# python tooling for linting and uploading to PyPI
RUN python3.8 -m easy_install pip \
  && python3.8 -m pip install -r requirements.txt

RUN make build

ENTRYPOINT ["make"]
CMD ["build"]