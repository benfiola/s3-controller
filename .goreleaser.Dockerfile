FROM debian:bookworm-slim AS homelab_helper_builder
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/s3-controller /usr/bin/
ENTRYPOINT ["s3-controller"]
