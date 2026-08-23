# Internal BGP (iBGP)

**Internal BGP (iBGP)** is BGP peering between routers in the same autonomous system. It distributes externally learned and locally originated BGP routes across that AS while preserving BGP policy and path attributes.

An IGP and iBGP solve different problems:

- the **IGP** reaches router loopbacks, internal links, and BGP next hops;
- **iBGP** carries Internet, customer, VPN, or other policy-controlled routes.

```text
                         AS 65001

             iBGP                          iBGP
       R1 <-----------> R2 <------------> R3
       |                 |                  |
  eBGP to ISP A       Route             eBGP to ISP B
                    Reflector

       +------------- OSPF or IS-IS -------------+
              provides internal reachability
```

## Defining Behavior

An iBGP session normally has these properties:

- both peers use the same ASN;
- advertising a route does not prepend the local ASN to `AS_PATH`;
- `NEXT_HOP` is often preserved;
- attributes such as `LOCAL_PREF` are carried across the AS;
- routes learned from one ordinary iBGP peer are not advertised to another ordinary iBGP peer.

The final rule is commonly called **iBGP split horizon**. Because the ASN is not added inside the AS, `AS_PATH` alone cannot prevent an internal advertisement loop.

## The Full-Mesh Requirement

Without route reflectors or confederations, every iBGP-speaking router must peer with every other iBGP-speaking router.

For `n` routers, the number of sessions is:

\[
\frac{n(n-1)}{2}
\]

| Routers | Full-mesh sessions |
| ---: | ---: |
| 3 | 3 |
| 10 | 45 |
| 50 | 1,225 |
| 100 | 4,950 |

A full mesh is simple and avoids route-reflector path-hiding behavior, but it becomes operationally expensive as the network grows.

## Route Reflectors

A **route reflector (RR)** relaxes the full-mesh requirement. Its clients peer with the reflector rather than with every other client.

```text
                 +----------------+
                 | Route Reflector|
                 +----------------+
                   /      |      \
                Client  Client  Client
```

A route reflector can advertise:

- a route from an eBGP peer to clients and non-clients;
- a route from a client to other clients and non-clients;
- a route from a non-client to clients, but not to other non-clients.

Route reflection adds `ORIGINATOR_ID` and `CLUSTER_LIST` information for loop prevention. Deploy redundant reflectors for availability, but remember that redundancy does not automatically guarantee identical best paths.

### Route-Reflector Example

Reflector:

```text
router bgp 65001
 bgp router-id 10.0.0.1
 neighbor 10.0.0.2 remote-as 65001
 neighbor 10.0.0.2 update-source lo
 neighbor 10.0.0.3 remote-as 65001
 neighbor 10.0.0.3 update-source lo
 !
 address-family ipv4 unicast
  neighbor 10.0.0.2 activate
  neighbor 10.0.0.2 route-reflector-client
  neighbor 10.0.0.3 activate
  neighbor 10.0.0.3 route-reflector-client
 exit-address-family
```

Client:

```text
router bgp 65001
 bgp router-id 10.0.0.2
 neighbor 10.0.0.1 remote-as 65001
 neighbor 10.0.0.1 update-source lo
 !
 address-family ipv4 unicast
  neighbor 10.0.0.1 activate
 exit-address-family
```

The IGP must provide reachability between `10.0.0.1` and `10.0.0.2`.

## Confederations

A **BGP confederation** divides a public AS into multiple internal member ASes. Sessions between member ASes behave similarly to eBGP internally, while external peers see the confederation as one AS.

Confederations can reduce full-mesh scaling pressure and create policy boundaries, but they add operational complexity. Route reflection is more common in many modern networks. Avoid combining both techniques unless there is a clear design reason.

## NEXT_HOP and `next-hop-self`

When an edge router advertises an eBGP-learned route to an iBGP peer, it commonly preserves the external next hop.

```text
ISP next hop: 192.0.2.1
        |
Edge router: 192.0.2.2 / loopback 10.0.0.2
        |
Core iBGP router
```

The core router must be able to resolve `192.0.2.1`. There are two common designs:

1. advertise the external link into the IGP; or
2. configure the edge router to advertise itself as the next hop.

```text
router bgp 65001
 address-family ipv4 unicast
  neighbor 10.0.0.3 next-hop-self
 exit-address-family
```

