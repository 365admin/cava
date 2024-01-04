FROM golang:1.20-buster

WORKDIR /app.
COPY ./projects/cli .
RUN go install
EXPOSE 5001
ENTRYPOINT ["tail", "-f", "/dev/null"]
# CMD [ "cava" ,"serve"]