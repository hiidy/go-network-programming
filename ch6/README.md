# UDP Datagram Programming - Speed, Unreliability, and Application-Layer Compensation

## prob

Since TCP is already so reliable, why does the industry (e.g., in gaming or the QUIC protocol) go to such great lengths to re-implement a reliability mechanism on top of UDP? Compared to TCP, what crucial "reliability" components are missing from the simple "ACK + Retransmit" protocol we wrote today?

Hint: This question will lead you to think about TCP's "Head-of-Line Blocking" problem and the immense complexity involved in building a truly industrial-grade reliable UDP protocol.

## ans

The core issue is **TCP's Head-of-Line Blocking problem**. TCP guarantees in-order delivery of all data, which paradoxically becomes a liability. If packets 1, 2, 3, 4, 5 are sent and packet 2 is lost, packets 3, 4, and 5 that have already arrived cannot be delivered to the application—they must wait until packet 2 is retransmitted. In gaming, this is fatal: fresh position updates are blocked waiting for stale data that's already useless by the time it arrives.

Looking at what's missing from the simple ACK + retransmit protocol we wrote today: First, there's no **congestion control**. Simply sending as fast as possible and retransmitting losses will overwhelm the network and cause even more packet loss. There's no **flow control** either, so a fast sender can overflow a slow receiver's buffer. There's no **RTT measurement and adaptive timeouts**—with fixed timeouts, you get unnecessary retransmissions on slow networks and slow recovery on fast networks. There's no **Selective ACK (SACK)**, so even if only one packet in the middle is lost, everything after it must be retransmitted.

The reason QUIC re-implements reliability on top of UDP is to create **multiple independent streams** within a single connection so that packet loss in one stream doesn't block others, to enable **connection migration** so connections survive switching from WiFi to cellular, and to **integrate encryption by default** to reduce round trips from TLS handshakes. Ultimately, it means overcoming TCP's limitations while re-implementing all the complex mechanisms TCP has developed over decades.
