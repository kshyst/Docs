# HTTP Streaming

Traditional streaming protocols (RTSP, RTP, RTMP) create direct, continuous connections between the client and server, offering low-latency streaming often used for real-time applications. However, these protocols require dedicated media servers and specific network configurations, limiting scalability and compatibility.

HTTP streaming, in contrast, uses standard web servers and is firewall-friendly. This method is more scalable, thanks to its compatibility with HTTP-based CDNs, and it supports adaptive bitrate streaming. Protocols like WebRTC, meanwhile, enable real-time peer-to-peer streaming, but they're less suited for large-scale content distribution.

## How does HTTP streaming work?

HTTP streaming works by breaking media content into small, sequential chunks (e.g., 2–10 seconds) and delivering them over standard HTTP protocol. A manifest file (like .m3u8 for HLS or .mpd for DASH) lists available chunk URLs and quality options. The client downloads and plays each chunk immediately, enabling near real-time playback. Adaptive Bitrate Streaming (ABR) adjusts chunk quality based on network speed, reducing buffering. HTTP streaming also leverages CDNs for caching, improving load times and scalability by distributing content across multiple servers.