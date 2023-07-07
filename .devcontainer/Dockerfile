FROM debian:bookworm

ARG USER_NAME=vscode \
  USER_UID=1000 \
  USER_GID=1000

ARG GOLANG_VERSION=1.19.10 \
  GOLANG_SHASUM_amd64=8b045a483d3895c6edba2e90a9189262876190dbbd21756870cdd63821810677 \
  GOLANG_SHASUM_arm64=df98698821211c819e8b2420c77a0f802d989e377718578a31b1f91f6be2c5b4

ADD docker.asc /etc/apt/keyrings/
RUN export DEBIAN_FRONTEND=noninteractive ARCH=$(dpkg --print-architecture) \
  && apt-get update \
  && apt-get -y install --no-install-recommends \
  apt-transport-https \
  ca-certificates \
  curl \
  git \
  gnupg \
  locales \
  lsb-release \
  make \
  openssh-client \
  unzip \
  && groupadd --gid $USER_GID $USER_NAME \
  && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USER_NAME \
  && curl -fsSL -o /tmp/golang.tar.gz https://golang.org/dl/go${GOLANG_VERSION}.linux-${ARCH}.tar.gz \
  && bash -c "echo \${GOLANG_SHASUM_${ARCH}}  /tmp/golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf /tmp/golang.tar.gz \
  && rm /tmp/golang.tar.gz \
  && echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh \
  && su -lc "echo PATH=\$HOME/go/bin:\$PATH >> \$HOME/.profile" ${USER_NAME} \
  && su -lc "go install golang.org/x/tools/cmd/goimports@latest" ${USER_NAME} \
  # Docker
  && echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.asc] \
  https://download.docker.com/linux/debian bookworm stable" > /etc/apt/sources.list.d/docker.list \
  && apt-get update \
  && apt-get install -y docker-ce-cli
