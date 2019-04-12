# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/knative/docs/helloworld
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o helloworld dmesg.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tcp-rdtsc tcp-rdtsc.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tcp-load tcp-load.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM golang

#ADD apt.conf /etc/apt/apt.conf
# Install sysbench
#ADD iozone.test /mnt/mvm-test/iozone.test
RUN apt-get update
RUN apt install --yes sysbench
ADD iozone /usr/bin/iozone
ADD getpid /usr/bin/getpid

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/knative/docs/helloworld/helloworld /helloworld
COPY --from=builder /go/src/github.com/knative/docs/helloworld/tcp-rdtsc /usr/bin/tcp-rdtsc
COPY --from=builder /go/src/github.com/knative/docs/helloworld/tcp-load /usr/bin/tcp-load

# Run the web service on container startup.
CMD ["/helloworld"]
