# VertexLogger
## VertexLogger is a logger that allows you to log all the actions of your code. It is suitable for all projects, from Hello World to pet projects with a microservice architecture
![MainImage](images/MainImage.jpg)

## Documentation

VertexLogger contains several methods:
1) Info (logger/logger.go)
2) Warn (logger/logger.go)
3) Error (logger/logger.go)
When calling these methods, you need to put your text in the parameters.

Structure project:
```Structure
VertexLogger-Go/
├─project/
|   ├─logger/
|   |    └──logger.go
|   ├─testing/
|   |    └──testing.go
|   ├─go.mod
|   └──main.go
├─images/
|    └──MainImage.jpg
└── README.md
```
## Structure line in log file

[date]-[time] | [LEVEL] - [text] (2024-13:20 | [WARN] - Kekus)