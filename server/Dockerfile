FROM ubuntu

RUN mkdir /MC
WORKDIR /MC

RUN apt-get -y update
RUN apt-get -y install git

RUN git clone https://github.com/oonray/Munin.git

RUN ls -lah
FROM openjdk
RUN ls -lah
WORKDIR Munin

RUN ls -lah
RUN java -jar bin/modpack.jar --installServer


