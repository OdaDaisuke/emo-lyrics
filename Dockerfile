FROM google/cloud-sdk:alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/bin:$PATH

# Install Go
ENV GO_VERSION 1.8.3
RUN curl -Lso go.tar.gz "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz
ENV PATH /usr/local/go/bin:$PATH

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# Install GAE for Go SDK
RUN gcloud components install app-engine-go
RUN chmod 755 /google-cloud-sdk/platform/google_appengine/appcfg.py
ENV PATH /google-cloud-sdk/platform/google_appengine:$PATH

EXPOSE 8080