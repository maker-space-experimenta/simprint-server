
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


## Build prusa-slicer
FROM keyglitch/docker-slic3r-prusa3d AS prusa


## DEPLOYMENT
FROM golang:1.19.3

WORKDIR /app
COPY --from=prusa /Slic3r/slic3r-dist /prusa 
COPY --from=build_go /printer_service /app/printer_service
COPY --from=build_webapp /app/dist /app/static
COPY ./images /app/images
COPY ./slicer-configs /app/slicer-configs
RUN ls -la /app
RUN apt update
RUN apt install -y \
    libsm6 \
    libxext6 \
    ffmpeg \
    libfontconfig1 \
    libxrender1 \
    libgl1-mesa-glx \
    libasound2 \
    libatk1.0-0 \
    libc6 \
    libcairo2 \
    libcups2 \
    libdbus-1-3 \
    libexpat1\ 
    libfontconfig1\ 
    libgcc1\ 
    libgconf-2-4\ 
    libgdk-pixbuf2.0-0\ 
    libglib2.0-0\ 
    libgtk-3-0\ 
    libnspr4\ 
    libpango-1.0-0\ 
    libpangocairo-1.0-0\ 
    libstdc++6\ 
    libx11-6\ 
    libx11-xcb1\ 
    libxcb1\ 
    libxcursor1\ 
    libxdamage1\ 
    libxext6\ 
    libxfixes3\ 
    libxi6\ 
    libxrandr2\ 
    libxrender1\ 
    libxss1\ 
    libxtst6\ 
    libnss3\ 
    libpangoxft-1.0-0\ 
    libgtk2.0-0\ 
    libglu1-mesa



EXPOSE 5000
CMD [ "/app/printer_service" ]