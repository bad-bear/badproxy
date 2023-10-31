# General Notes

- Good Practice clear IPtables 
```console
sudo iptables -F
```

- When making default changes on INPUT/OUTPUT need to make two rules for both



# Form
### Flags
* null means no input that applies to all

- p - protocol options: all, tcp, udp, icmp
- s - source options: ip, network addr, hostname, null
- d - destination options: ip, network addr, hostname, null
- j - target options: ACCEPT, DROP, QUEUE, RETURN
- i - in interface options: manual input
- o - out interface options: manual input


### Flags Additional Options
- sport
- dport


## Add Firewall Rule
Ex: iptables -A INPUT -s xxx.xxx.xxx.xxx -j DROP
Ex: iptables -A INPUT -p tcp -s xxx.xxx.xxx.xxx -j DROP
Ex: iptables -A INPUT -i eth0 -p tcp -s xxx.xxx.xxx.xxx -j DROP
Ex: iptables -A OUTPUT -o eth0 -p tcp -d 192.168.100.0/24 --dport 22 -m state --state NEW,ESTABLISHED -j ACCEPT


Inputs Needed:
- Table
- Protocol
- Source IP
- Destination IP
- Interface
- Target

## Add Geo Rule
## Add Port Rule

Ex: iptables -A OUTPUT -p udp -o eth0 --dport 53 -j ACCEPT
Ex: iptables -A INPUT -p udp -i eth0 --sport 53 -j ACCEPT

Inputs Needed:
- Table
- Protocol
- Interface
- Port
- Target

## Add IP Rate Rule
Ex: iptables -p tcp --syn --dport 23 -m connlimit --connlimit-above 2 -j REJECT
Ex: iptables -A FORWARD -p tcp --syn -m limit --limit 1/s -j ACCEPT

Inputs Needed:
- protocol
- time integer
- time type

## Add Bandwidth Rule
