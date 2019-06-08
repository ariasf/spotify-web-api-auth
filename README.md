# spotify-web-api-auth
Using AWS lambda and Golang implement the endpoints for Spotify Athorization Code Flow  for mobile apps. The main idea is to make it easier to fork this code and deploy it for your integration with spotify authorization code workflow, normally used for mobile apps.

## References:

[Spotify Documentation on Swap and Refresh](https://developer.spotify.com/documentation/ios/guides/token-swap-and-refresh/)

[Using Go Gin with AWS Lambda](https://github.com/awslabs/aws-lambda-go-api-proxy)


## endpoints 

**Login:**{{API_URL}}/spotify/login

**Swap Token:** {{API_URL}}/spotify/swap

**Refresh Token:**{{API_URL}}/spotify/refreshtoken


## Start 
1) cd into /src/main 
2) >go get -u github.com/golang/dep/cmd/dep
3) > dep init 
4) > dep ensure 
5) We use godep for package management see [dep](https://golang.github.io/dep/docs/installation.html)
6) > go build main.go
7) Make sure you do **step 6** before zipping the project, otherwise  the zip file won't have the latest build

   
### zipping the exec file aws lambdas  
 1) make sure to set **GOHOSTARCH="amd64"** and **GOOS="linux"** env vars 
 2) cd to `src/main`  and type `chmod +x main`  and    `zip main.zip main`
 3) if you are in using windows cd into `src/main` and type `bash` 
 
 # AWS Set up 
 
 ## Create Lambda 
 
 
 ### 1) create Spotify lambda 
 
 ![Create Spotify Lambda](src/imgs/create%20spotify%20lambda.PNG)
 
 ### 2 Upload the zip file of your lambda to aws 
 
 ![uplod lambda zip file.PNG](src/imgs/uplod%20lambda%20zip%20file.PNG)
 
 ### 3 Set env vars for lambda 
 
 ![set env vars for spotify.PNG](src/imgs/set%20env%20vars%20for%20spotify.PNG)
 
 ### 4 Create an API with AWS API Gateway 
 
 ![create API in AWS gateway.PNG](src/imgs/create%20API%20in%20AWS%20gateway.PNG)
 
 ### 5 Create Spotify Resource for your API 
 
 ![create spotify resource.PNG](src/imgs/create%20spotify%20resource.PNG)
 
 #### 5.1 Create the other endpoints/resources with its methods 
 
 ![create all the other resources login swap and refresh token.PNG](src/imgs/create%20all%20the%20other%20resources%20login%20swap%20and%20refresh%20token.PNG)
 
 #### 5.2 create the methods of the resources using lambda proxy integration 
 
 ![create methods with lambda proxy integration.PNG](src/imgs/create%20methods%20with%20lambda%20proxy%20integration.PNG)
 
 #### 5.3 all endpoints and methods listed 
 
 ![all endpoints with all methods.PNG](src/imgs/all%20endpoints%20with%20all%20methods.PNG)
 
 ### 6 deploy the your API 
 
 This is an important step if you do not deploy the api the endpoints won't be available outside API Gateway console 
 
 ![deploy API.PNG](src/imgs/deploy%20API.PNG)
 
 
 
 
 
 
 
 
 
 
