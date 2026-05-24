FROM golang:1.26-bookworm

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    curl \
    wget \
    ca-certificates \
    gnupg \
    unzip \
    xz-utils \
    build-essential \
    jq \
    && rm -rf /var/lib/apt/lists/*

# Install Node.js (LTS) for frontend tooling / npm packages
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install Claude CLI
RUN npm install -g @anthropic-ai/claude-code

# Install templ (Go HTML templating, pairs well with HTMX)
RUN go install github.com/a-h/templ/cmd/templ@latest

# Set up a non-root user
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update && apt-get install -y sudo \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && rm -rf /var/lib/apt/lists/*

# Make Go tools available to the dev user
RUN mkdir -p /root/go && mkdir -p /home/$USERNAME/go && cp -r /root/go /home/$USERNAME/go \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME/go

USER $USERNAME

ENV GOPATH=/home/dev/go
ENV PATH=$GOPATH/bin:$PATH
ENV HOME=/home/dev

WORKDIR /workspace
