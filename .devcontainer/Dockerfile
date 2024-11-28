FROM golang:1.23-bullseye

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    curl -fsSL https://deb.nodesource.com/setup_16.x | bash - && \
    apt update && \
    apt install -y sudo git make curl gnupg iputils-ping dnsutils build-essential nodejs="16.*" zip && \
    echo 'alias ll="ls -lsah"' >> ~/.bashrc && \
    go install golang.org/x/tools/gopls@latest && \
    # Base Go tools needed for VS code Go extension
    go install golang.org/x/tools/cmd/guru@latest && \
    go install golang.org/x/tools/cmd/gorename@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    # Extra tools integrating with VS code
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/cweill/gotests/...@latest &&\
    go install github.com/davidrjenni/reftools/cmd/fillstruct@latest && \
    # Terminal tools
    go install github.com/golang/mock/mockgen@latest && \
    go install github.com/vektra/mockery/v2/...@latest && \
    2>&1 && \
    rm -rf $GOPATH/pkg/* $GOPATH/src/* /root/.cache/go-build && \
    export PATH="$PATH:$(go env GOPATH)/bin"


