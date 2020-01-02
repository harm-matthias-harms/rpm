## BUILD
FROM node:10-alpine AS builder

WORKDIR /app

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY . .

RUN yarn run generate

## PRODUCTION
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html/