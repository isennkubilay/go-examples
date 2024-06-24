FROM ubuntu:22.04
ARG DEBIAN_FRONTEND=noninteractive
WORKDIR /app
RUN apt-get update -y && \
    apt-get install --no-install-recommends \
    -y vim \
        nano \
        net-tools \
        curl \
        tcpdump \
        iftop \
        netcat \
        dnsutils \
        strace \
        htop \
        iputils-ping \
        nano \
        traceroute \
        nmap \ 
        iperf3 \
        python3 \ 
        python3-pip \
        htop \
        wget \
        tar \
        tshark \
        vnstat \ 
        bmon \
        network-manager \
        mtr \
        tzdata 

RUN wget https://fastdl.mongodb.org/tools/db/mongodb-database-tools-ubuntu2004-x86_64-100.5.0.tgz && \
    tar -zxvf mongodb-database-tools-ubuntu2004-x86_64-100.5.0.tgz && \
    mv mongodb-database-tools-ubuntu2004-x86_64-100.5.0/bin/* /usr/bin/ && \
    rm -rf mongodb-database-tools-ubuntu2004-x86_64-100.5.0.tgz mongodb-database-tools-ubuntu2004-x86_64-100.5.0