FROM golang:1.15
ARG GITHUB_TOKEN
RUN go get -u github.com/cloudspannerecosystem/wrench
RUN go get -u go.mercari.io/yo
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
WORKDIR /go/src/github.com/knocknote/spanner-test
RUN git clone https://github.com/knocknote/go-tool-yo-template
