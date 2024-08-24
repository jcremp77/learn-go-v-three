# NX-OS EVPN VXLAN Fabric Architecture

Covers the deployment of a spine/leaf fabric using eBGP as the 'underlay' and the 'overlays'. 

## Abreviations

### Nodes

CL	= Compute Leaf
SL	= Service Leaf
BL	= Border Leaf
SP	= Spine
LF	= Leaf
LFP	= Leaf Pair
CLP	= Compute Leaf Pair
SLP	= Service Leaf Pair
BLP	= Border Leaf Pair
:: = denotes physical link

### Interfaces

>e1.1 		= Ethernet1/1
>
>lf01.e1.1	= Leaf switch 01 Ethernet1/1

## Fabric Naming Convention

>Single # 	= single digit
>
>Double ## 	= double digit

\<site#\> \<pod#\> \<func + instance##\>; all names are lowercase.

Example: 's1\-p1\-sp01' is interpreted as	'site1'\-'pod1'\-'spine01'

Example: 's1\-p1\-lf01' is interpreted as 'site1'\-'pod1'\-'leaf01'

## Architecture Summary

- Clos fabric
- eBGP for both the 'underlay' and 'overlays'.
- 2 x spines and '8' leaf nodes.
- 4 x leaf nodes for compute.
- 2 x leaf nodes for services.
- 2 x leaf nodes for borders.
- All leaf nodes will be in vPC pairs.

## Fabric Routing

### IPv4

1. eBGP Underlay:	10.1.1.0/26
2. VPC Keepalive: 	10.1.1.64/27
3. Loop0:			N/A (unassigned due to issues with Nexus 9Kv)
4. Loop1 (eVPN):	10.0.1.0/27
5. Loop2 (overlay):	10.0.1.32/27
6. Loop3 (overlay): 10.0.1.64/27
7. Loop4 (overlay): 10.0.1.96/27

### VRF

1. Underlay: 		Global (default)
2. Customer 1:		Red
3. Customer 2:		Blue
4. Customer 3:		Green

### ASN's

1. 65500 - Spines
2. 65501 - Leaf pair 01 *compute*
3. 65502 - Leaf pair 02 *compute*
4. 65503 - Leaf pair 03 *services*
5. 65505 - Leaf pair 04 *border*

## Fabric Nodes

### Spines

1. Spines only connect to leaf nodes (E.g., no inter-switch links)
2. Spine to leaf connectivity will leverage /31 subnets:
	1. First IP assigned to spine side
	2. Second IP assigned to leaf side
	
### Leafs

1. Leaf pairs:
	1. s1-p1-lf01 & s1-p1-lf02 		*leaf pair 1*
	2. s1-p1-lf03 & s1-p1-lf04 		*leaf pair 2*
	3. s1-p1-lf05 & s1-p1-lf06 		*leaf pair 3*
	4. s1-p1-lf07 & s1-p1-lf08 		*leaf pair 4*
2. vPC pair connectivity:
	1. 4 x links for vPC
	2. 2 x vPC keepalive
	3. 2 x vPC peer-link
3. vPC link details:
	1. vPC keepalive:
		1. 'Port-channel 10' 'L3':
			1. Member: Ethernet1/8
			2. Member: Ethernet1/9
	2. vPC peer-link:
		1. 'Port-channel 20' 'L2':
			1. Member: Ethernet1/10
			2. Member: Ethernet1/11
4. vPC peer connectivity
	1. 	lf01.e1.8 :: lf02.e1.8 		*port-channel10*
	2.	lf01.e1.9 :: lf02.e1.9		*port-channel10*
	3.	lf01.e1.10 :: lf02.e1.10	*port-channel20*
	4.	lf01.e1.11 :: lf02.e1.11	*port-channel20*

## Fabric uplink details:
	1. 	lf01.e1.1 :: sp01.e1.1 		10.1.1.1  :: 10.1.1.0	/31
	2. 	lf01.e1.2 :: sp02.e1.1 		10.1.1.21 :: 10.1.1.20	/31
	3. 	lf02.e1.1 :: sp01.e1.2 		10.1.1.3  :: 10.1.1.2	/31
	4. 	lf02.e1.2 :: sp02.e1.2 		10.1.1.23 :: 10.1.1.22	/31
	5. 	lf03.e1.1 :: sp01.e1.3 		10.1.1.5  :: 10.1.1.4	/31
	6. 	lf03.e1.2 :: sp02.e1.3 		10.1.1.25 :: 10.1.1.24	/31
	7. 	lf04.e1.1 :: sp01.e1.4 		10.1.1.7  :: 10.1.1.6	/31
	8. 	lf04.e1.2 :: sp02.e1.4 		10.1.1.27 :: 10.1.1.26	/31
	9. 	lf05.e1.1 :: sp01.e1.5 		10.1.1.9  :: 10.1.1.8	/31
	10.	lf05.e1.2 :: sp02.e1.5 		10.1.1.29 :: 10.1.1.28	/31
	11. lf06.e1.1 :: sp01.e1.6 		10.1.1.11 :: 10.1.1.10	/31
	12.	lf06.e1.2 :: sp02.e1.6 		10.1.1.31 :: 10.1.1.30	/31
	13.	lf07.e1.1 :: sp01.e1.7 		10.1.1.13 :: 10.1.1.12	/31
	14.	lf07.e1.2 :: sp02.e1.7 		10.1.1.33 :: 10.1.1.32	/31
	15.	lf08.e1.1 :: sp01.e1.8 		10.1.1.15 :: 10.1.1.14	/31
	16. lf08.e1.2 :: sp02.e1.8 		10.1.1.35 :: 10.1.1.34	/31


