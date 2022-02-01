# System Description

Shadowcorp, a geographically distributed corp, wants to create a url shortener for its clients.
Shadowcorp aims to monetize the service in the future and would like to collect as much user data as possible.
Each link must be unique.

## Initial information from Product Managers

* They estimate that 
  * users will generate 10 billion links per year using our service
  * each link will see over 100 clicks on average over its lifetime.
  * the number of clicks would be higher close to the link's creation. 
  * the number of clicks would drop off significantly over the lifetime of the link
  * users need to sign up to use the product
  * the latency of the redirect should be low (no estimates given)
  * service operators can delete proxied links but users can't
  * the availability of the url shortener should be high (99\% suggested).
  * Short links should be usable at most 5s after creation.


## Deductions from meeting with PMs

* On average, url lengths of 2000 characters i.e., 2000 bytes are acceptable on most browsers. 
* If we have 2kb on average for each link, we need a storage that can handle 2kb X 10 billion for a year.
  * 2 x 10^3 x 10^10 = 2 x 10^13 bytes of storage for the links.
* Putting 10^13 bytes on a single machine would require the machine to have about 20TB of storage.
* However, we need availability and reliability for our service. A single machine is a point of failure.
* We need distributed data stores for our service.

## System needs

* A program to perform the shortening of the links
* A data store to store the short links and map it to its proxy.
* A data store for user details
* A data store for link access details

## Service Architecture

* Loadbalancer that connects to multiple shortening services.

* Shortening service.
  * The shortening service is stateless.
  * It receives requests from the user
  * It retrieves responses from the data layer

* User details data layer
  * exposed by API
  * replicated data store
  * partitioned data store

* User details data layer
  * exposed by API
  * replicated data store
  * partitioned data store
  * key-value data store useful here

* Link access data layer
  * exposed by API
  * replicated data store
  * partitioned data store
  * relational data store useful here

* Link Data layer
  * exposed by API 
  * replicated data store
  * partitioned data store
  * key value data store is best for the mapping of short links to long links

