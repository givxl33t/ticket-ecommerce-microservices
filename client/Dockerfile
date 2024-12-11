# Common build stage
FROM node:20-alpine as common-build-stage

WORKDIR /app
COPY package.json .
COPY . .

# Development build stage
FROM common-build-stage as development-build-stage

RUN npm install

CMD ["npm", "run", "dev"]

# Production build stage
FROM common-build-stage as production-build-stage

RUN npm install --omit=dev

RUN npm run build

CMD ["npm", "run", "start"]