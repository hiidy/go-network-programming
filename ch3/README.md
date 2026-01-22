# ch3 - TCP Connection Lifecycle-Visualizing the Three-Way Handshake and Four-Way Handshake

## prob

If the server actively calls conn.Close() to close the connection, what would the four-way handshake process look like? On which side would the TIME_WAIT and CLOSE_WAIT states appear respectively?

##  ans

If the server actively calls `conn.Close()`:

**Four-way handshake process:**

1. Server → Client: FIN (Server initiates close)
2. Client → Server: ACK (Client acknowledges)
3. Client → Server: FIN (Client also closes)
4. Server → Client: ACK (Server acknowledges)

**State distribution:**

- **TIME_WAIT**: Appears on the **server** (the active closer always enters TIME_WAIT)
- **CLOSE_WAIT**: Appears on the **client** (the passive closer enters CLOSE_WAIT until it calls `Close()`)

The rule is simple: **whoever calls `Close()` first gets TIME_WAIT**.
