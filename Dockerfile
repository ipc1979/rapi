FROM golang:1.18

# Set destination for COPY
WORKDIR /vemo

# Download Go modules
COPY go.mod ./
#RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY controllers/ ./controllers/
COPY model/ ./model/
COPY repositories/ ./repositories/
COPY restapi/ ./restapi/
COPY services/ ./services/
COPY vemoapp.go ./

# Build
RUN go build -o /vemoapp

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
#
CMD ["/vemoapp"]