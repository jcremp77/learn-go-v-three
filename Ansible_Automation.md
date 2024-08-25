# Ansible EVPN/VXLAN Fabric Automation
<!-- blank -->
## Create Lab Topology
<!-- blank -->
1. Deploy nodes in lab and establish layer-1 connectivity per the architecture.
2. Test console reachability to all nodes.
<!-- blank -->
<!-- blank -->
## Prepare For Automation
<!-- blank -->
1. Deploy a linux or windows server with: 'ansible', 'python3', 'Jinja2', and 'Terraform'.
	1. This will be known as the 'tools' server.
	2. Install 'openssh' and generate key-pair for later use.
<!-- blank -->
2. Setup out-of-band network for management access to/from all nodes including the tools server.
<!-- blank -->
3. Establish console session to all network nodes and perform the following tasks:
<!-- blank -->
4. Configure nxos boot: <pre>**boot nxos bootflash:/nxos.9.3.9.bin**</pre>
<!-- blank -->
5. Change the hostname: <pre>**hostname s1-p1-sp01**</pre>
<!-- blank -->
6. Generate crypto key: <pre>**cry key gen rsa mod 2048**<pre>
<!-- blank -->
7. Add Secure-Copy: <pre>**feature scp-server**</pre>
<!-- blank -->
8. Configure management interface:
		<pre>
		interface mgmt 0
		ip address ip.ip.ip.ip/cidr</pre>
9. Create user and role: <pre>**user ansible role network-admin password somepassword**</pre>
<!-- blank -->
10. From tools server - copy ssh public key to switch: <pre>**scp ~/.ssh/id_rsa.pub ansible@192.168.1.161:**</pre>
<!-- blank -->
11. From switch - enable authentication of the previously created user via ssh keys:
	<pre>**username ansible sshkey file bootflash:id_rsa_ansible.pub**</pre>
12. Setup Ansible to login using ssh_private_key_file as opposed to password.
...
		
	