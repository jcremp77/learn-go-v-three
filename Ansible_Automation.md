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
	1. Enter configuration mode: **Config t**
	2. Run the following command: **boot nxos bootflash:/nxos.9.3.9.bin**
	3. Save configuration: **copy run start**
	4. Change the hostname: **hostname s1-p1-sp01**
	5. Generate crypto key: **cry key gen rsa mod 2048**
	6. Add Secure-Copy: **feature scp-server**
	7. Configure management interface:
		<pre>
		interface mgmt 0
		ip address ip.ip.ip.ip/cidr</pre>
	8. Create user and role: **user ansible role network-admin password somepassword**
<!-- blank -->
4. From tools server - scp the public key into bootflash:
	<pre>
	Ex. Ubuntu: "scp ~/.ssh/id_rsa.pub ansible@192.168.1.161:"</pre>
5. From switch - enable authentication of the previously created user via ssh keys
	<pre>
	example: "username ansible sshkey file bootflash:id_rsa_ansible.pub"</pre>
6. Setup Ansible to login using ssh_private_key_file as opposed to password.
<!-- blank -->
7. ...
		
	