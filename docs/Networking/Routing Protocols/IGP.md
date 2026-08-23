# Interior Gateway Protocols (IGPs)

An **Interior Gateway Protocol (IGP)** exchanges routing information within one administrative domain or autonomous system. Its main job is to provide fast, dependable reachability between internal routers, links, loopbacks, and subnets.

Common IGPs include:

- **OSPF** — a link-state protocol widely used in enterprise and service-provider networks;
- **IS-IS** — a link-state protocol common in large service-provider networks;
- **RIP** — a distance-vector protocol mainly used in small or legacy environments;
- **EIGRP** — an advanced distance-vector protocol used primarily in Cisco-oriented networks.

An IGP is a protocol category, not one specific protocol.

## IGP and BGP Roles

| IGP | BGP |
| --- | --- |
| Optimizes internal reachability | Applies routing policy within or between autonomous systems |
| Usually converges quickly after local failures | Scales to very large route sets and policy boundaries |
| Uses topology or distance information | Uses path attributes such as AS_PATH and LOCAL_PREF |
| Commonly carries links and loopbacks | Commonly carries Internet, customer, VPN, or service routes |
| Operates within one administrative domain | eBGP crosses AS boundaries; iBGP operates inside an AS |

A common service-provider design uses OSPF or IS-IS to carry router loopbacks and point-to-point links, then establishes iBGP between those loopbacks.

!!! warning
    Do not redistribute a full Internet BGP table into an IGP. IGPs are not designed for that route scale or policy model, and the resulting churn can destabilize the network.

## Main Algorithm Families

### Link-State Protocols

OSPF and IS-IS flood information describing links and their state. Each router builds a link-state database and independently runs a shortest-path-first calculation.

Simplified process:

1. Discover neighbors with hello messages.
2. Form adjacencies with eligible neighbors.
3. Flood link-state information.
4. Build a synchronized topology database.
5. Run Dijkstra's shortest-path-first algorithm.
6. Install selected routes in the routing table.

**Advantages:** fast convergence, strong topology visibility, hierarchy, and good scalability.

**Costs:** more concepts, database synchronization, and careful area or level design.

### Distance-Vector Protocols

Distance-vector protocols advertise destinations and a distance or metric to neighbors rather than a complete topology map.

RIP uses hop count. EIGRP uses the Diffusing Update Algorithm and a composite metric.

**Advantages:** potentially simpler configuration and lower conceptual overhead in small networks.

**Costs:** protocol-dependent scaling limits and less complete topology visibility.

## Protocol Comparison

| Property | OSPF | IS-IS | RIP | EIGRP |
| --- | --- | --- | --- | --- |
| Family | Link state | Link state | Distance vector | Advanced distance vector |
| Typical metric | Cost | Configured interface metric | Hop count | Composite metric |
| Hierarchy | Areas, with Area 0 backbone | Level 1 and Level 2 | None | Stub and summarization features |
| Equal-cost multipath | Yes | Yes | Yes, implementation-dependent count | Yes; unequal-cost also supported |
| Primary use | Enterprise and provider networks | Provider and large routed networks | Small or legacy networks | Cisco-centric enterprise networks |
| IP transport | IP protocol 89 | Runs directly over Layer 2 | UDP 520 | IP protocol 88 |

## Metrics and Administrative Preference

An IGP metric compares paths learned by the **same protocol**. Examples include OSPF cost, IS-IS metric, and RIP hop count.

When routes to the same prefix come from different sources, a router commonly uses a separate local preference called **administrative distance**, **route preference**, or **protocol preference**. Its name and default values depend on the vendor.

Do not confuse:

- **longest-prefix match**, which chooses the most specific destination during forwarding;
- **administrative preference**, which chooses between route sources for the same prefix length;
- **protocol metric**, which chooses between paths within a routing protocol.

## OSPF Overview

OSPF organizes routers and links into areas. Area `0.0.0.0` is the backbone for a multi-area OSPF domain.

Common concepts include:

- router ID;
- neighbors and adjacencies;
- designated and backup designated routers on multi-access networks;
- link-state advertisements (LSAs);
- area border routers (ABRs);
- autonomous system boundary routers (ASBRs);
- internal, inter-area, and external routes.

Minimal FRRouting example:

```text
interface eth0
 ip ospf area 0.0.0.0
 ip ospf network point-to-point

interface lo
 ip ospf area 0.0.0.0

router ospf
 ospf router-id 10.0.0.1
 passive-interface lo
```

A passive interface advertises its connected network without attempting to form neighbors on that interface.

## IS-IS Overview

IS-IS uses Level 1 for routing within an area and Level 2 for routing between areas. A router may operate as Level 1, Level 2, or Level 1-2.

Common concepts include:

