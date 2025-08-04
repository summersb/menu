FROM golang:tip-bookworm

# Create non-root user and group
RUN useradd -m -u 1000 devuser

WORKDIR /app

USER devuser
# Only copy go.mod and go.sum first
COPY --chown=devuser:devuser go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code (for production only, dev will mount)
COPY --chown=devuser:devuser . .

# Install a live reload tool (optional)
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air"]
