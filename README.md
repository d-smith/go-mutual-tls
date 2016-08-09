# Go Mutual TLS Example
How to make use of mutual TLS authentication in Go

## Notes:


Inspect the certificate:

openssl x509 -in ../cert.pem -text -noout

This sample updated the source forl with a required http/2 cipher. The client
was also updated to use plain vanilla go for making the request.
