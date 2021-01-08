# Filename: Dockerfile 
FROM ubuntu
WORKDIR /usr/src/app
COPY . .
EXPOSE 9090
CMD ["./cudui"]