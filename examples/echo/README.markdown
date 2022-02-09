echo-server
===========

The following example implement the protocol of "echo" server.

An echo server return either in TCP or in UDP what was sent to it as-is.
The RFC is require to listen on port number 7 - A low port (bellow 1024).

The following example show to use libcap-ng to get access to the port without
using root privileges to gain that access.


libcap-ng
---------

The execution ELF file was set with the following attributes (of libcap):
  - CAP_NET_RAW - Access to raw binding (UDP).
  - CAP_NET_BIND - Access for TCP binding.
  - CAP_SET_PCAP - Gain access to change capabilities.
  - CAP_SET_FCAP - Gain access to change file(s) capabilities.

The setting of such attributes were made by using the following command:

    sudo setcap CAP_NET_RAW,CAP_NET_BIND_SERVICE,CAP_SETPCAP,CAP_SETFCAP+ep echo-server

Without the command, libcap and libcap-ng will ignore the request.
