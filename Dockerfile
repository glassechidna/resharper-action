# DOT_NET_VERSION need put here to be able to use it in FROM
# ARG DOT_NET_VERSION

FROM golang:1.14 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w"
RUN ls -lh


# put the resharper binary in a scratch container
FROM mcr.microsoft.com/dotnet/sdk:6.0-jammy-amd64
ARG RESHARPER_CLI_VERSION

RUN echo 'export PATH="$PATH:~/.dotnet/tools"' >> .bashrc
ENV PATH $PATH:~/.dotnet/tools

# install jb cli(include inspectcode)
RUN dotnet --info
RUN dotnet tool install JetBrains.ReSharper.GlobalTools --global --version $RESHARPER_CLI_VERSION


# RUN mkdir -p /usr/local/share/dotnet/sdk/NuGetFallbackFolder

COPY --from=builder /build/resharper-action /usr/bin
CMD resharper-action
# CMD dotnet exec

# # =======================================================
# # DOT_NET_VERSION need put here to be able to use it in FROM
# ARG DOT_NET_VERSION
#
# FROM golang:1.14 AS builder
#
# WORKDIR /build
# COPY . .
# RUN CGO_ENABLED=0 go build -ldflags="-s -w"
# RUN ls -lh
#
#
# # put the resharper binary in a scratch container
# FROM ubuntu:22.04
# ARG RESHARPER_CLI_VERSION
#
# RUN apt-get update && apt-get install -y \
#     wget \
#     curl \
#     unzip \
#     && rm -rf /var/lib/apt/lists/*
#
# RUN apt remove -y dotnet*  aspnetcore* netstandard*; true
#
# # RUN echo 'Package: *' >> /etc/apt/preferences.d/dotnet
# # RUN echo 'Pin: origin "packages.microsoft.com"' >> /etc/apt/preferences.d/dotnet
# # RUN echo 'Pin-Priority: 1001' >> /etc/apt/preferences.d/dotnet
#
# RUN wget https://packages.microsoft.com/config/ubuntu/22.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb \
#   && dpkg -i packages-microsoft-prod.deb \
#   && rm -f packages-microsoft-prod.deb \
#   && apt-get update \
#   && apt-get install -y dotnet6 dotnet-runtime-6.0 \
#   && rm -rf /var/lib/apt/lists/*
#
#
# # RUN mkdir -p /usr/local/share/dotnet/sdk/NuGetFallbackFolder && mkdir -p /usr/share/dotnet/host/fxr
# #
# # # install jb cli(include inspectcode)
# # RUN dotnet tool install JetBrains.ReSharper.GlobalTools --global --version $RESHARPER_CLI_VERSION
# #
# # ENV PATH $PATH:/root/.dotnet/tools
# #
# #
# # COPY --from=builder /build/resharper-action /usr/bin
# # CMD resharper-action
# # # CMD dotnet exec
# # =======================================================

# FROM docker pull mcr.microsoft.com/dotnet/sdk:7.0-jammy
# FROM docker pull mcr.microsoft.com/dotnet/sdk:6.0-jammy-amd64

# install jb cli(include inspectcode)
# RUN dotnet tool install JetBrains.ReSharper.GlobalTools --global --version $RESHARPER_CLI_VERSION
