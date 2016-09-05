package main

import (
	"./asymptotic"
	"./mergesort"
	"fmt"
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
	asymptotic.Loops(5)

	asymptotic.NestedLoop(5)

	asymptotic.ConsecutiveStatements(5)

	asymptotic.IfThenElseStatements(5)

	asymptotic.LogarithmicComplexity(5)

	asymptotic.LogarithmicComplexityWorst(5)

	asymptotic.LogarithmicComplexityLogn1(20)

	asymptotic.LogarithmicComplexityLogn2(20)

	asymptotic.RecursiveFunction1(20)

	/** Merge Sort Example
	**/

	slice := mergesort.GenerateSlice(50)
	fmt.Printf("\n --- unsorted --- \n\n", slice)
	fmt.Printf("\n --- sorted ---\n\n", mergesort.MergeSort(slice), "\n")

}
