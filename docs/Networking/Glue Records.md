# Glue Records

For example we want to resolve this with the given dns:
```
example.com
Nameserver: ns1.example.com
```

To resolve example.com, a resolver must ask ns1.example.com.

But to contact ns1.example.com, it must first resolve ns1.example.com.

## What is Glue Record

A glue record is an IP address (A or AAAA record) for the nameserver that is stored at the parent zone, not in the child zone.

So the parent zone (.com) provides:

```text
example.com      NS   ns1.example.com
ns1.example.com  A    192.0.2.10   ← glue record
```

This allows resolvers to reach the nameserver without needing to resolve the domain first.

## Where glue records live

- Stored in the parent zone
- Managed by the domain registrar / registry
- Automatically included when required

## Example

In real world
You define nameservers
If nameserver is under your domain:
- Registrar asks for IP address
- That IP becomes a glue record
