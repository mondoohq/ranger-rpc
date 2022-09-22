# Ranger Guard Middleware

Both humans and agents can be authenticated and authorized for API access. When a request reaches the endpoint, it goes through several stages. The following diagram illustrates the flow:

                ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
                                   ranger guard                  │
                │
┌─────────────┐                                                  │
│             │ │      ┌─────────────┐        ┌─────────────┐
│    Human    │───┬───▶│   authn 1   │   ┌───▶│   authz 1   │    │
│             │ │ │    └─────────────┘   │    └─────────────┘       ┌─────────────┐
└─────────────┘   │           │          │           │           │  │             │
                │ │           │          │           │          ┌──▶│   Service   │
┌─────────────┐   │           ▼          │           ▼          ││  │             │
│             │ │ │    ┌─────────────┐   │    ┌─────────────┐   │   └─────────────┘
│    Agent    │───┘    │   authn 2   │───┘    │   authz 2   │───┘│
│             │ │      └─────────────┘        └─────────────┘
└─────────────┘                                                  │
                │
                                                                 │
                └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─

Ranger Guard allows the specification of multiple authentication and authorization modules. *Authentication* modules are called in sequence until one of authentication middleware succeeds. If the incoming request cannot be authenticated, it is rejected with 401 HTTP status code. Otherwise the user is authenticated and its identifier is made available to subsequent authorization steps. If no authentication middleware is configured, the request is just denied.

Until now, we authenticated the user, the incoming request must still be *authorized*. If no authorization middleware is defined, all request are denied. If multiple authorization modules are configured, Ranger Guard checks each one, and if any module authorizes the request, then the request can proceed. If all of the modules deny the request, then the request is denied with 403 HTTP status code.

## Authenticators

Ranger Guard uses Ranger's plugin mechanism to register its authentication modules. The aim of the authentication modules is to:

1. ensure the entity is the right one
2. ensure nobody tampered with the data in transit

This allows the server to cover 3 use cases:

1. the server is able to verify that only valid clients are talking to its api
2. ranger-generated client signs the request payload 
3. ranger-generated client use using cert pinning and verifies https cert to ensure it is the correct server

Out-of-the-box the following authentication method are available:

 * Certificate Authenticator
 * Static Token Authenticator


```
                                                            ┌───────────┐
                                                            │Google/Okta│
                                         ┌─────────────────▶│ /Keycloak │
                                         │                  └───────────┘
┌───────────────┐                 ┌─────────────┐
│      UI       │   Auth Header   │             │
│               │────────────────▶│OIDC Verifier│─────┐     ┌───────────┐
│               │                 │             │     │     │           │
└───────────────┘                 └─────────────┘     │     │           │
                                                      ├────▶│    API    │
┌───────────────┐                 ┌─────────────┐     │     │           │
│ Client + Cert │                 │             │     │     │           │
│               │────────────────▶│Cert Verifier│─────┘     └───────────┘
│               │                 │             │
└───────────────┘  Auth Header +  └─────────────┘
                  Signed Messages        │                  ┌───────────┐
                                         └─────────────────▶│    AMS    │
                                                            └───────────┘
```

## Guard Example

```
# start the server
$ cd examples/rangerguard/server
$ go run main.go

# start the client
$ cd examples/rangerguard/client
$ go run client.go 
```

## FAQ

### How can the client trust the server?

The current pattern only authenticates the client to the server but not the server to the client. Clients should use certificate pinning and verify that the https certificate of the server is valid.

### Generate certificates

This based on [How to Generate & Use Private Keys using OpenSSL's Command Line Tool](https://gist.github.com/briansmith/2ee42439923d8e65a266994d0f70180b). Ranger Guard recommends the use of elliptic curves:

```bash
# list available curves
openssl ecparam -list_curves

# generate a private key
openssl ecparam -name secp384r1 -genkey -noout -out private-key.pem

# convert to PKCS8 private key
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in private-key.pem -out private-key.p8 

# generate corresponding public key
openssl ec -in private-key.pem -pubout -out public-key.pem

# create a self-signed certificate
openssl req -new -x509 -key private-key.pem -out cert.pem -days 360
```

To generate rsa certs:

```
#!/bin/bash
openssl genpkey -algorithm RSA \
    -pkeyopt rsa_keygen_bits:3072 \
    -pkeyopt rsa_keygen_pubexp:65537 | \
    openssl pkcs8 -topk8 -nocrypt -outform pem > rsa/rsa-3072-private-key.p8

openssl pkey -pubout -inform pem -outform pem \
    -in rsa/rsa-3072-private-key.p8 \
    -out rsa/rsa-3072-public-key.pem
```

## References

- The auth method is inspired by Amazon [best-practices](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
- Go [HTTPS configuration](https://gist.github.com/denji/12b3a568f092ab951456)
- Generate a [private/public key](https://rietta.com/blog/2012/01/27/openssl-generating-rsa-key-from-command/)
- [Golang & Cryptography. RSA asymmetric algorithm](https://medium.com/@raul_11817/golang-cryptography-rsa-asymmetric-algorithm-e91363a2f7b3)
- handle [unencrypted private/public key](https://help.globalscape.com/help/secureserver3/Generating_an_unencrypted_private_key_and_self-signed_public_certificate.htm)