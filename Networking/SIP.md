# SIP (Session Initiator Protocol)

SIP is a signaling protocol used to start, maintain, and end communication sessions. Think of it as the “handshake” or the “digital operator.” It does not carry your voice; it simply sets up the connection so your voice can travel.


- **Setup (SIP INVITE)**: The caller’s phone sends a message to Asterisk saying, “I want to talk, and here are the audio formats (Codecs) I support.”
- **Negotiation (SDP)**: Asterisk looks at the request. If it accepts, they agree on a format (e.g., G.711 or Opus). This negotiation happens inside the SIP packet body using a sub-protocol called SDP.
- **Connection (200 OK / ACK)**: Asterisk answers the call.
- **The Split (SIP vs. RTP)**:
  - **SIP** sits idle, monitoring the line to see if anyone hangs up.
  - **RTP** (Real-time Transport Protocol) takes over. This is the actual “Stream” arrow in your diagram. RTP carries the voice data from the user to Asterisk.
- **Teardown (BYE)**: When the “Room Type Microservice” sends an “Action” to hang up, Asterisk sends a SIP BYE message to the caller’s phone to end the connection.