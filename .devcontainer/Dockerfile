FROM golang:1.26-bookworm

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git curl wget ca-certificates gnupg unzip xz-utils \
    build-essential jq python3 python3-pip ripgrep \
    postgresql-client docker docker-compose sudo \
    && rm -rf /var/lib/apt/lists/*

# Install Node.js LTS
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# Set up non-root user FIRST
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# Set Go env for dev user before installing tools
ENV GOPATH=/home/$USERNAME/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV HOME=/home/$USERNAME

# Install Go tools as dev user so they land in the right GOPATH
USER $USERNAME

RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/a-h/templ/cmd/templ@latest

# Install Claude CLI (needs sudo for global npm install)
RUN sudo npm install -g @anthropic-ai/claude-code

# Pre-create the mod cache dir with correct ownership
RUN mkdir -p $GOPATH/pkg/mod $GOPATH/bin /home/dev/.claude
WORKDIR /workspace
