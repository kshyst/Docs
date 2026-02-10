# Glossary

**ACD (Automatic Call Distributor)** - A device or system that distributes incoming calls to a specific group of terminals that agents use. It is often part of a computer telephony integration (CTI) system.

**CODEC (Coder/Decoder)** – A software library that contains the algorithms necessary to convert an analog signal to and from a digital one. Examples: G.711 G.729 GSM

**Context** – The dialplan is composed of one or more extension contexts. Each extension context is itself simply a collection of extensions. Each extension context in a dialplan has a unique name associated with it. The use of contexts can be used to implement a number of important features, such as security, routing, autoattendant, multilevel menus, authentication, callback, privacy, macros, etc…

**DAHDI (Digium Asterisk Hardware Device Interface)** – A high density kernel telephony interface for PSTN hardware.

**Dialplan** – A dial plan establishes the expected number and pattern of digits for a telephone number. This includes country codes, access codes, area codes and all combinations of digits dialed. For instance, the North American public switched telephone network (PSTN) uses a 10-digit dial plan that includes a 3-digit area code and a 7-digit telephone number. Most PBXs support variable-length dial plans that use 3 to 11 digits. Dial plans must comply with the telephone networks to which they connect.

**E&M (Ear & Mouth)** - A type of signaling commonly used over T1 and E1 interfaces.

**Encode** – The process of converting an analog signal into a digital signal that can be manipulated easily by a computer.

**FXO (Foreign Exchange Office)** – A device usually found on the customer end that is powered by the channel and can interface into the telephone company’s network. Digium makes FXO modules that interface with PSTN lines using FXS signalling in either Loopstart(fxs_ls) or the more common Kewlstart(fxs_ks) modes.

**FXS (Foreign Exchange Station)** – A device usually located on the telephony company’s property, a FXS device send power through a channel to a phone on the other end. Digium makes FXS modules that interface with PSTN phones using FXO signalling in either Loopstart(fxo_ls) or the more common Kewlstart(fxo_ks) modes.

**G.711** – An uncompressed codec that samples a 64kbps channel at 8 bits per sample using pulse code modulation. The Two varients of G.711 are known formally as uLaw and aLaw.

**G.729** – The G.729 codec is an industry standard which allows for stuffing more calls in limited bandwidth to utilize IP voice in more cost effective ways. A typical call consumes 64Kbps of voice bandwidth. G.729 reduces the call to 8Kbps (normal IP overhead adds to this number). Many people are using Asterisk with G.729 to replace expensive gateways.

**GSM** – A compressed speech codec that uses a rate of 13 kbps.

**H.323** – A VoIP protocol that was deployed early and is widely adopted.

**IAX (Inter-Asterisk eXchange)** – A VOIP protocol designed to be much more NAT friendly. IAX currently only transports audio.

**IVR (Interactive Voice Response)** – An automated voice system that allows callers to navigate a phone system and be directed to the correct extension by pressing a series of numbers on a tuch-tone phone. (I.E. Push 1 for sales, push 2 for support, etc..)

**MGCP (Media Gateway Control Protocol)** – A VOIP Protocol that has both signaling and control and was designed to reduce complexity between media gateways.

**PBX (Private Branch Exchange)** – A telephone exchange that serves a particular business or office, as opposed to one that a common carrier or telephone company operates for many businesses or for the general public.

**PRI (Primary Rate Interface)** – A PRI is a truly digital circuit, containing 24 ISDN channels. One of these channels is a D channel and used for signaling. The rest are B channels and used to transport audio.

**PSTN (Public Switched Telephone Network)** – Originally a network of fixed-line analog telephone systems, the PSTN is now almost entirely digital and includes mobile as well as fixed telephones. The network works in much the same way that the Internet is the network of the world’s public IP-based packet-switched networks.

**REN (Ringer Equivalency Number)** – A number which represents the ringer loading effect on a line. A ringer equivalency number of 1 represents the loading effect of a single traditional telephone set ringing circuit. Most modern telephones probably will have a REN lower than 1. The total REN expresses the total loading effect of the equipment on the ringing current generator (FXS). The REN of a Digium FXS board is 5 (representing “extension,” i.e., parallel-connected telephones). The actual number of devices on the line may be greater than the REN limit, if their respective individual RENs are less than 1.

**SIP (Session Initiation Protocol)** – A signaling protocol, widely used for controlling multimedia communication sessions such as voice and video calls over Internet Protocol (IP). SIP adoption amongst hardware and software vendors continues to expand.

**TDM (Time Division Multiplexing)** – A processes of splitting one medium into two or more channels by using timed segments to transmit information.

**Transcode** – The process of converting a channel with one type of encoding to a different type of encoding in real time.

**VoIP (Voice Over Internet Protocol)** – A general method for transporting voice through the internet.

**Zaptel** – The Zaptel project has been renamed ‘DAHDI’ as of May 2008. DAHDI is a series of drivers for telephony hardware devices.