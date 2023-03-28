package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // create a new Gin router
    r := gin.Default()

    // define an endpoint that returns "Hello, world!" when accessed
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, world!"})
    })

    // define an endpoint for communication with the frontend
    r.POST("/frontend", func(c *gin.Context) {
        // extract the request body as a JSON object
        var request map[string]string
        if err := c.ShouldBindJSON(&request); err != nil {
            // return an error if the request body is not valid JSON
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }

        // extract the "message" field from the request
        message, ok := request["message"]
        if !ok {
            // return an error if the "message" field is missing
            c.JSON(400, gin.H{"error": "missing message field"})
            return
        }

        // do something with the message (e.g. log it)
        log.Printf("frontend message: %s", message)

        // return a success response
        c.JSON(200, gin.H{"status": "success"})
    })

    // start the server on port 8080
    r.Run(":8080")
}
