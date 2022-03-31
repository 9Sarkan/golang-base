#######################
#     build image     #
#######################
FROM golang:1.18 AS Builder

# set working dir
WORKDIR $GOPATH/src/base/

# Create appuser
ENV USER=baseuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

COPY . .

# install dependencies
RUN go mod download
RUN go mod verify

# build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/base

#######################
# build a small image #
#######################
FROM scratch

# Import from builder stage
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/base /application/base

USER baseuser:baseuser

CMD [ "/application/base" ]