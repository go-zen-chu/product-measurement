FROM grafana/grafana:latest

USER root

RUN grafana-cli plugins install yesoreyeram-infinity-datasource

USER grafana

ENV GF_INSTALL_PLUGINS="yesoreyeram-infinity-datasource"
ENV GF_PATHS_PROVISIONING="/etc/grafana/provisioning"
ENV GF_SECURITY_ALLOW_EMBEDDING=true
ENV GF_AUTH_ANONYMOUS_ENABLED=true

RUN mkdir -p "$GF_PATHS_PROVISIONING/datasources" \
    && mkdir -p "$GF_PATHS_PROVISIONING/dashboards"

COPY ./provisioning/datasources /etc/grafana/provisioning/datasources/
COPY ./provisioning/dashboards /etc/grafana/provisioning/dashboards/

EXPOSE 3000
