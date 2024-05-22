# Ranger RPC

Ranger RPC is a simple and fast proto-based RPC framework

## Design Goals

- simple & fast to use
- work with standard Protocol Buffers service definitions
- works with GO's standard http server and hence does not require HTTP 2
- minimal runtime dependencies
- make it easy to generate json/yaml/protobuf based APIs

## References

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [GRPC Service Definition](https://grpc.io/docs/what-is-grpc/core-concepts/#service-definition)

## Kudos

This implementation is inspired by:

- [Rob Pike's video](https://www.youtube.com/watch?v=ENLWEfi0Tkg)
- [GRPC](https://grpc.io/)
- [Twirp](https://github.com/twitchtv/twirp).

We at Mondoo are huge NASA fans and we want to dedicate the name to the NASA [Ranger missions](https://science.nasa.gov/missions/ranger/) whose objective was to obtain the first close-up images of the surface of the Moon. These missions, which were the first American spacecraft to land on the moon, helped lay the groundwork for the Apollo program.

## Authors

- Christoph Hartmann
- Dominik Richter
