FROM golang
#MOVE TO AND MAKE BUILD DIRECTORY
WORKDIR /build
# Copy the code into the container
COPY . .
# run go build to compile the binary
RUN go build -o main .
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/main"]