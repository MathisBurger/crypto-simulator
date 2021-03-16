FROM node:14-alpine AS webBuild

WORKDIR /webapp

COPY ./web .

RUN npm ci
RUN npm run build --prod

FROM golang:1.16-alpine AS backendBuild

WORKDIR /backend

COPY . .

RUN go mod tidy

RUN go build

FROM alpine:3 AS final

WORKDIR /app

COPY --from=webBuild /webapp/dist .
COPY --from=backendBuild /backend/crypto-simulator .
COPY --from=backendBuild /backend/.env .

RUN chmod +x ./crypto-simulator

EXPOSE 8080

CMD ["./crypto-simulator"]