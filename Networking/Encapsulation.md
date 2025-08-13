# Encapsulation in TCP/IP

- Application data: It all starts when the user inputs the data they want to send into the application. For example, you write an email or an instant message and hit the send button. The application formats this data and starts sending it according to the application protocol used, using the layer below it, the transport layer.
- Transport protocol segment or datagram: The transport layer, such as TCP or UDP, adds the proper header information and creates the TCP segment (or UDP datagram). This segment is sent to the layer below it, the network layer.
- Network packet: The network layer, i.e. the Internet layer, adds an IP header to the received TCP segment or UDP datagram. Then, this IP packet is sent to the layer below it, the data link layer.
- Data link frame: The Ethernet or WiFi receives the IP packet and adds the proper header and trailer, creating a frame.

![encap](img/encap.svg)

## The life of a packet

- On the TryHackMe search page, you enter your search query and hit enter.
- Your web browser, using HTTPS, prepares an HTTP request and pushes it to the layer below it, the transport layer.
- The TCP layer needs to establish a connection via a three-way handshake between your browser and the TryHackMe web server. After establishing the TCP connection, it can send the HTTP request containing the search query. Each TCP segment created is sent to the layer below it, the Internet layer.
- The IP layer adds the source IP address, i.e., your computer, and the destination IP address, i.e., the IP address of the TryHackMe web server. For this packet to reach the router, your laptop delivers it to the layer below it, the link layer.
- Depending on the protocol, The link layer adds the proper link layer header and trailer, and the packet is sent to the router.
- The router removes the link layer header and trailer, inspects the IP destination, among other fields, and routes the packet to the proper link. Each router repeats this process until it reaches the router of the target server.