FROM node:lts-alpine

RUN npm install -g http-server

RUN mkdir /frontend

COPY frontend/awesome-scrambler/package.json /frontend
COPY frontend/awesome-scrambler/package-lock.json /frontend

WORKDIR /frontend

RUN npm install

COPY frontend/awesome-scrambler/ /frontend

RUN npm run build

EXPOSE 80

CMD [ "http-server", "dist" ]
