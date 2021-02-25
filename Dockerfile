FROM ubuntu:xenial


# apt packages
ENV INSTALL_DEPS \
  bazel \
  ca-certificates \
  git \
  make \
  software-properties-common \
  unzip \
  wget \
  maven \
  patch \
  python3.8 \
  python3.8-distutils \
  python3-setuptools
RUN apt-get update \
  && apt-get install -y -q --no-install-recommends curl openjdk-8-jdk \
  && echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list \
  && curl https://bazel.build/bazel-release.pub.gpg | apt-key add - \
  # software-properties-common allows adding PPA (Personal Package Archive) repositories
  && apt install -y -q --no-install-recommends software-properties-common \
  # deadsnakes is a PPA with newer releases of python than default Ubuntu repositories
  && add-apt-repository ppa:deadsnakes/ppa \
  && apt-get update \
  && apt-get install -y -q --no-install-recommends ${INSTALL_DEPS} \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# protoc
ENV PROTOC_VER=3.12.4
ENV PROTOC_REL=protoc-"${PROTOC_VER}"-linux-x86_64.zip
RUN wget https://github.com/google/protobuf/releases/download/v"${PROTOC_VER}/${PROTOC_REL}" \
  && unzip ${PROTOC_REL} -d protoc \
  && mv protoc /usr/local \
  && ln -s /usr/local/protoc/bin/protoc /usr/local/bin


# go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GORELEASE go1.14.7.linux-amd64.tar.gz
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

WORKDIR ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
COPY . .

# python must be on PATH for the execution of py_binary bazel targets, but
# the distribution we installed doesn't provide this alias
RUN ln -s /usr/bin/python3.8 /usr/bin/python

# python tooling for linting and uploading to PyPI
RUN python3.8 -m easy_install pip \
  && python3.8 -m pip install -r requirements.txt

RUN make build

ENTRYPOINT ["make"]
CMD ["build"]
