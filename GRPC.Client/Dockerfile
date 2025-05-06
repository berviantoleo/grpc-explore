FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine AS build
RUN apk add grpc-plugins
ENV PROTOBUF_PROTOC=/usr/bin/protoc
ENV GRPC_PROTOC_PLUGIN=/usr/bin/grpc_csharp_plugin
WORKDIR /app
COPY . .
RUN dotnet restore
RUN dotnet publish -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:8.0-alpine AS runtime
COPY --from=build /app/publish /app/publish
WORKDIR /app/publish
CMD ["dotnet", "GRPC.Client.dll"]
