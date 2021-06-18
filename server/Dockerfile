FROM openjdk

RUN mkdir /MC
WORKDIR /MC
RUN git clone git@github.com:oonray/Munin.git
WORKDIR Munin
RUN java -jar bin/modpack.jar --installServer


