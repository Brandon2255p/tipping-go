FROM golang as builder

ADD . /go/src/Brandon2255p/tipping-go
WORKDIR $GOPATH/src/Brandon2255p/tipping-go

# RUN go get -d -v
RUN go get github.com/derekparker/delve/cmd/dlv
# RUN go get github.com/looplab/eventhorizon
# RUN go build -o main

EXPOSE 8080
ENTRYPOINT ["./main"]


# # STEP 2 build a small image
# # start from scratch
# FROM scratch
# # Copy our static executable
# COPY --from=builder /go/bin/main /go/bin/main
# ENTRYPOINT ["/go/bin/main"]