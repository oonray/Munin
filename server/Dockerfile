FROM openjdk:8-jre-alpine

RUN mkdir /MC

RUN apk -y update
RUN apk -y install git

RUN git clone https://github.com/oonray/Munin.git /MC/Munin

WORKDIR /MC/Munin

RUN ls -lah
RUN java -jar bin/modpack.jar --installServer


