FROM golang:1.14 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w"
RUN ls -lh

FROM mcr.microsoft.com/dotnet/core/sdk:5.0
ENV RESHARPER_CLI_VERSION=2020.3.2

RUN mkdir -p /usr/local/share/dotnet/sdk/NuGetFallbackFolder

WORKDIR /resharper
RUN \
  curl -o resharper.tgz -L "https://download.jetbrains.com/resharper/ReSharperUltimate.$RESHARPER_CLI_VERSION/JetBrains.ReSharper.CommandLineTools.Unix.$RESHARPER_CLI_VERSION.tar.gz" \
  && tar -xvf resharper.tgz \
  && rm resharper.tgz \
  && rm -rf macos-x64
ENV PATH="/resharper:${PATH}"

# this is the same as the base image
WORKDIR /

COPY --from=builder /build/resharper-action /usr/bin
CMD resharper-action
