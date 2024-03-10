FROM node:18-alpine AS frontend
WORKDIR /app
RUN chmod +x ./
COPY frontend/package*.json frontend/ ./
RUN npm install --production
COPY frontend .

FROM nginx:alpine
COPY --from=frontend ./public /usr/share/nginx/html

# Modify Nginx configuration (within Dockerfile or separate file)
RUN echo "location / {" >> /etc/nginx/conf.d/default.conf
RUN echo "  root /usr/share/nginx/html/frontend/build;  # Adjust path based on your build structure" >> /etc/nginx/conf.d/default.conf
RUN echo "}" >> /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
