# Architecture
<!-- blank -->
**EVPN/VXLAN Fabric with Nexus 9000v**
<!-- blank -->
<pre>
1. eBGP for both the 'underlay' and 'overlays'.
2. 2 x spines and '8' leaf nodes.
3. 4 x leaf nodes for compute.
4. 2 x leaf nodes for services.
5. 2 x leaf nodes for borders.
6. Leaf nodes will be in vPC pairs.
</pre>
<!-- blank -->
<!-- blank -->
## Purpose
<!-- blank -->
**Test bed for network automation using Ansible, Jinja2, and Terraform.**
<!-- blank -->
<!-- blank -->
## Nodes Abbreviations and Usage
<!-- blank -->
|**ID**   | **Definition**           |
|:--------|:-------------------------|
| CL	  | Compute Leaf             |
| SL	  | Service Leaf             |
| BL	  | Border Leaf              |
| SP      | Spine                    |
| LF	  | Leaf                     |
| LP	  | Leaf Pair                |
| CLP     | Compute Leaf Pair        |
| SLP     | Service Leaf Pair        |
| BLP     | Border Leaf Pair         |
| ::      | Border Leaf Pair         |
| \#      | single digit placeholder |
| \##     | double digit placeholder |
| nxx     | node name\+node instance |
| Int.    | interface                |
<!-- blank -->
<!-- blank -->
## Interface Abbreviations
<!-- blank -->
|**Int abbr.**| **Definition**                |
|:------------|:------------------------------|
| e1.1	      | Ethernet1/1                   |
| lf01.e1.1   | leaf switch 01 Ethernet1/1    |
<!-- blank -->
<!-- blank -->
## Fabric Naming Convention
<!-- blank -->
|**Site**|**Pod** |**nxx** |
|:------:|:------:|:------:|
| s1	 | p1     | sp01   |
<!-- blank -->
<pre>
Hostname: s1-p1-sp01
	
Note: names are all lowercase
</pre>
<!-- blank -->
<!-- blank -->
## IPv4
<!-- blank -->
|**Description**    | **CIDR**       | **Comment ** |
|:------------------|:---------------|:-------------|
| Underlay	        | 10.1.1.0/26    | eBGP         |
| vPC Keepalive     | 10.1.1.64/27   |              |
| Loopback0	        | not used       |              |
| Loopback1	        | 10.0.1.0/27    | VTEP         |
<!-- blank -->
<pre>
Note: loopback0 will not be used due to a bug with N9Kv image and vPC.
</pre>
<!-- blank -->
<!-- blank -->
## VRF
<!-- blank -->
|**Description**    | **VRF Name**   | **Comment ** |
|:------------------|:---------------|:-------------|
| Underlay	        | Global         | default      |
| vPC Keepalive     | Keepalive	     |              |
| Tenant A          | Red            | Customer 1   |
| Tenant B          | Blue           | Customer 2   |
| Tenant C          | Green          | Customer 3   |
| ...               | ...            | ...          |
<!-- blank -->
<!-- blank -->
## ASN's
<!-- blank -->
|**AS#**      | **Designation**               |
|:------------|:------------------------------|
| 65500	      | Spines                        |
| 65501       | Compute leaf pair 1           |
| 65502       | Compute leaf pair 2           |
| 65503       | Service leaf pair 1           |
| 65504       | Border leaf pair 1            |
<!-- blank -->
<!-- blank -->
## Spines
<!-- blank -->
<pre>
Spines 'only' connect to leaf nodes (E.g., no inter-switch links)
</pre>
<!-- blank -->
<!-- blank -->
## vPC Peering
<!-- blank -->
|**vPC Peer A**|**vPC Peer B**|**vPC Domain**|**Leaf Pair**|
|:-------------|:-------------|:-------------|:------------|
| s1-p1-lf01   | s1-p1-lf02   | 1            | CLP01       |
| s1-p1-lf03   | s1-p1-lf04   | 1            | CLP02       |
| s1-p1-sl01   | s1-p1-sl02   | 1            | SLP01       |
| s1-p1-bl01   | s1-p1-bl02   | 1            | BLP01       |
<!-- blank -->
<!-- blank -->
|**Peer Function**|**Peer Interface**|**PO No.**|**PO Mode**|
|:----------------|:-----------------|:--------:|:---------:|
| keepalive       | lf01.e1.8        | 10       | L3        |
| keepalive       | lf01.e1.9        | 10       | L3        |
| peer-link       | lf01.e1.10       | 20       | L2 Trunk  |
| peer-link       | lf01.e1.11       | 20       | L2 Trunk  |
<!-- blank -->
<!-- blank -->
|**vPC LP** |**Keepalive Int**|**Peer A IPv4**|**Peer B IPv4**|**Mask**|
|:----------|:----------------|:--------------|:--------------|:-------|
| CLP01     | port-channel10  | 10.1.1.64     | 10.1.1.65     | /31    |
| CLP02     | port-channel10  | 10.1.1.66     | 10.1.1.67     | /31    |
| SLP01     | port-channel10  | 10.1.1.68     | 10.1.1.69     | /31    |
| CLP01     | port-channel10  | 10.1.1.70     | 10.1.1.71     | /31    |
<!-- blank -->
<!-- blank -->
## Fabric uplink details:
<!-- blank -->
|**Spine.Int**|**Leaf.Int**|    |**SP IPv4**     |**LF IPv4**     |**Mask**|
|:------------|:-----------|:--:|:---------------|:---------------|:------:|
|sp01.e1.1    |lf01.e1.1   |    | 10.1.1.0       | 10.1.1.1       | /31    |
|sp02.e1.1    |lf01.e1.2   |    | 10.1.1.20      | 10.1.1.21      | /31    |
|sp01.e1.2    |lf02.e1.1   |    | 10.1.1.2       | 10.1.1.3       | /31    |
|sp02.e1.2    |lf02.e1.2   |    | 10.1.1.22      | 10.1.1.23      | /31    |
|sp01.e1.3    |lf03.e1.1   |    | 10.1.1.4       | 10.1.1.5       | /31    |
|sp02.e1.3    |lf03.e1.2   |    | 10.1.1.24      | 10.1.1.25      | /31    |
|sp01.e1.4    |lf04.e1.1   |    | 10.1.1.6       | 10.1.1.7       | /31    |
|sp02.e1.4    |lf04.e1.2   |    | 10.1.1.26      | 10.1.1.27      | /31    |
|sp01.e1.5    |sl01.e1.1   |    | 10.1.1.8       | 10.1.1.9       | /31    |
|sp02.e1.5    |sl01.e1.2   |    | 10.1.1.28      | 10.1.1.29      | /31    |
|sp01.e1.6    |sl02.e1.1   |    | 10.1.1.10      | 10.1.1.11      | /31    |
|sp02.e1.6    |sl02.e1.2   |    | 10.1.1.30      | 10.1.1.31      | /31    |
|sp01.e1.7    |bl01.e1.1   |    | 10.1.1.12      | 10.1.1.13      | /31    |
|sp02.e1.7    |bl01.e1.2   |    | 10.1.1.32      | 10.1.1.33      | /31    |
|sp01.e1.8    |bl02.e1.1   |    | 10.1.1.14      | 10.1.1.15      | /31    |
|sp02.e1.8    |bl02.e1.2   |    | 10.1.1.34      | 10.1.1.35      | /31    |


