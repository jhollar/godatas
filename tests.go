package main

import (
	"./igomqtt"
)

/**
  This is an example of how you can publish an Infinite Video Go
  Message to a specific Topic.
**/

func main() {
	/**
	  Construct a message object which requires the following:

	  BrokerUri:  The MQTT Broker URI
	  ClientId:   The unique client identifier for this request
	  Topic:      Messages in MQTT are published on topics. Topics are treated
	              as a hierarchy, using a slash (/) as a separator. Multiple
	              clients could then receive messages by creating a subscription
	              to this specific topic.
	  CertFile:   relative or absolute location of the certificate file
	  KeyFile:    relative or absolute location of the key file.

	**/
	igoMessage := igomqtt.GetIgMessage("ssl://iv-1msdev.ytv.io:8883", "jehollar", "test", "./igomqtt/ca.crt", "./igomqtt/ca.key")

	/**
	  Publish a message to a specific topic

	  Message Payload:  The message you are publishing

	  igoMessage: is the structure received from the GetIgMessage function
	**/
	igomqtt.PublishMessage("TEST MESSAGE Payload", &igoMessage)

}
