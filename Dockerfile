FROM gcr.io/distroless/static:nonroot
# Copy the binary from the build stage
COPY gofrommediumtohugo /gofrommediumtohugo
# Run the binary
ENTRYPOINT ["/gofrommediumtohugo"]