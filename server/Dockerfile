FROM ubuntu

RUN mkdir /MC

RUN apt-get -y update
RUN apt-get -y install git

RUN git clone https://github.com/oonray/Munin.git /MC/Munin

RUN ls -lah
FROM openjdk
RUN ls -lah
WORKDIR /MC/Munin

RUN ls -lah
RUN java -jar bin/modpack.jar --installServer


