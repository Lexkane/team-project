FROM ubuntu:latest

COPY src/team-project /opt/team-project/bin/
COPY src/team_project_config.json /opt/team-project/config/

RUN chmod +x /opt/team-project/bin/*

EXPOSE 8080

WORKDIR /opt

ENTRYPOINT ["/opt/team-project/bin/team-project"]
CMD ["-config", "/opt/team-project/config/team_project_config.json"]
