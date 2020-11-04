FROM node:latest
WORKDIR /usr/client/
COPY yarn.lock .
COPY package.json .
RUN yarn install
COPY . .
