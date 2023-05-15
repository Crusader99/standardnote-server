# Lightweight StandardNotes Server

[![Docker Pulls](https://img.shields.io/docker/pulls/crusaders/standardnote-server)](https://hub.docker.com/r/crusaders/standardnote-server)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/mdouchement/standardfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdouchement/standardfile)](https://goreportcard.com/report/github.com/mdouchement/standardfile)
[![License](https://img.shields.io/github/license/mdouchement/standardfile.svg)](http://opensource.org/licenses/MIT)

This is a 100% Golang implementation of the [Standard Notes](https://docs.standardnotes.com/specification/sync) protocol. It aims to be **portable** and **lightweight**.

### Running your own server

You can run your own Standard File server, and use it with any SF compatible client (like Standard Notes).
This allows you to have 100% control of your data.
This server implementation is built with Go and can be deployed in seconds:

- `git clone https://github.com/Crusader99/standardnote-server.git`
- `cd standardnote-server`
- `docker compose up`

<details>
<summary>Running without docker-compose is also possible</summary>

`docker run -p 5000:5000 -v $(pwd)/db:/etc/standardfile/database:z -v $(pwd)/standardfile.yml:/etc/standardfile/standardfile.yml:z -it crusaders/standardnote-server`

</details>

### Technologies / Frameworks

- [Golang](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)
- [Echo](https://github.com/labstack/echo)
- [BoltDB](https://github.com/etcd-io/bbolt) + [Storm](https://github.com/asdine/storm) Toolkit
- [Gowid](https://github.com/gcla/gowid)
- [OTP](https://github.com/pquerna/otp)


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

## Differences to repository from mdouchement


* Provides updated **subscription** premium features out of the box:
```
This option enables paid features in the official StandardNotes client.
If you want to enables these features, you should consider to
donate to the StandardNotes project as they say:

Building Standard Notes has high costs. If everyone evaded contributing financially,
we would no longer be here to continue to build upon and improve these services for you.
Please consider donating to https://standardnotes.com/donate if you do not plan on purchasing a subscription.
https://docs.standardnotes.com/self-hosting/subscriptions/

This project https://github.com/mdouchement/standardfile does not intend to
conflict with the business model of StandardNotes project or seek compensation.
```

* 

* **2FA** (aka `verify_mfa`) implemented using [OTP-Library](https://github.com/pquerna/otp):
```
Requires enable_subscriptions=true in configuration file.
```

* Experimental support for encrypted **file upload/download**: Current state is tracked in [#93](https://github.com/mdouchement/standardfile/pull/93)

These features will be merged in [mdouchement's repository](https://github.com/mdouchement/standardfile) when pull requests accepted.

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

[MIT](https://github.com/Crusader99/standardnote-server/blob/master/LICENSE)