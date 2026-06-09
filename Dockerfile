FROM golang:1.22

COPY . .

ENV CACHE_TIME 300
ENV API_KEY test
ENV CHANNEL_ID test

CMD [ "go", "run", "." ]