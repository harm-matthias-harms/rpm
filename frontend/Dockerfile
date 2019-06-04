## BUILD
FROM node:alpine AS builder

WORKDIR /app

COPY package.json .
RUN npm ci

COPY . .

RUN npm run generate

## PRODUCTION
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html/