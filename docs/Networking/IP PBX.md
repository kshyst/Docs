# IP PBX (Internet Protocol Private Branch Exchange)

A business telephone system that delivers voice and video via IP.

A modern switchboard

## How It Works

Traditional PBX systems (Time Division Multiplexing TDM) requires seperate network for voice and data.

- **Digitization**: Speaking into an IP phone converts analog voice signals into digital packets.
- **Compression**: Compressed using Codec (G.711, G.729, etc) to save bandwidth.
- **Transmission**: Data travels over LAN to the IP PBX server.
- **Routing**: IP PBX server determines where the call should go.

### Internal Call

Routes data directly to another IP phone on the network.

### External Call

Sends the data through a `SIP Trunk` (a virtual phoneline) over the internet to reach the recipient on the standard telephone network (PTSN) or mobile network.

#### SIP (Session Initial Protocol)

Standard signaling protocol used to set up, connect, and disconnect communication sessions.

## Components of an  IP PBX System

### IP PBX Server

A computer or a VM that manages all call switching and routing

Used Softwares:

- Asterisk
- 3CX
- FreePBX
- Avaya Aura
- Cisco Unified Communication Manager

### Endpoints (Phones)

- IP Hardphones: Physical desk phone. Looks like normal phones but plugs into an Ethernet port (RJ45) rather than a phone jack (RJ11).
- Softphones: Software applications installed on computers or smartphones

### SIP Trunks / VoIP Provider

Instead of physical wires from telephone company. Provided by an ITSP (Internet Telephony Service Provider) that connects your IP PBX to outside world via internet.

### VoIP Gateway

Old analog phone lines (POTS) or legacy hardware device like fax, This gateway is used to convert analog signals into IP packets.

## Key Features

- **Unified Communication**
- **Voicemail to Email**
- **IVR (Interactive Voice Response)**
- **Call Queues**
- **Find Me / Follow Me**: Calls can ring a desk phone, then a mobile phone, then a laptop sequentially or simultaneously.
- **CDR (Call Detail Records)** : Detailed logs of call duration, source, and destination for billing and analysis.

```ini
; sip.conf configuration for a generic extension

[1001]
type=friend              ; Can make and receive calls
context=from-internal    ; Security context
host=dynamic             ; The IP address of the phone can change
secret=SuperSecurePass!  ; The password for the phone to register
disallow=all             ; Disable all codecs first
allow=ulaw               ; Allow G.711 u-law
allow=alaw               ; Allow G.711 a-law
dtmfmode=rfc2833         ; Standard for sending key presses
callerid="John Doe" <1001>
nat=force_rport,comedia  ; Network Address Translation settings
qualify=yes              ; Check if the device is reachable
```

