# Stage 1: Build Next.js Frontend (adjust comment if using Vite)
FROM node:18-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install --production
COPY frontend .
RUN npm run build  # Use npm run build for production build

# Stage 2: Build Golang Backend
FROM golang:1.19-alpine AS backend
WORKDIR /go/src/block-cipher-webapp/backend
COPY backend/ ./
RUN go mod download
RUN go build -o backend .

# Stage 3: Combine frontend and backend (final image)
FROM nginx:alpine
COPY --from=frontend /app /usr/share/nginx/html
COPY --from=backend /go/src/block-cipher-webapp/backend /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
