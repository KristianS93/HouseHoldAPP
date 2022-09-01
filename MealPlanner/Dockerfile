# syntax=docker/dockerfile:1

# GO version 1.19 with alpine
FROM golang:1.19-alpine

# Current working directory on alpine
WORKDIR /app

#copy all to /app folder
COPY . ./
COPY ./data/ ./data/
COPY ./database/ ./database/

#Download modules
RUN go mod download

#Build the image with name docker-grocerylist name
RUN go build -o /docker-mealplanner

#Expose port 5003
EXPOSE 5005

#Run the go image
CMD [ "/docker-mealplanner" ]


# # GO version 1.19 with alpine
# FROM golang:1.19-buster AS build

# # Current working directory on alpine
# WORKDIR /app

# #copy all to /app folder
# COPY . ./

# #Download modules
# RUN go mod download

# #Build the image with name docker-grocerylist name
# RUN go build -o /docker-mealplanner

# #DEPLOY
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /docker-mealplanner /docker-mealplanner
# COPY ./database/.env ./database/
# # COPY ./data/ ./data/

# #Expose port 5005
# EXPOSE 5005

# USER nonroot:nonroot

# #Run the go image
# ENTRYPOINT [ "/docker-mealplanner" ]






#BUILD
# # GO version 1.19 with alpine
# FROM golang:1.19-buster AS build

# # Current working directory on alpine
# WORKDIR /app

# #copy all to /app folder
# COPY . ./

# #Download modules
# RUN go mod download

# #Build the image with name docker-grocerylist name
# RUN go build -o /docker-grocerylist


# #DEPLOY
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /docker-grocerylist /docker-grocerylist

# #Expose port 5003
# EXPOSE 5003

# USER nonroot:nonroot

# #Run the go image
# ENTRYPOINT [ "/docker-grocerylist" ]

################################
# syntax=docker/dockerfile:1

# # GO version 1.19 with alpine
# FROM golang:1.19-alpine

# # Current working directory on alpine
# WORKDIR /app

# #copy all to /app folder
# COPY . ./

# #Download modules
# RUN go mod download

# #Build the image with name docker-grocerylist name
# RUN go build -o /docker-mealplanner

# #Expose port 5003
# EXPOSE 5003

# #Run the go image
# CMD [ "/docker-mealplanner" ]

