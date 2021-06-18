FROM mcr.microsoft.com/java/jdk

RUN mkdir /MC
WORKDIR /MC
RUN apt-get install git
RUN git clone git@github.com:oonray/Munin.git
WORKDIR Munin
RUN java -jar bin/modpack.jar --installServer


