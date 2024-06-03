# Use a Node.js base image for building
FROM node:20-alpine AS builder

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the .env file
# COPY .env .

# Copy and build the app
COPY . .
RUN npm run build

# Production environment
FROM nginx:alpine

# Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built Svelte app to Nginx directory
COPY --from=builder /app/dist/client /usr/share/nginx/html

# Expose port 80
EXPOSE 80
