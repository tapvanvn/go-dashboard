FROM tapvanvn/rpi64_dashboard_base AS build

WORKDIR /

RUN apk update && apk add --no-cache git curl openssh-client gcc g++ musl-dev

RUN mkdir -p /src

COPY ./ /src/

RUN cd /src && go get ./... && go build

FROM arm64v8/alpine
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
        
COPY --from=build               /src/go-dashboard / 
COPY static/                    /static

ENV PORT=80

ENTRYPOINT ["/go-dashboard"]