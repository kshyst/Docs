# Routing Protocols

- **OSPF (Open Shortest Path First)**: Every router on network share information about  the network topology and the most efficient path for data transmissions. In this protocol, routers exchange updates about the state of their connected links and networks. This way each router has a complete map of the network and can determine the best routes.
- **EIGRP (Enhanced Interior Gateway Routing Protocol)**: Built by cisco, it combines aspects of different routing algorithm. Every router shares what network they can reach and the cost(like bandwidth or delay) associated with those routes.
- **BGP (Border Gateway Protocol)**: Primary routing protocol used in internet. It allows different networks(Like ISPs) to exchange routing information and establish paths for data to travel between these networks. BGP helps ensure data can route efficiently across multiple networks.
- **RIP (Routing Information Protocol)**: Simple routing protocol in small networks. Routers using RIP share information about the networks they can reach and number of hops(routers) required to get there. As a result each router builds a routing table based on this information and then chooses the route with fewest hops to reach dest.

