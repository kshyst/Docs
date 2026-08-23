# External BGP (eBGP)

**External BGP (eBGP)** is BGP peering between routers in different autonomous systems. It is used for ISP transit, private network interconnections, Internet exchanges, cloud connectivity, and customer-provider routing.

```text
       AS 65001                             AS 65002

  203.0.113.0/24 -- R1 192.0.2.1 ===== 192.0.2.2 R2 -- 198.51.100.0/24
                              eBGP
```

## Defining Behavior

Compared with iBGP, eBGP normally has these properties:

- the peers use different ASNs;
- the advertising router prepends its ASN to `AS_PATH`;
- a router rejects a route when its own ASN already appears in `AS_PATH`, unless explicitly configured otherwise;
- many implementations change `NEXT_HOP` to the advertising router;
- directly connected peering is a common default expectation;
- an eBGP-learned route is normally eligible for advertisement to other eBGP and iBGP peers, subject to policy.

These are common defaults rather than a substitute for platform documentation.

## AS_PATH and Loop Prevention

Suppose a prefix originates in AS 65003 and crosses AS 65002 before reaching AS 65001:

```text
Prefix: 198.51.100.0/24
AS_PATH seen by AS 65001: 65002 65003
```

If AS 65002 later receives an advertisement whose path contains `65002`, it rejects that path. This prevents inter-AS routing loops.

Operators may prepend their own ASN multiple times to make a path less attractive:

```text
65002 65002 65002 65003
```

Prepending is only a hint. A remote network may prioritize `LOCAL_PREF` or another policy over path length.

## Direct and Multihop Peering

### Directly Connected eBGP

Direct peering uses addresses on a shared link. It is simple, makes next-hop resolution obvious, and supports failure detection tied to the interface.

```text
router bgp 65001
 neighbor 192.0.2.2 remote-as 65002
```

Many implementations send directly connected eBGP traffic with IP TTL 1 by default.

### eBGP Multihop

Multihop peering establishes the session between addresses that are not directly connected, often loopbacks. It requires explicit multihop configuration and a route to the peer address.

```text
router bgp 65001
 neighbor 198.51.100.2 remote-as 65002
 neighbor 198.51.100.2 update-source lo
 neighbor 198.51.100.2 ebgp-multihop 2
```

Multihop can improve session resilience across parallel links, but it also introduces recursive routing and security considerations. Use the smallest practical TTL and strict filtering.

## Policy at the AS Boundary

Every eBGP peer should have an explicit import and export policy.

### Import Checklist

- Permit only prefixes the peer is authorized to send.
- Enforce minimum and maximum prefix lengths.
- Apply a maximum-prefix limit.
- Reject default routes unless they are expected.
- Reject special-use and inappropriate private prefixes.
- Validate route origins with RPKI when applicable.
- Set local preference and communities according to business intent.

### Export Checklist

- Permit only locally authorized and customer prefixes.
- Prevent accidental transit between providers or peers.
- Do not export internal infrastructure routes.
- Aggregate where operationally appropriate.
- Attach documented communities.
- Remove private ASNs only when the design requires it.

!!! danger
    An established eBGP session is not evidence that its routes are safe. Missing or incorrect export policy can leak a full routing table and cause an outage far beyond the local network.

## Relationship-Based Policy

Internet policy often reflects commercial relationships:

| Relationship | Typical routes accepted | Typical routes exported |
| --- | --- | --- |
| Customer | Customer and customer-downstream routes | Local, customer, peer, and provider routes |
| Peer | Peer and peer-customer routes | Local and customer routes |
| Provider | Broad Internet reachability | Local and customer routes |

This pattern helps prevent providing unpaid transit. Actual contracts and routing policy always take precedence.

## Traffic Engineering

### Influencing Outbound Traffic

The local AS controls outbound path selection directly. Common tools include:

- `LOCAL_PREF`;
- local weight on platforms that support it;
- selective route acceptance;
- IGP cost to the BGP next hop;
- equal-cost multipath where supported.

