# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|


  config.vm.define "testbox0" do |config|
    config.vm.box = "bento/centos-7"
    config.vm.hostname= "testbox0"
    config.vm.network "private_network", ip: "192.168.33.10" 
  end

  config.vm.define "testbox1" do |config|
    config.vm.box = "bento/centos-7"
    config.vm.hostname= "testbox1"
    config.vm.network "private_network", ip: "192.168.33.11"  
  end

  config.vm.define "testbox2" do |config|
    config.vm.box = "bento/centos-7"
    config.vm.hostname= "testbox2"
    config.vm.network "private_network", ip: "192.168.33.12"  
    config.vm.network "forwarded_port", guest: 22, host: 22
  end

  config.vm.define "testbox3" do |config|
    config.vm.box = "bento/centos-7"
    config.vm.hostname= "testbox3"
    config.vm.network "private_network", ip: "192.168.33.13"
  end

  config.vm.define "testbox4" do |config|
    config.vm.box = "bento/centos-7"
    config.vm.hostname= "testbox4"
    config.vm.network "private_network", ip: "192.168.33.14" 
  end

  config.vm.provision "shell", inline: <<-SHELL
    touch /home/vagrant/chef-log.txt
    echo "generating random 1gb size file"
    dd if=/dev/urandom of=example-log.txt bs=1048576 count=1024
    
    echo "generating random 2gb size file"
    dd if=/dev/urandom of=example-log-2.txt bs=1048576 count=2048
    rm -rf ~/.ssh/known_hosts
  SHELL
end
