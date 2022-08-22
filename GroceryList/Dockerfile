# syntax=docker/dockerfile:1

#BUILD
# GO version 1.19 with alpine
FROM golang:1.19-buster AS build

# Current working directory on alpine
WORKDIR /app

#copy all to /app folder
COPY . ./

#Download modules
RUN go mod download

#Build the image with name docker-grocerylist name
RUN go build -o /docker-grocerylist


#DEPLOY
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-grocerylist /docker-grocerylist

#Expose port 5003
EXPOSE 5003

USER nonroot:nonroot

#Run the go image
ENTRYPOINT [ "/docker-grocerylist" ]
