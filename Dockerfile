FROM gcr.io/distroless/static:latest
LABEL description = "IBM Cloud Object Storage COSI driver"

COPY ./bin/ibm-cosi-driver ibm-cosi-driver
ENTRYPOINT ["/ibm-cosi-driver"]
