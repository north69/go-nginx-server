# Official Nginx image (Yes, in this article I always use the `latest`. Kill me!)
FROM nginx:alpine

# Copy Nginx configuration file to the container
COPY ./docker/nginx/nginx.conf /etc/nginx/conf.d

# Expose port 80
EXPOSE 80