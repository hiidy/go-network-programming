# 07.Socket Options Fine-Tuning the Kernel Network Stack

## prob
Today we discussed SO_REUSEADDR in detail, which allows a new listening service to bind to an address still occupied by the TIME_WAIT state.

In the configuration of some high-performance Web servers (like Nginx), you might see another option with a very similar name: SO_REUSEPORT.

Please investigate: What is the core difference between SO_REUSEPORT and SO_REUSEADDR? What completely different problem does SO_REUSEPORT solve? How should we set it in Go?

This question will lead you into a more advanced field of server performance optimization—utilizing "Kernel-Level Load Balancing" to process new connections in parallel across multi-core CPUs.

## ans

---

## Core Difference

| Option | Problem it solves |
|--------|------------------|
| `SO_REUSEADDR` | Allows **sequential** restart of the same application by coexisting with sockets in `TIME_WAIT` state |
| `SO_REUSEPORT` | Allows **multiple processes to simultaneously listen** on the same IP:Port for kernel-level load balancing |

## What Problem Does SO_REUSEPORT Solve?

In a typical single-process server, all incoming connections go through one `accept()` call on one thread. Even with multiplexing (epoll/kqueue), this becomes a bottleneck when handling tens of thousands of connections per second.

`SO_REUSEPORT` solves this by allowing multiple processes to bind to the same port simultaneously. The kernel then distributes incoming connections across these processes, effectively **parallelizing the accept() operation** and maximizing single-machine performance before scaling out.

## How to Set It in Go

```go
lc := net.ListenConfig{
    Control: func(network, address string, c syscall.RawConn) error {
        var opErr error
        e-rr := c.Control(func(fd uintptr) {
            opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
        })
        if err != nil {
            return err
        }
        return opErr
    },
}

listener, err := lc.Listen(context.Background(), "tcp", ":8080")
```

Note: All processes sharing the same port must set `SO_REUSEPORT`. Once a connection is accepted, it belongs exclusively to that process—TCP's connection-oriented nature remains intact after the handshake.
