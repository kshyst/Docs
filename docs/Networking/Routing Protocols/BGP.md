# Border Gateway Protocol (BGP)

**Border Gateway Protocol (BGP)** is the Internet's interdomain routing protocol. It exchanges IP-prefix reachability together with attributes that describe each path. Operators use those attributes and local policy to decide which routes to accept, prefer, and advertise.

BGP is a **path-vector** protocol. Unlike an IGP, it does not build a complete map of the network or select a path from link bandwidth alone. Its central concerns are **policy**, **scalability**, and **loop prevention between autonomous systems**.

## Core Terms

| Term | Meaning |
| --- | --- |
| Autonomous system (AS) | A network under a common routing policy, identified by an ASN |
| ASN | A 16-bit or 32-bit autonomous system number |
| Peer or neighbor | A router with which a BGP session is configured |
| NLRI | Network Layer Reachability Information, such as an IPv4 or IPv6 prefix |
| Path attribute | Metadata attached to a route, such as AS_PATH or LOCAL_PREF |
| RIB | Routing Information Base containing candidate or selected routes |
| Policy | Rules that accept, reject, modify, or prioritize routes |

Private ASNs are useful inside labs or private interconnections. Common private ranges are `64512–65534` and `4200000000–4294967294`. Public Internet advertisements require an ASN allocated through the appropriate registry process.

## How a BGP Session Works

BGP peers establish a TCP connection, normally to destination port **179**. TCP provides reliable, ordered delivery; BGP itself handles routing state and updates.

A simplified session sequence is:

1. The peers establish TCP connectivity.
2. They exchange `OPEN` messages containing ASN, hold time, BGP identifier, and capabilities.
3. They confirm the session with `KEEPALIVE` messages.
4. They exchange reachable and withdrawn prefixes in `UPDATE` messages.
5. They send periodic keepalives and notifications when errors occur.

The principal BGP finite-state-machine states are `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`, and `Established`. Routes are exchanged only in `Established`.

!!! tip
    An `Active` session is not healthy or established. It normally means the router is retrying the TCP connection. Check addressing, routing, TCP port 179, peer ASN, source address, and authentication.

## Important Path Attributes

| Attribute | Purpose |
| --- | --- |
| `AS_PATH` | Lists ASes traversed; provides loop detection and influences path selection |
| `NEXT_HOP` | Address through which the advertised prefix is reachable |
| `LOCAL_PREF` | Selects the preferred outbound path inside an AS; higher is commonly preferred |
| `MED` | Suggests a preferred entry point to a neighboring AS; lower is commonly preferred |
| `ORIGIN` | Indicates how the route entered BGP: IGP, EGP, or incomplete |
| `COMMUNITY` | Tags routes so policy can classify and act on groups of prefixes |
| `ATOMIC_AGGREGATE` / `AGGREGATOR` | Records information related to route aggregation |

`LOCAL_PREF` is normally used throughout one AS and is not sent to eBGP peers. `MED` is generally sent to external neighbors, but whether and how it is compared depends on implementation and policy.

## Best-Path Selection

There is no single universal best-path sequence shared by every implementation. A typical router evaluates only valid routes and then considers factors such as:

1. locally configured preference or weight;
2. highest `LOCAL_PREF`;
3. locally originated routes;
4. shortest `AS_PATH`;
5. `ORIGIN` type;
6. lowest applicable `MED`;
7. eBGP-learned versus iBGP-learned paths;
8. IGP cost to the BGP next hop;
9. implementation-specific tie-breakers.

Always consult the router vendor's documentation before depending on an exact selection order.

!!! warning
    A BGP route can be the protocol's best path yet remain unusable if its `NEXT_HOP` is not present in the routing table.

## BGP Is Policy-Driven

BGP configuration should begin with policy, not merely neighbor establishment.

**Import policy** controls what the local AS accepts:

- permit only expected prefixes;
- enforce sensible prefix-length limits;
- reject the local AS in received paths where appropriate;
- set local preference or communities;
- cap the number of received prefixes.

