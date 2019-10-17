# iron/go is the alpine image with only ca-certificates added
FROM golang:1.9

WORKDIR /app
# Now just add the binary
ADD users /app/
ENTRYPOINT ["./users"]
EXPOSE 8080