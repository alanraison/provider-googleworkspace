FROM debian:bookworm

ARG USER_NAME=vscode \
  USER_UID=1000 \
  USER_GID=1000

ARG GOLANG_VERSION=1.19.11 \
  GOLANG_SHASUM_amd64=ee18f98a03386e2bf48ff75737ea17c953b1572f9b1114352f104ac5eef04bb4 \
  GOLANG_SHASUM_arm64=ae22c047e0e63d2d28205b529baaf9d9ca0c93e890c309af62cd116b9efebcbd

ARG UP_VERSION=v0.18.0

ADD docker.asc \
  kubernetes-archive-keyring.asc \
  cloud.google.asc \
  hashicorp-archive-keyring.asc \
  /etc/apt/keyrings/

RUN export DEBIAN_FRONTEND=noninteractive ARCH=$(dpkg --print-architecture) \
  && apt-get update \
  && apt-get install -y --no-install-recommends \
  apt-transport-https \
  ca-certificates \
  # Docker
  && echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.asc] \
  https://download.docker.com/linux/debian bookworm stable" > /etc/apt/sources.list.d/docker.list \
  # Kubectl
  && echo "deb [signed-by=/etc/apt/keyrings/kubernetes-archive-keyring.asc] https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list \
  # gke-gcloud-auth-plugin
  && echo "deb [signed-by=/etc/apt/keyrings/cloud.google.asc] https://packages.cloud.google.com/apt cloud-sdk main" > /etc/apt/sources.list.d/google-cloud-sdk.list \
  # Terraform
  && echo "deb [signed-by=/etc/apt/keyrings/hashicorp-archive-keyring.asc] https://apt.releases.hashicorp.com bookworm main" > /etc/apt/sources.list.d/hashicorp.list \
  && apt-get update \
  && apt-get -y install --no-install-recommends \
  apt-transport-https \
  build-essential \
  ca-certificates \
  curl \
  docker-ce-cli \
  google-cloud-cli \
  google-cloud-cli-gke-gcloud-auth-plugin \
  git \
  kubectl \
  locales \
  make \
  openssh-client \
  python3 \
  terraform \
  unzip \
  && groupadd --gid $USER_GID $USER_NAME \
  && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USER_NAME \
  # Go
  && curl -fsSL -o /tmp/golang.tar.gz https://golang.org/dl/go${GOLANG_VERSION}.linux-${ARCH}.tar.gz \
  && bash -c "echo \${GOLANG_SHASUM_${ARCH}}  /tmp/golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf /tmp/golang.tar.gz \
  && rm /tmp/golang.tar.gz \
  && echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh \
  && su -lc "echo PATH=\$HOME/go/bin:\$PATH >> \$HOME/.profile" ${USER_NAME} \
  && su -lc "go install golang.org/x/tools/cmd/goimports@latest" ${USER_NAME} \
  # Up
  && curl -fsSL -o /tmp/up.tar.gz "https://cli.upbound.io/stable/${UP_VERSION}/bundle/up/linux_${ARCH}.tar.gz" \
  && tar -xzf /tmp/up.tar.gz -C /usr/local/bin --strip-components=1 \
  && rm /tmp/up.tar.gz

