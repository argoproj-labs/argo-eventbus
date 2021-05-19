####################################################################################################
# base
####################################################################################################
FROM alpine:3.12.3 as base
RUN apk update && apk upgrade && \
    apk add ca-certificates && \
    apk --no-cache add tzdata

####################################################################################################
# argo-eventbus
####################################################################################################
FROM scratch as argo-eventbus
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY dist/argo-eventbus /bin/argo-eventbus
ENTRYPOINT [ "/bin/argo-eventbus" ]

