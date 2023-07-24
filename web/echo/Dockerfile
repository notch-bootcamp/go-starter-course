# Step 1
FROM golang:1.20 AS build
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /echo-app
# Step 1.5: intermediate step with UPX compress would go here for optimization
# Step 2
# use small distroless to run the code faster
FROM gcr.io/distroless/static-debian11
COPY --from=build echo-app /
ENTRYPOINT ["/echo-app"]
