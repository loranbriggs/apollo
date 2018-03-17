#!/bin/bash
# Install Apollo

# ensure you have:
# set hostname and user
# remove user pi
# configured wifi in raspi-config
# enabled ssh
# set auto login

if [ $(id -u) -ne 0 ]; then
echo "Please run as root"
  exit
fi

apt-get update
apt-get upgrade

echo
echo
echo
echo "Time to install Apollo!!!!!!!"

########################## Install Dependacies ###########################

apt-get install matchbox-window-manager xserver-xorg x11-xserver-utils midori unclutter xinit lsof

########################## Download Scripts ###########################

mkdir -p ~/scripts/

wget https://raw.githubusercontent.com/loranbriggs/apollo/master/scripts/launchapollo.sh -O ~/scripts/launchapollo.sh
chmod +x ~/scripts/launchapollo.sh

wget https://raw.githubusercontent.com/loranbriggs/apollo/master/apollo -O ~/apollo

wget https://raw.githubusercontent.com/loranbriggs/apollo/master/scripts/apollo -O /etc/ini.d/apollo
chmod +x /etc/ini.d/apollo

wget https://raw.githubusercontent.com/loranbriggs/apollo/master/scripts/fstab -O /etc/fstab

wget https://raw.githubusercontent.com/loranbriggs/apollo/master/scripts/mountfs.sh -O ~/scripts/mountfs.sh
chmod +x ~/scripts/mountfs.sh

########################## Launch on Startup ###########################

cat <<EOT >> ~/.bashrc
if [ -z "${SSH_TTY}" ]; then
  sudo service apollo start
  xinit ~/scripts/launchappollo.sh
fi
EOT
fi

##########################  DONE!!!!!!!!  ###########################

echo "Installed dependacies, launch scripts, and configured auto boot."
reboot
