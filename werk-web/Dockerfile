FROM node:16

WORKDIR /usr/app

COPY . .

RUN npm ci --only=production
RUN npm run build

CMD [ "npm", "start" ]
