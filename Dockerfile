# step 1. compile
FROM "sky0621dhub/dockerfile-gowithdep" AS builder

COPY . /go/src/github.com/sky0621/go-webapi-for-gke-study
WORKDIR /go/src/github.com/sky0621/go-webapi-for-gke-study
RUN dep ensure
#RUN go test ./...
RUN CGO_ENABLED=0 go build -o go-webapi-for-gke-study github.com/sky0621/go-webapi-for-gke-study

# -----------------------------------------------------------------------------
# step 2. build
FROM scratch
COPY --from=builder /go/src/github.com/sky0621/go-webapi-for-gke-study/ .
EXPOSE 80
ENTRYPOINT [ "./go-webapi-for-gke-study" ]
