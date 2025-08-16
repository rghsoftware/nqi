# hadolint global ignore=DL3008,DL3013
FROM mcr.microsoft.com/devcontainers/python:3.11

# Install additional tools
RUN apt-get update && apt-get install -y --no-install-recommends \
  postgresql-client \
  redis-tools \
  httpie \
  jq \
  tree \
  && rm -rf /var/lib/apt/lists/*

# Install uv globally
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN curl -LsSf https://astral.sh/uv/install.sh | sh

# Install global Python tools
RUN pip install --no-cache-dir --upgrade pip \
  && pip install --no-cache-dir\
  ipython \
  rich \
  httpie \
  litecli \
  pgcli

# Create workspace directory
WORKDIR /workspace

# Set up shell
RUN echo 'alias ll="ls -la"' >> ~/.bashrc \
  && echo 'alias gs="git status"' >> ~/.bashrc \
  && echo 'alias gp="git pull"' >> ~/.bashrc \
  && echo 'alias dc="docker-compose"' >> ~/.bashrc
