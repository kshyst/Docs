# Routing Protocols: BGP and IGPs

Routing protocols let routers exchange reachability information and select paths automatically. The most important design boundary is the **autonomous system (AS)**: a network or group of networks operated under one administrative policy.

- An **Interior Gateway Protocol (IGP)** carries routes *within* an AS.
- **External BGP (eBGP)** exchanges routes *between* autonomous systems.
- **Internal BGP (iBGP)** distributes BGP routes *within* one AS.

```text
                    AS 65001                              AS 65002

       +--------+      iBGP       +--------+   eBGP   +--------+
 LAN --| R1     |<--------------->| R2     |<-------->| R3     |-- LAN
       +--------+                  +--------+          +--------+
           ^                           ^                   ^
           +-------- OSPF/IS-IS -------+                   |
                    (IGP)                         OSPF/IS-IS (IGP)
```

## Choosing the Right Protocol

| Requirement | Typical choice | Reason |
| --- | --- | --- |
| Reach every router and infrastructure subnet inside one AS | OSPF or IS-IS | Fast convergence and topology awareness |
| Exchange routes with another organization or ISP | eBGP | Policy control and AS-level loop prevention |
| Carry Internet or customer routes across your own AS | iBGP | Preserves BGP attributes and routing policy internally |
| Small, simple internal network | Static routes or an IGP | Full BGP may add unnecessary complexity |

!!! note
    BGP and an IGP usually complement each other. The IGP provides internal reachability to router loopbacks and next hops; BGP carries externally learned, customer, VPN, or other policy-controlled prefixes.

## Control Plane and Forwarding Plane

A routing protocol operates in the **control plane**. It learns candidate routes and offers them to the router's routing table. The router then programs selected paths into its **forwarding plane**, which moves packets.

A route may appear in a protocol database but not in the active routing table because:

- another protocol has a more preferred route;
- the BGP next hop is unreachable;
- policy rejects the route;
- the route is invalid or suppressed;
- equal-cost multipath limits have been reached.

## Protocol Roles at a Glance

| Property | IGP | eBGP | iBGP |
| --- | --- | --- | --- |
| Scope | Inside an AS | Between ASes | Inside an AS |
| Primary goal | Fast internal reachability | Interdomain policy | Internal distribution of BGP routes |
| Common metric/input | Cost, bandwidth, delay, topology | BGP path attributes and policy | BGP path attributes and policy |
| AS_PATH changed? | Not applicable | Yes, normally prepends local AS | No |
| Default next-hop behavior | Protocol-specific | Commonly changed to self | Commonly preserved |
| Loop prevention | Protocol algorithm and topology data | AS_PATH | iBGP advertisement rules, cluster data, and AS_PATH |

## Recommended Learning Path

1. [IGP](IGP.md): understand internal reachability and convergence.
2. [BGP](BGP.md): learn BGP sessions, attributes, policy, and route selection.
3. [eBGP](eBGP.md): connect autonomous systems safely.
4. [iBGP](iBGP.md): distribute BGP routes inside an AS without creating loops.

## Design Rules of Thumb

- Keep the IGP simple and stable; advertise infrastructure prefixes rather than every service route.
- Establish iBGP between stable loopback addresses and let the IGP reach those loopbacks.
- Apply explicit import and export policy at every eBGP boundary.
- Filter prefixes and AS paths; do not trust a peer merely because the BGP session is established.
- Ensure every BGP next hop is resolvable.
- Use route reflectors deliberately and deploy at least two when availability matters.
- Monitor session state, rejected updates, route counts, next-hop reachability, and policy changes.

## Further Reading

- [RFC 4271 — Border Gateway Protocol 4](https://www.rfc-editor.org/rfc/rfc4271)
- [RFC 4456 — BGP Route Reflection](https://www.rfc-editor.org/rfc/rfc4456)
- [RFC 8212 — Default External BGP Route Propagation Behavior](https://www.rfc-editor.org/rfc/rfc8212)
- [RFC 2328 — OSPF Version 2](https://www.rfc-editor.org/rfc/rfc2328)
- [RFC 1195 — Use of IS-IS for Routing in TCP/IP](https://www.rfc-editor.org/rfc/rfc1195)
