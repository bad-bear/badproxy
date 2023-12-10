#!/bin/bash

echo
echo '[APT Update]'
sudo apt update -y

echo
echo '[APT Upgrade]'
sudo apt upgrade -y

echo
echo '[Install GoLang]'
sudo apt remove --autoremove golang-go
sudo wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

echo
echo '[Setup IPtable Logging]'
sudo touch /var/log/fw.log
sudo apt install rsyslog
sudo systemctl enable rsyslog
sudo chmod 777 /etc/rsyslog.conf 
sudo echo 'kern.warning /var/log/fw.log' >> /etc/rsyslog.conf
sudo echo ':msg, contains, "IPTABLES" -/var/log/fw.log' >> /etc/rsyslog.conf
sudo echo '& ~' >> /etc/rsyslog.conf
sudo chmod 644 /etc/rsyslog.conf 
sudo service rsyslog restart

echo 
echo '[Setup Log Rules]'
sudo iptables -I INPUT -j LOG --log-prefix "[IPTABLES]: " --log-level 4
sudo iptables -I OUTPUT -j LOG --log-prefix "[IPTABLES]: " --log-level 4
sudo iptables -I FORWARD -j LOG --log-prefix "[IPTABLES]: " --log-level 4

# These two aren't necessary
#sudo apt install snapd
#sudo snap install --classic code
#sudo snap install --classic sqlitebrowser

echo '[Setup Complete]'
echo 'Run the following commands to start application:'
echo
echo 'export PATH=$PATH:/usr/local/go/bin'
echo 'sudo ./run.sh'

