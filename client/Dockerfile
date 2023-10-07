# Stage 1: Build an Angular Docker Image
# Build Stage
FROM node:alpine3.18 as BuildStage
LABEL stage=build

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the working directory
COPY package*.json /app/

# Install dependencies
RUN npm install

# Copy the entire application source to the working directory
COPY . /app

# Define an argument for the build configuration (e.g., production)
ARG configuration=production

# Build the Angular application with the specifeid configuration
RUN npm run build -- --output-path=./dist/out --configuration $configuration

# Stage 2: Serve Angular App using Nginx
# Deploy Stage
FROM nginx:alpine3.18-slim AS DeployStage
LABEL stage=deploy

# Create a non-root group
# RUN addgroup -S nonroot

# Create a non-root user and add them to the nonroot group
# RUN adduser -S -G nonroot nonroot

# Start the application as the nonroot user
# USER nonroot:nonroot

# Copy the compiled Angular application to the default Nginx public folder
COPY --from=BuildStage /app/dist/out/ /usr/share/nginx/html

# Copy the custom Nginx configuration file to the default Nginx configuration folder
COPY /nginx-custom.conf /etc/nginx/conf.d/default.conf

# Expose the port that your Angular app will listen on
EXPOSE 80

# The official Nginx image is configured to start Nginx automatically when
# a container is launched, and it handles serving static files which includes
# the compiled Angular application.  Therefore, we can rely on the default
# behavior of the Nginx image to serve the Angular applictation without
# specifying an ENTRYPOINT or CMD instruction.
