# apache-pulsar-demo
high-performance Apache Pulsar project designed to deliver scalable, fault-tolerant messaging and real-time streaming

## Apache Pulsar Messaging Project

This project demonstrates a simple end-to-end messaging pipeline using Apache Pulsar, featuring a Producer, Consumer, and Prometheus-based metrics collection.

## Overview

Apache Pulsar is a cloud-native, distributed messaging and streaming platform. This project showcases a minimal setup where:

#### A Producer publishes messages to a Pulsar topic

#### A Consumer subscribes to that topic and processes incoming messages

#### Prometheus monitors Pulsar component metrics using a prometheus.yml configuration file

## Producer

The Producer is responsible for sending messages to a specified Pulsar topic.

### Key responsibilities:

Establish a connection to the Pulsar broker

### Create or use an existing topic

Publish messages asynchronously or synchronously

Handle retries, batching, and acknowledgments

Producers are typically used for event generation, logging, data ingestion, or streaming workloads.

## Consumer

The Consumer listens on the same topic and processes the messages published by the Producer.

### Key responsibilities:

Subscribe to a topic using a subscription mode (e.g., Exclusive, Shared, Failover)

Receive and acknowledge messages

Process messages reliably and fault-tolerantly

Handle redelivery and backoff if configured

Consumers form the core of data processing pipelines and event-driven architectures.

## Role of prometheus.yml

The prometheus.yml file defines how Prometheus scrapes metrics from Pulsar services or other system components.

This file typically includes:

Pulsar endpoints exposing metrics

Scrape intervals

Job names

Optional labels for environment or instance tagging

Prometheus metrics help monitor:

Producer/Consumer throughput

Message backlog and latency

Broker performance and system health
```
Project Structure (Example)
.
├── producer/
│   └── producer.py
├── consumer/
│   └── consumer.py
├── prometheus.yml
└── README.md
```

Getting Started

Start Apache Pulsar (standalone or cluster).

Run the Producer to publish messages.

Run the Consumer to process messages.

Launch Prometheus with prometheus.yml to begin scraping metrics.
