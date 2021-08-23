## BUILD
FROM node:12-alpine AS builder

WORKDIR /app

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY . .

RUN yarn generate

# PRODUCTION
FROM nginx:alpine

COPY default.conf /etc/nginx/conf.d/

COPY --from=builder /app/dist /usr/share/nginx/html/