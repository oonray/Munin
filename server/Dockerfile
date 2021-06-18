FROM openjdk:8-jre-alpine

RUN mkdir /MC

RUN apk update
RUN apk add git

RUN git clone https://github.com/oonray/Munin.git /MC/Munin

WORKDIR /MC/Munin

RUN echo "eula=true" >> eula.txt

RUN ls -lah
RUN java -jar bin/modpack.jar --installServer

CMD ["java", "-Xmx1024M", "-Xms1024M", "-jar", "minecraft_server.1.16.5.jar", "nogui"]


