# About IP and Some definitions

### Definitions

**CNAME (Canonical Name)**: An alias for other domain names.

**IP address**: Identifies a device

**Subnet Mask**: Defines the network it belongs to 
> A 32-bit number that distinguished the network part of an IP address from host part. It helps local network to identify each other and determine which devices are on the same subnet.

**Default Gateway**: The route to other networks

> It is the IP address of a router or other device that acts as path to other networks or the internet. When a device wants to communicate with a device on a different network, it sends the data to the default gateway which then forwards the data to the correct destination.

**CIDR**: Classless Inter-Domain Routing
> Before CIDR networks where categorized into class A,B and C.
> 
> CIDR uses VLSM (Variable-Length Subnet Masking) allowing networks to be divided into subnets of varying sizes.
> 
> If subnet mask is 255.255.255.0 the CIDR notation would be /24 
> 
> To get the network portion of an IP address we bitwise AND the IP address and the subnet mask or simply if the CIDR is for example /25 we choose only the 25 first bits of the IP address and that would be our network portion.
> 
> Subnet mask is just some consecutive 1's followed by consecutive 0's and 32 bits.

**MAC Address**: Media Access Control
> 
