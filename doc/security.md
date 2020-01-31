# Security

The following design decisions are respected in the implementation of the service.

## Insecure Direct Object Reference 

The so called IDOR attack is possible if an application exposes storage internal identifiers to the world, like
auto incrementing user or session ids from a database. An attacker can simply guess valid identifiers from it. See
also [[OWASP]](https://owasp.org/www-project-cheat-sheets/cheatsheets/Insecure_Direct_Object_Reference_Prevention_Cheat_Sheet.html)
and [[bugcrowd]](https://www.bugcrowd.com/blog/how-to-find-idor-insecure-direct-object-reference-vulnerabilities-for-large-bounty-rewards/).

However we won't apply the according recommendation, because the additional layer of indirection - especially demonstrated
by the OWASP example itself - negates any index properties from the storage layer, which is in any but the simplest
scenarios unacceptable. Thinking further, to optimize this, we need to track another map, to efficiently do reverse 
lookups, which must also survive server restarts (so be persistent). Also in most cases it is not feasible to resolve
to session local ids, because otherwise Apps will break on session timeouts entirely. Therefore we decide to expose
the internal ids, however we use random ones. For this we follow the [proposal](https://neilmadden.blog/2018/08/30/moving-away-from-uuids/) 
of Neil Madden, the author of the book *API Security in Action*, to avoid inefficiencies and broken entropy of UUIDs.
For sure, each resource access will be checked for proper authorization.