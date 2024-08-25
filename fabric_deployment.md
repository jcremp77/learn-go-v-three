# Ansible Fabric Deployment

## Deploy NX-OS Features

1. feature bgp
<pre>**ansible-playbook -i ./inventories/nxos-lp2.ini ./playbooks/nxos-feature/bgp-enable.yml --ask-vault-pass -vvvv**</pre>

feature interface-vlan
feature lacp
feature vpc
feature nv overlay


Enable VLAN-based VXLAN
Enable VXLAN
Enable VN-Segment for VLANs.
Enable Switch Virtual Interface (SVI).
Enable the EVPN control plane for VXLAN.
