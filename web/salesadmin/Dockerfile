FROM node:10.16.3-alpine

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY ./web/salesadmin/package.json /app/package.json
#
## need to run this command before any npm install
RUN npm config set unsafe-perm true
RUN npm install --silent
RUN npm install react-scripts -g --silent
COPY ./web/salesadmin /app
RUN npm run build
RUN npm install -g serve
CMD ["serve", "-s", "build", "-l", "4000"]
