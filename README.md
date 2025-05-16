# Lightweight StandardNotes Server

[![Docker Pulls](https://img.shields.io/docker/pulls/crusaders/standardnote-server)](https://hub.docker.com/r/crusaders/standardnote-server)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/mdouchement/standardfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdouchement/standardfile)](https://goreportcard.com/report/github.com/mdouchement/standardfile)
[![License](https://img.shields.io/github/license/mdouchement/standardfile.svg)](http://opensource.org/licenses/MIT)

**Portable** and **lightweight** Golang implementation of the [Standard Notes](https://docs.standardnotes.com/specification/sync) protocol for self-hosting.

### Running your own server

Create a config file `standardfile.yml`:
```
address: "0.0.0.0:5000"
no_registration: false
show_real_version: false
database_path: "/etc/standardfile/database"
secret_key: jwt-development
session:
  secret: paseto-development
  access_token_ttl: 1440h
  refresh_token_ttl: 8760h
enable_subscription: true
files_server_url: "http://localhost:5000"
```

Setup requires Docker:
`docker run -p 5000:5000 -v $(pwd)/db:/etc/standardfile/database:z -v $(pwd)/standardfile.yml:/etc/standardfile/standardfile.yml:z -it crusaders/standardnote-server`

Done! You can register and login using Standard Notes after configuring `http://localhost:5000` as custom server.


<details>
<summary>Build the image</summary>

- Requires Earthly for containerized build: https://github.com/earthly/earthly
- `git clone https://github.com/Crusader99/standardnote-server.git`
- `cd standardnote-server`
- `earthly +build`

</details>

### Technologies / Frameworks

- [Golang](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)
- [Echo](https://github.com/labstack/echo)
- [BoltDB](https://github.com/etcd-io/bbolt) + [Storm](https://github.com/asdine/storm) Toolkit
- [Gowid](https://github.com/gcla/gowid)
- [OTP](https://github.com/pquerna/otp)


## Differences to repository from mdouchement

* **2FA** (aka `verify_mfa`) implemented using [OTP-Library](https://github.com/pquerna/otp):
```
Requires enable_subscriptions=true in configuration file.
```

* Experimental support for encrypted **file upload/download**: Current state is tracked in [#93](https://github.com/mdouchement/standardfile/pull/93)

* Images are provided on [DockerHub](https://hub.docker.com/r/crusaders/standardnote-server) (including an aarch64 image for Raspberry Pi)

These features will be merged in [mdouchement's repository](https://github.com/mdouchement/standardfile) when pull requests accepted.


## Differences from reference implementation

<details>
<summary>Drop the POST request done on Extensions (backups too)</summary>

> This feature is pretty undocumented and I feel uncomfortable about the outgoing traffic from my server on unknown URLs.

</details>

<details>
<summary>Drop V1 support</summary>

> All stuff used in v1 and not in v2 nor v3

</details>

<details>
<summary>JWT revocation strategy after password update</summary>

> Reference implementation use a pw_hash claim to check if the user has changed their pw and thus forbid them from access if they have an old jwt.

<hr>

> Here we will revoke JWT based on its `iat` claim and `User.PasswordUpdatedAt` field.
> Looks more safer than publicly expose any sort of password stuff.
> See `internal/server/middlewares/current_user.go`

</details>

<details>
<summary>Session use PASETO tokens instead of random tokens</summary>

> Here we will be using PASETO to strengthen authentication to ensure that the tokens are issued by the server.

</details>

## Not working yet

- [Note revisions](https://github.com/mdouchement/standardfile/issues/31)
- [Integrity check](https://github.com/mdouchement/standardfile/issues/75)

## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

This project is licensed under the [MIT license](https://github.com/Crusader99/standardnote-server/blob/master/LICENSE).
