# Roles and permissions
Groups, users and permission rules provide the central core organization principle in *mercurius*. The system has the 
following functional requirements:
* The system must support an arbitrary amount of groups. 
* A group must be able to contain other groups and users.
* The system has a global set of permission rules.
* Rules must be able to be configured and instantiated with different parameters.
* Rules must be inherited by all children (groups and users).
* If no other rules forbids it, any rule must be addable to enforce more limitations.
* If no other rules forbids it, any rule must be removable to enforce less limitations.
* Groups represent different user role domains, like an uber-admin > tenant-owner > tenant > customer-organization > 
customer-client > end-user-group > end-user and so forth. Each level is subject to contain user entities.

## Tenant
An installation must organize multiple tenants. Each tenant has its own. For this reason, we introduce the wording *tenant* for the first level user role.