FROM node:18-alpine

WORKDIR /rudderstack-app

COPY package.json .
COPY package-lock.json .

RUN npm install 

COPY . .

RUN ls

CMD ["npm", "start"]