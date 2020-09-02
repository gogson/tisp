你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Tisp

[![wercker](https://img.shields.io/wercker/ci/wercker/docs.svg?style=flat-square)](https://app.wercker.com/raviqqe/tisp/runs)
[![codeclimate](https://img.shields.io/codeclimate/github/kabisaict/flow.svg?style=flat-square)](https://codeclimate.com/github/raviqqe/tisp)
[![Go Report Card](https://goreportcard.com/badge/github.com/raviqqe/tisp?style=flat-square)](https://goreportcard.com/report/github.com/raviqqe/tisp)
[![codecov](https://img.shields.io/codecov/c/github/raviqqe/tisp.svg?style=flat-square)](https://codecov.io/gh/raviqqe/tisp)
[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Google group](https://img.shields.io/badge/join-us-ff69b4.svg?style=flat-square)](https://groups.google.com/forum/#!forum/tisp-aliens)

![logo](img/front.png)

Tisp is a "Time is Space" programming language.
It aims to be simple, canonical, and practical.

This baby project is under heavy development.
Any contributions would be appreciated.
[Join the Google group today!](https://groups.google.com/forum/#!forum/tisp-aliens)

## Current status

See [the issues](https://github.com/raviqqe/tisp/issues).
Tisp is actively developed and work in progress.
Stay tuned!

## Installation

```
go get github.com/raviqqe/tisp/src/cmd/tisp
```

You need Go 1.8+.

## Features

- Purely functional programming
  - Impure function calls in pure functions are detected and raise errors at
    runtime.
- Implicit parallelism and concurrency
  - Most of the time, you don't need to parallelize your code explicitly.
    Tisp does it all for you!
  - Inherently, programs written in Tisp run concurrently and can run
    parallelly.
- Optional injection of parallelism and causality
  - You can also increase parallelism of your code or run functions
    sequentially using `par` or `seq` primitives.
    (`par` is not implemented yet.)
- Asynchronous IO
  - Every IO can be run asynchronously by the `par` primitive.
- Dynamic typing

## Documentation

Visit [here](https://raviqqe.github.io/tisp/).

## Examples

Try scripts in [test directory](test) (`test/*.tisp`) by running:

```
tisp test/$filename.tisp
```

Some examples in [examples directory](examples) work but not all because
some features are not implemented yet.

## Contributing

Please send pull requests, issues, questions or anything else.
See also [the documentation](https://raviqqe.github.io/tisp/for_developers/)
on how to develop Tisp.

## FAQ

### What languages is Tisp inspired by?

The following is their list with ideas and technologies which come from them.

- Haskell
  - The concept of "Time is Space"
  - Lazy evaluation and data structures to realize it
- Clojure
  - Everything is a value; no object system
  - The syntax and macro system
- OCaml
  - The syntax close to pure lambda calculus and of mutual recursion
- Python
  - The Zen (See `python -c 'import this'`.)
  - The syntax of function calls with positional and keyword arguments
- Go
  - Simplicity
  - Tisp utilizes Go's coroutine runtime.

## License

[MIT](LICENSE)
