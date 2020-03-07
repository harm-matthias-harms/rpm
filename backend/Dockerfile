############################
# STEP 1 build executable binary
############################
FROM golang:1.14-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/github.com/harm-matthias-harms/rpm/backend

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/backend .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/backend /go/bin/backend

# Use an unprivileged user.
USER appuser

#Expose port of backend
EXPOSE 3001

# Run the hello binary.
ENTRYPOINT ["/go/bin/backend"]