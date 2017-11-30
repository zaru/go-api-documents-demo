FROM golang:1.9.2-stretch

RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    lsof \
    vim \
    nodejs \
    mysql-client

# bashrc
RUN echo 'PS1="${debian_chroot:+($debian_chroot)}\e[35m\u@\h:\w\e[0m\\n$ "' >> ~/.bashrc; \
    echo "alias ll='ls -la'" >> ~/.bashrc

WORKDIR /go/src/github.com/zaru/go-api-documents-demo

# Golang package
RUN go get -u \
    github.com/golang/dep/cmd/dep \
    github.com/tockins/realize \
    github.com/pressly/goose/cmd/goose \
    github.com/golang/lint/golint

# node.js package
RUN npm install --unsafe-perm -g yarn aglio drakov dredd

COPY . .

EXPOSE 1323