- NET and NSAP addressing;
- intermediate-system neighbors;
- link-state protocol data units (LSPs);
- designated intermediate systems on broadcast networks;
- Level 1 and Level 2 databases;
- wide metrics;
- type-length-value (TLV) extensions.

Because IS-IS runs directly over Layer 2 rather than over IP, an incorrect IP configuration does not necessarily prevent an adjacency. It can still prevent IP routes and forwarding from working correctly.

Illustrative FRRouting example:

```text
router isis CORE
 net 49.0001.0100.0000.0001.00
 is-type level-2-only
 metric-style wide

interface eth0
 ip router isis CORE
 isis network point-to-point

interface lo
 ip router isis CORE
 isis passive
```

## Hierarchical Design

Hierarchy limits topology information and failure impact.

### OSPF Areas

- Keep Area 0 contiguous.
- Connect non-backbone areas through Area 0 under normal design.
- Summarize at ABRs where addressing allows it.
- Use stub area types only when their route restrictions match requirements.

### IS-IS Levels

- Level 1 routers know their local area and use the nearest Level 1-2 router for external destinations.
- Level 2 provides the inter-area backbone.
- Level 1-2 routers connect both scopes.

Start with a single area or level when scale does not require hierarchy. Premature partitioning adds complexity without benefit.

## Convergence

**Convergence** is the process of reaching a consistent routing state after a change. It includes:

1. detecting the failure;
2. informing other routers;
3. recalculating paths;
4. updating the routing and forwarding tables.

Ways to improve convergence include:

- point-to-point network types on point-to-point links;
- BFD for rapid failure detection;
- tuned hello and dead timers;
- fast reroute features;
- stable route summarization;
- avoiding excessive route redistribution;
- sufficient control-plane CPU and memory.

Faster is not always safer. Overly aggressive timers can cause false failures during congestion or control-plane load.

## Route Redistribution

Redistribution imports routes from another protocol or source. It creates a boundary where loops, feedback, and inconsistent metrics can occur.

Use these safeguards:

- redistribute at as few points as possible;
- filter exact prefixes;
- tag routes and reject them on re-entry;
- choose metrics explicitly;
- summarize where safe;
- document the source of default routes;
- prefer BGP for complex policy between routing domains.

!!! danger
    Mutual redistribution at multiple routers can create persistent routing loops even when every individual protocol is loop-free.

## Authentication and Security

IGPs normally trust routers that can form an adjacency, so protect the routing domain:

- authenticate protocol messages where supported;
- make user-facing interfaces passive;
- filter routing protocol packets at network boundaries;
- use infrastructure ACLs and control-plane policing;
- log adjacency changes;
- protect configuration access;
- avoid placing untrusted devices on routing links.

Authentication verifies participation; it does not make incorrect metrics or malicious route advertisements safe.

## Troubleshooting Workflow

1. **Check interfaces:** Are the link and line protocol up, with correct addressing and MTU?
2. **Check neighbor discovery:** Do hello/dead timers, area or level, authentication, and network type match?
3. **Check adjacency state:** Identify where database exchange stops.
4. **Check the protocol database:** Does the expected topology or prefix exist?
5. **Check route calculation:** Compare metrics, area restrictions, and route preference.
6. **Check the routing table:** Was another route source selected?
7. **Check the forwarding table and data plane:** Test the actual path and return path.
8. **Check redistribution and filters:** Look for tags, route maps, summaries, and default-route conditions.

Common FRRouting commands:

```text
show ip ospf neighbor
show ip ospf interface
show ip ospf database
show ip route ospf
show isis neighbor
show isis database
show ip route isis
show bfd peers
```

## Practical Design Checklist

- [ ] Router IDs and loopback addresses are unique and stable.
- [ ] Point-to-point links use the appropriate network type.
- [ ] User-facing interfaces are passive.
- [ ] Metrics reflect intended primary and backup paths.
- [ ] Areas or levels are no more complex than scale requires.
- [ ] Redistribution points are few, filtered, tagged, and documented.
- [ ] Every iBGP next hop is reachable through the IGP.
- [ ] Authentication and control-plane filtering are enabled where supported.
- [ ] Convergence and failure behavior have been tested.

## Further Reading

- [BGP](BGP.md)
- [iBGP](iBGP.md)
- [RFC 2328 — OSPF Version 2](https://www.rfc-editor.org/rfc/rfc2328)
- [RFC 5340 — OSPF for IPv6](https://www.rfc-editor.org/rfc/rfc5340)
- [RFC 1195 — Use of IS-IS for Routing in TCP/IP](https://www.rfc-editor.org/rfc/rfc1195)
- [RFC 2453 — RIP Version 2](https://www.rfc-editor.org/rfc/rfc2453)
