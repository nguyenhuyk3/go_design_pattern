FROM golang:1.23.1-alpine

# * Set up work dir in container
WORKDIR /app

# * Copy all content from the current dir on my Lap into 
# * work dir in container (/app)
# * First dot is source 
# * Second dot is destination
COPY . .

# * Compile all code in current dir (.) and
# * create binary file named myApp
RUN go build -o my-design-pattern-app .

# * Run app wwhen container start 
CMD ["./my-design-pattern-app"]

