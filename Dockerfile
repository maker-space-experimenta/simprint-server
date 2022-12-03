
## Build webapp
FROM alpine AS build_webapp

WORKDIR /app
RUN cd /app
RUN apk add --update npm
RUN npm --version
RUN node --version
COPY ./webapp ./
RUN npm install
RUN npm run build


## Build golang
FROM golang:1.19.3 AS build_go

WORKDIR /app
RUN cd /app
COPY ./ ./
RUN go mod download
RUN go build -o /printer_service ./cmd/makerspace_printer_kiosk/main.go
RUN ls -la /


## DEPLOYMENT
FROM golang:1.19.3

WORKDIR /app
COPY --from=build_go /printer_service /app/printer_service
COPY --from=build_webapp /app/dist /app/static
COPY ./images /app/images
COPY ./slicer-configs /app/slicer-configs
RUN ls -la /app

EXPOSE 5000
CMD [ "/app/printer_service" ]