`next-hop-self` is often easier to operate and prevents the IGP from needing every external peering subnet. Exact behavior for reflected routes may require platform-specific options such as forcefully rewriting the next hop.

## Loopback-Based Sessions

iBGP sessions are commonly sourced from loopback interfaces:

```text
router bgp 65001
 neighbor 10.0.0.2 remote-as 65001
 neighbor 10.0.0.2 update-source lo
```

A loopback remains logically up while at least one physical path exists, allowing the IGP to reroute the TCP session after a link failure. The BGP peer address and configured update source must agree at both ends.

## Synchronization and Route Propagation

Old BGP implementations had a **synchronization rule** requiring iBGP-learned routes to appear in the IGP before use. This rule is obsolete and disabled by default on modern platforms. Do not redistribute the full Internet table into an IGP.

Instead:

- use the IGP for infrastructure reachability;
- use iBGP to carry BGP routes;
- ensure BGP next hops resolve through the IGP;
- apply policy at controlled boundaries.

## Route-Reflector Design Considerations

Route reflectors select and advertise BGP paths rather than merely relaying every path. This can cause **path hiding**: a reflector's best route may not be the best route from a client's location.

Mitigations include:

- placing reflectors according to network topology;
- using redundant reflector clusters;
- keeping client-to-reflector IGP costs intentional;
- using ADD-PATH where supported and justified;
- validating failure and traffic-engineering scenarios in a lab.

Do not place route reflectors directly in the forwarding path merely because they are control-plane hubs. Many designs separate route reflection from packet forwarding.

## Minimal iBGP Example

Two routers in AS 65001 peer using loopbacks:

```text
! Router R1
router bgp 65001
 bgp router-id 10.0.0.1
 neighbor 10.0.0.2 remote-as 65001
 neighbor 10.0.0.2 update-source lo
 !
 address-family ipv4 unicast
  neighbor 10.0.0.2 activate
  neighbor 10.0.0.2 next-hop-self
 exit-address-family
```

```text
! Router R2
router bgp 65001
 bgp router-id 10.0.0.2
 neighbor 10.0.0.1 remote-as 65001
 neighbor 10.0.0.1 update-source lo
 !
 address-family ipv4 unicast
  neighbor 10.0.0.1 activate
 exit-address-family
```

This configuration assumes the IGP already carries `10.0.0.1/32` and `10.0.0.2/32`.

## Troubleshooting iBGP

| Symptom | Likely checks |
| --- | --- |
| Session will not establish | Loopback reachability, update source, TCP/179, ASN, authentication |
| Route appears on one router only | Full-mesh rule, missing RR-client configuration, export policy |
| Route is present but invalid | Unreachable `NEXT_HOP`, missing IGP route |
| Unexpected path through the network | `LOCAL_PREF`, RR topology, IGP cost, path hiding |
| Duplicate or unstable paths | RR cluster configuration, inconsistent policy, next-hop changes |
| VPN/EVPN routes absent | Address-family activation, route-target policy, capabilities |

Useful FRRouting commands include:

```text
show bgp summary
show bgp neighbors 10.0.0.1
show bgp ipv4 unicast
show bgp ipv4 unicast 203.0.113.0/24
show ip route 10.0.0.1
show ip route 192.0.2.1
```

## Design Checklist

- [ ] Router loopbacks are unique and advertised by the IGP.
- [ ] Every BGP next hop is resolvable.
- [ ] The topology uses a full mesh, route reflectors, or confederations intentionally.
- [ ] At least two route reflectors exist where availability requires them.
- [ ] Reflector clusters and client roles are documented.
- [ ] `next-hop-self` behavior is deliberate at edge routers and reflectors.
- [ ] Import and export policy is consistent across equivalent peers.
- [ ] Failure scenarios have been tested without redistributing BGP into the IGP.

## Further Reading

- [BGP fundamentals](BGP.md)
- [eBGP](eBGP.md)
- [IGP](IGP.md)
- [RFC 4456 — BGP Route Reflection](https://www.rfc-editor.org/rfc/rfc4456)
- [RFC 5065 — Autonomous System Confederations for BGP](https://www.rfc-editor.org/rfc/rfc5065)
- [RFC 7911 — Advertisement of Multiple Paths in BGP](https://www.rfc-editor.org/rfc/rfc7911)
