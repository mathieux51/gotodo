# FROM node:lts-alpine
# RUN npm run build
# COPY __sapper__/build __sapper__/build
# CMD ["node", "__sapper__/build"]
FROM nginx

WORKDIR /go/src/github.com/mathieux51/gotodo/ui
RUN npm run export
COPY __sapper__/export /usr/share/nginx/html