### Influencing Inbound Traffic

A remote AS ultimately controls its own outbound choice. The local AS can only influence it with signals such as:

- advertising more-specific prefixes;
- AS-path prepending;
- `MED` when the neighbor honors it;
- provider-defined BGP communities;
- selective advertisement across links.

More-specific routes are powerful because forwarding uses longest-prefix match. Use them carefully and ensure they are accepted by upstream prefix-length policy.

## Complete FRRouting Example

```text
router bgp 65001
 bgp router-id 192.0.2.1
 neighbor 192.0.2.2 remote-as 65002
 neighbor 192.0.2.2 description TRANSIT-AS65002
 neighbor 192.0.2.2 password example-secret
 !
 address-family ipv4 unicast
  network 203.0.113.0/24
  neighbor 192.0.2.2 activate
  neighbor 192.0.2.2 prefix-list TRANSIT-IN in
  neighbor 192.0.2.2 prefix-list TRANSIT-OUT out
  neighbor 192.0.2.2 maximum-prefix 1000
 exit-address-family

ip prefix-list TRANSIT-IN seq 10 permit 0.0.0.0/0
ip prefix-list TRANSIT-IN seq 100 deny any
ip prefix-list TRANSIT-OUT seq 10 permit 203.0.113.0/24
ip prefix-list TRANSIT-OUT seq 100 deny any
```

This intentionally accepts only a default route and exports one documentation prefix. The password is illustrative, not a production secret.

## Private ASNs

Private ASNs are not intended to appear in the public Internet routing table. A provider may remove a customer's private ASN before exporting routes globally:

```text
neighbor 192.0.2.2 remove-private-AS
```

Behavior differs when a path mixes public and private ASNs or contains repeated ASNs. Verify the platform's exact semantics.

## Fast Failure Detection

Common options include:

- short BGP keepalive and hold timers;
- Bidirectional Forwarding Detection (BFD);
- interface-state tracking for direct peers;
- redundant physical links and sessions.

Aggressive timers can create instability during transient congestion or control-plane load. Choose timers based on measured requirements rather than using the lowest possible values.

## Troubleshooting eBGP

| Symptom | Likely checks |
| --- | --- |
| Session remains `Active` | IP reachability, TCP/179, ACLs, peer ASN, source address, TTL |
| `OPEN` error or immediate reset | ASN mismatch, unsupported capability, authentication |
| Session established but receives zero routes | Address-family activation, peer export policy, local import policy |
| Route received but not installed | Next-hop reachability, best path, RIB preference, RPKI state |
| Prefix not advertised | Prefix absent from local RIB, export filter, aggregation, route validity |
| Intermittent resets | Link loss, BFD, timer expiry, maximum-prefix, CPU or memory pressure |

Useful FRRouting commands:

```text
show bgp ipv4 unicast summary
show bgp neighbors 192.0.2.2
show bgp ipv4 unicast neighbors 192.0.2.2 routes
show bgp ipv4 unicast neighbors 192.0.2.2 advertised-routes
show ip route 192.0.2.2
```

## Security Recommendations

- Use explicit import and export filters.
- Configure maximum-prefix thresholds.
- Protect TCP sessions with authentication where supported.
- Use GTSM for suitable direct peerings.
- Apply control-plane ACLs that permit BGP only from known peers.
- Monitor RPKI validation and decide how to treat invalid routes.
- Alert on session resets, unexpected origin AS changes, and route-count spikes.

## Further Reading

- [BGP fundamentals](BGP.md)
- [iBGP](iBGP.md)
- [RFC 7454 — BGP Operations and Security](https://www.rfc-editor.org/rfc/rfc7454)
- [RFC 5082 — Generalized TTL Security Mechanism](https://www.rfc-editor.org/rfc/rfc5082)
- [RFC 8212 — Default External BGP Route Propagation Behavior](https://www.rfc-editor.org/rfc/rfc8212)
