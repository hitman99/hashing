package design

import (
    . "goa.design/goa/v3/dsl"
)

var _ = API("hashing", func() {
    Title("Hashing Service")
    Description("Service for hashing strings")
    Server("hashing", func() {
        Host("0.0.0.0", func() {
            URI("http://0.0.0.0:8080")
        })
    })
})

var _ = Service("hashing", func() {
    Description("The hashing service performs hashing operations.")

    Method("sha256", func() {
        Payload(func() {
            Field(1, "string", String, "String to hash")
            Required("string")
        })

        Result(HashingResult)

        HTTP(func() {
            GET("/sha256/{string}")
        })
    })
    Method("sha512", func() {
        Payload(func() {
            Field(1, "string", String, "String to hash")
            Required("string")
        })

        Result(HashingResult)

        HTTP(func() {
            GET("/sha512/{string}")
        })
    })
    Method("md5", func() {
        Payload(func() {
            Field(1, "string", String, "String to hash")
            Required("string")
        })

        Result(HashingResult)

        HTTP(func() {
            GET("/md5/{string}")
        })
    })
    Method("sha1", func() {
        Payload(func() {
            Field(1, "string", String, "String to hash")
            Required("string")
        })

        Result(HashingResult)

        HTTP(func() {
            GET("/sha1/{string}")
        })
    })
    Method("sha384", func() {
        Payload(func() {
            Field(1, "string", String, "String to hash")
            Required("string")
        })

        Result(HashingResult)

        HTTP(func() {
            GET("/sha384/{string}")
        })
    })

    Files("/openapi.json", "./static/openapi.json")
})

var HashingResult = ResultType("HashingResult", func(){
    Description("Hash response")
    Attributes(func() {
        Attribute("type", String, "Hashing algorithm")
        Attribute("hash", String, "Hash string in hex format")
    })
})