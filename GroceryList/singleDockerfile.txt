# syntax=docker/dockerfile:1

# GO version 1.19 with alpine
FROM golang:1.19-alpine

# Current working directory on alpine
WORKDIR /app

#copy all to /app folder
COPY . ./

#Download modules
RUN go mod download

#Build the image with name docker-grocerylist name
RUN go build -o /docker-grocerylist

#Expose port 5003
EXPOSE 5003

#Run the go image
CMD [ "/docker-grocerylist" ]