# Stage 1: build the Vite app
FROM node:20-alpine AS builder

WORKDIR /app

# Install dependencies (cached layer)
COPY package.json package-lock.json* ./
RUN npm install

# Copy source and build
COPY . .
RUN npm run build

# Stage 2: serve with Nginx
FROM nginx:alpine

# Remove default Nginx static files
RUN rm -rf /usr/share/nginx/html/*

# Copy Vite build output
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy custom Nginx config for SPA routing (history API fallback)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
