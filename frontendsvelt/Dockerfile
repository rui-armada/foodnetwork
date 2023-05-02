FROM node:16-alpine

WORKDIR /app

# copy files and install dependencies
COPY . ./
RUN npm install
RUN npm run build

EXPOSE 5000

CMD ["npm", "start"]