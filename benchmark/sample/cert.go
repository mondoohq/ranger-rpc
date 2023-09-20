// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package sample

// openssl genrsa -out server.key 2048
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// TODO: regenerate!!!

var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIICvjCCAaYCCQClVNOARrZ/XDANBgkqhkiG9w0BAQsFADAhMQswCQYDVQQGEwJV
UzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTE4MDcyMzE5MDAwNFoXDTI4MDcyMDE5
MDAwNFowITELMAkGA1UEBhMCVVMxEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJ
KoZIhvcNAQEBBQADggEPADCCAQoCggEBAMbSFS6JhnbEpM0xh/JfqZ8SxQ5yDkKE
JBSgUi/x3+iwW093loKZhYzBHR90Oy9whOwavbpFacRI/Qy1j25vYDqJ9nVCze7n
GoZtH/ucRXR2ugdhtfhWZawZ1KA9+SDtQC5hfIvqczID/d6lruY8h6bjBvZJB71W
aS5EKJITNRQJPFhAu7TwnuP/uvpaTcs/3BvGcB1XHHSZolnHgp5ONepzhfScex12
mgfO26vecplHE35OduAqrb88gtdcPQBDt11sQiAJ234AfQwMvniVUD4vtx3XPIKg
CjazUMtrbhjpY4FyONSRfCMvEqGXXZ1+2UfztsG8VQUZt161g0oIZJECAwEAATAN
BgkqhkiG9w0BAQsFAAOCAQEAv9OYlWmMzQNvG2htkrxgeymRHVLZX0V7Krsfttqc
Uik5Zvd+VtO9vlxMGb7Qno5xH6ZHbAhODKagFCSqZVsyZIjWv12KG8aKeMGY8ltC
zJKAkRXbSfYk8wU28nAXjliZ80ItuzfwgXZs77WTGdrSJ+w9v+umxYERP+X1X6ad
JKZbkbNTFR+ScNbZ//bTUegHtZN1h3qqQrjQ7JzZqfIvHaNDvGeILYJeg3GsrcpR
sENqM4ofD4TXVt6LhW9FKjQyd7IEkFP61yVPGt8akJ22LUBYYcE43ReGSfO4IR2x
BKfIZsBVfazmsugT8EemEsBl7srtkMq5vdHMnBJCPgkKmg==
-----END CERTIFICATE-----`)

var LocalhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAxtIVLomGdsSkzTGH8l+pnxLFDnIOQoQkFKBSL/Hf6LBbT3eW
gpmFjMEdH3Q7L3CE7Bq9ukVpxEj9DLWPbm9gOon2dULN7ucahm0f+5xFdHa6B2G1
+FZlrBnUoD35IO1ALmF8i+pzMgP93qWu5jyHpuMG9kkHvVZpLkQokhM1FAk8WEC7
tPCe4/+6+lpNyz/cG8ZwHVccdJmiWceCnk416nOF9Jx7HXaaB87bq95ymUcTfk52
4CqtvzyC11w9AEO3XWxCIAnbfgB9DAy+eJVQPi+3Hdc8gqAKNrNQy2tuGOljgXI4
1JF8Iy8SoZddnX7ZR/O2wbxVBRm3XrWDSghkkQIDAQABAoIBAG6goqycVTdsv451
SwGv/P/IP+GD0S9tu23GrzSCT2Z4CrazAgp1RfxFz+CamfwRjcSaNIua5/kR68vQ
kpiOXGr2LS6eF1whN38o5SzpjTP6hBRraAdge35BeTAYi7Cokpe8IsLvl11zHyVt
512wvII9vLf5dtcBZ9EYl8J/8X7NkbC/+x87cpPRkhkABYMMCKWfdL3SsvJ/3qLk
8WMy3Lgj02lG1EluQFg/h7hGrOBWiw8CzfQC1sjEVGixA4U18OTRJVw5ZIp1aMxj
KLBACCoIMYuRwiyPnqc/WY+9kENM6LaaJacZ0ZXYjX7sj0xm6SaW2zBF7E/Dyerz
8up43CkCgYEA5NPyLnRBBHOugq7KadJIq0KQ+6SePcMm8EaLS1MyvlMiQXAAyN5F
K8eg5VpCUQ5378xFZnUe/y4D+YC25LbZ1QdLG5hTDhKcW35+poSa33CWllUMGbZd
QpFlu1OMN0J40VFMVRKG8y0HpbPSxempQ9MDs3uw6Ach/EZ/317fICMCgYEA3m30
6xOgRbzZ0f4EBAT5L8w/rIMdtAQHQ5I7+a7K/NszBxGzAawH9OjOeR8Ih4oG5f2d
wpJwCaL1ztFOukmn9NKKuDfx4R6MRys3QnxcxvgVBjF/izb0tM+0jRB75+j6ne3p
vQ28sJYLxTgOQutZUJBNaV4260Y4IwC/s4O9mbsCgYA0FX3pTvLBlach/bD61y9N
M/CWJokSG8pQJG5uLbi+E2QXquuyzMzHwz9/FMVFd9qazU76nCv6/zlOYBrBAlGg
qHFTDZ/R8zB9rtQbCNHLi+/qtd70N0sQ7NFQCxs+NLYVRsDuGDJ5RUWZVM3j2GR7
mJseDkhc98qnhlBywkBdKwKBgBq6NbrlodWfasEb99mPy22d6mzNWI1gCotpEAHh
qgyWPlx0GQFzbYVVUDIns3ut70RFpGZT+FiAF29hoUcrQJ5fikG2nz8Az7RhkgNQ
NEnIV6Zl3kCZOvBbIQPuXiUwzqSZiQOpmenSLdnl8XjDFPlkTZkCtDCzQF2cYmys
wOSvAoGAe96Lu278OUia0orkoX/Zg8PxQxt2NlpjiGdEvUdNVamWm3yGa6meezFT
fd5FFKSciIHDiqKQ1IakGypW8XeRLScxISaFYeCoAxLBKTsCGhsxCaBZRPKrzk1D
NZNW6nmd1ItvUwupVPiaABfgbh0YJ2MJYZSqWBkVye+1JzezfHk=
-----END RSA PRIVATE KEY-----`)
