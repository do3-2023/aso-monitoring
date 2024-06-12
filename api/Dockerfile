# Run step
FROM node:20-alpine
WORKDIR /var/app

COPY package*.json ./
RUN npm ci

COPY . /var/app
RUN npm run build
RUN rm .env || true
RUN npm prune --omit=dev

CMD [ "npm", "run", "start:prod" ]

