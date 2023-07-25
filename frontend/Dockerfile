# build
FROM node:16 AS builder
# Create app directory
WORKDIR /usr/src/app
ARG SHOP_URL=${SHOP_URL}
ARG VUE_APP_API_URL=https://${SHOP_URL}/api
ARG NODE_ENV=dev
COPY ./package.json .
RUN npm install
COPY . .
RUN npm run build
RUN ls -la /usr/src/app/dist

# release
FROM nginx:1.22.1-alpine
RUN rm -rf /usr/share/nginx/html/*
# copy files from builder
COPY  --from=builder /usr/src/app/dist/ /usr/share/nginx/html/
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]