FROM node:8-alpine
# ENV PORT 3000
# EXPOSE 3000

WORKDIR /app
RUN npm install -g nodemon

COPY package*.json ./
RUN npm install
CMD ["nodemon", "/app/src/index.js"]