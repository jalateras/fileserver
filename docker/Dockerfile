FROM node:0.10.36

MAINTAINER Jim Alateras <jima@comware.com.au>

COPY . /app

WORKDIR /app

RUN \
  npm install --production && \
  chmod +x bin/docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["bin/docker-entrypoint.sh"]
CMD ["node", "index.js"]


