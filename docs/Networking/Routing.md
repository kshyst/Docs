# Routing

Routing is the process of selecting a path for packets between networks. Routers can learn routes from directly connected interfaces, static configuration, or dynamic routing protocols.

Dynamic routing protocols fall into two broad operational categories:

- **Interior Gateway Protocols (IGPs)** exchange routes inside one administrative domain. Common examples are OSPF, IS-IS, RIP, and EIGRP.
- **Border Gateway Protocol (BGP)** exchanges policy-controlled reachability. eBGP operates between autonomous systems, while iBGP distributes BGP routes inside one autonomous system.

## Routing Protocol Guides

- [BGP fundamentals](Routing%20Protocols/BGP.md)
- [External BGP (eBGP)](Routing%20Protocols/eBGP.md)
- [Internal BGP (iBGP)](Routing%20Protocols/iBGP.md)
- [Interior Gateway Protocols (IGPs)](Routing%20Protocols/IGP.md)
- [BGP and IGP overview](Routing%20Protocols/index.md)

## Route Selection Layers

When forwarding a packet, a router generally:

1. chooses the route with the longest matching prefix;
2. if identical prefixes come from different sources, applies its protocol or administrative preference;
3. if one protocol offers multiple paths, applies that protocol's metric and tie-breakers;
4. programs the selected next hop into the forwarding table.

A routing protocol's best path is not necessarily usable: its next hop must also resolve through the routing table.
