# igomessages

# Setup MQTT Broker
This project has been configured to use the open source MQTT Broker - Mosquitto (http://mosquitto.org/)
The following sections describe the installation and configurations made specific to this project's
requirements.

## Installation - Mosquitto
To host the mosquitto broker, use an AWS EC2 running Ubuntu 14.04 or higher.
To install mosquitto do the following:

* sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
* sudo apt-get update
* sudo apt-get install mosquitto
* sudo apt-get install mosquitto-clients

***
## Create a Certificate
To provide a level of security generate a certificate using the generateCA.sh script
found in ./igomessages/mosquitto/generateCA.sh

* ./generate-CA.sh <your broker's host system dns name or ip address>
* copy all generated files to /etc/mosquitto/certs
* sudo chown -R ubuntu:root /etc/mosquitto/certs
* sudo chmod -R 775 /etc/mosquitto/certs

***
## Setup Mosquitto Configuration
To ensure your configuration is using the proper configuration settings, here's an
example file you may use with modifications:

pid_file /var/run/mosquitto.pid

user root
autosave_interval 1800

connection_messages true
log_timestamp true

max_inflight_messages 0
max_queued_messages 9000

listener 1883

listener 8883
tls_version tlsv1
cafile /etc/mosquitto/certs/ca.crt
certfile /etc/mosquitto/certs/iv-1msdev.ytv.io.crt
keyfile /etc/mosquitto/certs/iv-1msdev.ytv.io.key

require_certificate false

persistence true
persistence_location /var/lib/mosquitto/
persistence_file iv1msdev.db

log_dest file /var/log/mosquitto/mosquitto.log
***
***
# Setup Go-lang Package
The foundation of this project is a go-lang package labeled "igomqtt". This package
contains a set of functions which allow the user to publish and subscribe messages.

## Pre-requisites
Once you have installed Go-lang and configured the GOPATH do the following:

* cd $GOPATH/src
* go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git
* cd $GOPATH
* git clone https://github3.cisco.com/jehollar/igomessages.git
* cd igomessages

## Test Installation
Once you have installed the go-lang package, to test it do the following:

* (Terminal 1)
* cd $GOPATH/igomessages
* go run subscribe.go

* (Terminal 2)
* cd $GOPATH/igomessages
* go run publish.go

***
***
