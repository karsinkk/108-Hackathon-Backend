FROM golang

# Create the directory where the application will reside
RUN mkdir /108

# Copy the application files (needed for production)
ADD . /108/

# Set the working directory to the app directory
WORKDIR /108

EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT /108/main