**Export policy** controls what the local AS announces:

- advertise only authorized prefixes;
- avoid leaking routes learned from one provider to another;
- remove private ASNs where required;
- attach agreed communities;
- originate a prefix only when intended.

Modern operational guidance favors **default reject** for eBGP when no explicit policy exists.

## Address Families and MP-BGP

Multiprotocol BGP (MP-BGP) carries multiple address families over BGP, including:

- IPv4 unicast;
- IPv6 unicast;
- VPNv4 and VPNv6;
- Ethernet VPN (EVPN);
- multicast and labeled-unicast routes.

A neighbor relationship and an address-family activation are separate concepts on many platforms. A TCP session can therefore be established while no routes are exchanged for a particular address family.

## Minimal FRRouting Example

The following example originates `203.0.113.0/24` and peers with AS 65002. Documentation prefixes are used intentionally.

```text
router bgp 65001
 bgp router-id 192.0.2.1
 neighbor 192.0.2.2 remote-as 65002
 !
 address-family ipv4 unicast
  network 203.0.113.0/24
  neighbor 192.0.2.2 activate
  neighbor 192.0.2.2 prefix-list FROM-AS65002 in
  neighbor 192.0.2.2 prefix-list TO-AS65002 out
 exit-address-family

ip prefix-list FROM-AS65002 seq 10 permit 198.51.100.0/24
ip prefix-list FROM-AS65002 seq 100 deny any
ip prefix-list TO-AS65002 seq 10 permit 203.0.113.0/24
ip prefix-list TO-AS65002 seq 100 deny any
```

The `network` statement normally requires a matching route in the local routing table. Exact syntax and behavior vary by software release.

## Operational Security

A production BGP design should consider:

- strict prefix and AS-path filters;
- maximum-prefix limits with an intentional recovery policy;
- TCP authentication where supported;
- Generalized TTL Security Mechanism (GTSM) for directly connected peers;
- RPKI Route Origin Validation for Internet routes;
- Bogon and special-use prefix filtering;
- peer and route-change monitoring;
- redundant sessions and controlled maintenance procedures.

RPKI origin validation checks whether an AS is authorized to originate a prefix. It does **not** validate the complete AS path.

## Troubleshooting Workflow

1. **Check IP reachability:** Can the configured source address reach the peer?
2. **Check TCP:** Is port 179 permitted, and which side is initiating or accepting the connection?
3. **Check neighbor parameters:** Verify peer address, local and remote ASN, source interface, password, and address family.
4. **Check session state:** Inspect the last reset reason and notification messages.
5. **Check received routes:** Determine whether updates arrived before import policy.
6. **Check policy:** Confirm that prefix lists, route maps, communities, and maximum-prefix controls allow the route.
7. **Check next-hop resolution:** Verify a route to `NEXT_HOP` through the IGP or a static route.
8. **Check best-path output:** Compare path attributes and implementation-specific tie-breakers.
9. **Check advertisement:** Confirm the route is eligible and export policy permits it.

Common FRRouting commands include:

```text
show bgp summary
show bgp neighbors 192.0.2.2
show bgp ipv4 unicast
show bgp ipv4 unicast 203.0.113.0/24
show bgp ipv4 unicast neighbors 192.0.2.2 advertised-routes
show ip route 192.0.2.2
```

## BGP Variants

- [eBGP](eBGP.md) connects routers in different autonomous systems.
- [iBGP](iBGP.md) distributes BGP information inside one autonomous system.
- [IGPs](IGP.md) provide the internal reachability on which many BGP designs depend.

## Further Reading

- [RFC 4271 — BGP-4](https://www.rfc-editor.org/rfc/rfc4271)
- [RFC 4760 — Multiprotocol Extensions for BGP-4](https://www.rfc-editor.org/rfc/rfc4760)
- [RFC 7454 — BGP Operations and Security](https://www.rfc-editor.org/rfc/rfc7454)
- [RFC 8212 — Default External BGP Route Propagation Behavior](https://www.rfc-editor.org/rfc/rfc8212)
