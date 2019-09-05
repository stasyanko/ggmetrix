# ggmetrix

Ggmetrix collects metrics from multiple (possibly distributed) sources with ease, that is provides metrics aggregation. For example you have a job queue and you want to measure the number of jobs per minute or you have a bunch of microservices and you want to get the number of requests to your system.

Ggmetrix is currently under development, pull requests are welcome!
Binary version is to be released, but you can compile it yourself like this:

`npm install`
`npm run prod`
`go build main.go`

Rename .env.example to .env and set your environmental variables.

<p align="center">
  <img width="450" height="400" src="https://raw.githubusercontent.com/stasyanko/ggmetrix/master/docs/images/workflow.png">
</p>

#### Features
 - Your custom SQL databases for metrics
 - Simple to setup - just upload binary to your server
 - Client libs for popular languages (in development)
 - Metrics aggregation out of the box
 - Admin panel included 

#### Why ggmetrix?

Monitoring systems are usually hard to setup and mantain. Ggmetrics is supposed to solve these problems and offers super simple setup and more awesome features like simple metrics aggregation. Yeah, I wanted to use metrics aggregation in my project, but this feature is either offered by paid services, or it is hard to setup in systems like Prometheus (though for other purposes it is great). That's why ggmetrix was created.

Created with :heart: with such great tools as gin framework, gorm, react.js.

#### TODO:
- ~~counter metrics type~~
- gauge metrics type
- basic auth for admin panel
- auth in header like jwt
- client libs for node.js, php, go
