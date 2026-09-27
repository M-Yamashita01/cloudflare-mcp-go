# Build stage
FROM golang:1.27-alpine AS build

WORKDIR /src

# Cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Build a static binary (no cgo) so it runs on a minimal base image
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/cloudflare-mcp-go .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

# The server needs a Cloudflare API token at request time only. It starts and
# lists its tools without a token, so MCP introspection works out of the box.
COPY --from=build /out/cloudflare-mcp-go /usr/local/bin/cloudflare-mcp-go

ENTRYPOINT ["/usr/local/bin/cloudflare-mcp-go"]
