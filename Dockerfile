FROM gcr.io/distroless/static:latest

COPY ./bin/ibm-cosi-driver ibm-cosi-driver
ENTRYPOINT ["/ibm-cosi-driver"]
