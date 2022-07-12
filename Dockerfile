# FROM golang:1.17.5-alpine AS builder
# LABEL stage=builder

# RUN apk add --no-cache upx 
# ENV GOPATH /go
# COPY  src/ /go/src/
# WORKDIR /go/src/

# RUN echo $GOPATH
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build . 
# RUN upx awslogcheck


FROM scratch AS final
LABEL maintainer="Sylvain Gaunet <sgaunet@gmail.com>"
# COPY --from=alpine /etc/ssl/certs/ /etc/ssl/certs/
COPY "resources" /
WORKDIR /
COPY gopsql /gopsql
USER MyUser
# CMD ["/gopsql", "-c", "/cfg.yaml"]