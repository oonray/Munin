FROM openjdk:8-jre-alpine

RUN mkdir /MC

RUN apk update
RUN apk add git

RUN git clone https://github.com/oonray/Munin.git /MC/Munin

WORKDIR /MC/Munin

RUN ls -lah
RUN java -jar bin/modpack.jar --installServer


