FROM    alpine:3.19

LABEL   description="Baremetal CSI Operator"

ADD     controller  node-controller

RUN apk update; apk upgrade

RUN addgroup -S bmcsi && adduser -S bmcsi -u 1000 -G bmcsi

USER 1000

ENTRYPOINT ["/node-controller"]
