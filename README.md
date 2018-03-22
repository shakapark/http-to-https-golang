# http-to-https-golang

 - How to build binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

 - How to build image
 docker build -t example-scratch -f Dockerfile .
 
 - Run the container
 docker run -p 80:8080 -d example-scratch
