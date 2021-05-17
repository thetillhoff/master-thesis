FROM gitpod/workspace-base:latest

RUN sudo sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -b /usr/local/bin
