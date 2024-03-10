FROM python:3.6.5

ENV DEBIAN_FRONTEND noninteractive
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p /workspace
ADD __init__.py typed.py sender.py info.py /workspace/

ENTRYPOINT ["/workspace/sender.py"]
