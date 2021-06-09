FROM ubuntu

COPY bedrock-server-1.16.221.01.zip /
RUN apt-get update && apt install -y unzip openssl
RUN unzip bedrock-server-1.16.221.01.zip
RUN chmod +x /bedrock_server

ENTRYPOINT /bedrock_server
EXPOSE 19132