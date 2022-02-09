# Metrics Aggregation Service

Design of a global service that aggregates metrics from all other services in the larger system

## Requirements
* Gather performance data from all services
  * Ideally in a manner that aids pattern detection
* Metrics are collected every minute (or less)
* Basic metrics are automatically collected
  * But devs set up other metrics too
* Metrics are presented on a dashboard
  * Must be possible to search, filter and tag

### Larger System Description
* Image hosting service
* URL Shortening Service
* Question & Answer Service

### Data Retention
* Recent data are more important than historical data
  * Meaning we could compress historical data

### Growth Projections
* Amount of metrics will double yearly

### Usage Requirements
* Most of the engineers will interact constantly with this service
  * Approx 100 engineers

### Budget
* High priority project for the company
  * If you can justify the cost, it will be approved
  * Small team of engineers for 1 year

## Error Budget 
* Some errors are allowed but not specified

## Calculations

## SLOs
* Most recent metrics must be viewable after 1m

## Architecture

## References
