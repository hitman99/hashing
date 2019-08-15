# Hashing API
This api will hash your damn strings, that is all.
Supported hashes (and corresponding endpoints):
* md5 - `/md5/<your string>`
* sha1 - `/sha1/<your string>`
* sha256 - `/sha256/<your string>`
* sha384 - `/sha384/<your string>`
* sha512 - `/sha512/<your string>`

## How to use it

To get a hash of your string, use the appropriate endpoint and provide a string like so:
```shell script
$ curl https://hashing.movetokube.com/sha1/whatever
{"type":"Sha1","hash":"d869db7fe62fb07c25a0403ecaea55031744b5fb"}
```

You'll get a JSON string in the response. The uppercased hash type in `type` field
and your hashed string in `hash` field (as in example above).

## Disclaimer
This application is not designed to save strings or hashes in any way. 
