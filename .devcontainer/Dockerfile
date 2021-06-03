# Dockerfile for InvoiceGen Dev Evn.
# Ubuntu - v20.04
# GoLang - v1.16.4
# node - v14.x
# npm - v6.x

#FROM ubuntu:20.04
FROM ubuntu:20.04

# install any updates (standard routine for a fresh system)
RUN apt-get update
RUN apt-get upgrade -y

RUN apt-get -y install build-essential

# Install wget to download files
RUN apt-get -y install wget

# Create temp directory
WORKDIR /temp_downloads

# Install GoLang
ARG GOFILE=go1.16.4.linux-amd64.tar.gz
RUN wget https://golang.org/dl/$GOFILE
RUN tar -C /usr/local -xzf $GOFILE

ENV PATH="/usr/local/go/bin:${PATH}"

ENV GOPATH="/go"

RUN go get -v golang.org/x/tools/gopls
RUN go get -v honnef.co/go/tools/...
RUN go get -v golang.org/x/lint/golint
RUN go get -v github.com/mgechev/revive
RUN go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get -v github.com/ramya-rao-a/go-outline
RUN go get -v github.com/go-delve/delve/cmd/dlv
RUN go get -v github.com/golangci/golangci-lint/cmd/golangci-lint

# Confirm go version
#RUN go version

# Install Nodejs and npm (npm is bundled with nodejs)
RUN wget https://deb.nodesource.com/setup_14.x

RUN bash setup_14.x
RUN apt-get install -y nodejs

# Confirm node and npm versions
RUN node -v
RUN npm -v

# Clean up
WORKDIR /
RUN rm -rf /temp_downloads

# Final confirmation message
RUN echo "Successfully setup Go ENV"