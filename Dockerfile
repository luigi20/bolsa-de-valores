FROM node:24-bookworm

WORKDIR /app

COPY package*.json ./

RUN npm install --legacy-peer-deps

COPY . .

RUN npx prisma generate

EXPOSE 3333

CMD ["npm", "run", "start:dev"]