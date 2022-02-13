gocapng
=======

The following code
-------------------

The following code provides a (C)Go binding for libcap-ng by Red-Hat using
__static__ linking.

The aim of the binding is to provide a Go'ish approach on using libcap-ng's API.
The code is only for Linux without a support for any additional OSes.

What is it all about
--------------------

The aim of POSIX's libcap and libcap-ng is to provide elevated permissions for
specific type of actions without becoming root, and only open it when required
and turn it off after performing the call itself.

In order to bind low port number (bellow 1024) you need to become root, or use
CAP_NET_BIND_SERVICE with libcap or libcap-ng.

The difference between libcap and libcap-ng is that libcap-ng is a library that
provides helps for taking care of POSIX libcap in simpler manner.


Building Requirements
---------------------

Every executable binary that requires to have capabilities attributes enabled.

The way to set capabilities is to use `setcap` command that arrive from `libcap`.

```shell
$ sudo setcap CAP_ .... <executable>
```

Without using `setcap`, the executable will not gain any required permission.



