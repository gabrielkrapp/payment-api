# Payment Processing System with Kafka Integration

## Overview

This project is a payment processing system that integrates Apache Kafka for message queue management and Amazon Simple Email Service (SES) for sending email notifications. The system is divided into two main parts:

1. **Payment API (`payment-api`)**: Responsible for processing payments and posting messages to a Kafka topic.
2. **Kafka Worker (`worker`)**: Consumes messages from the Kafka topic and sends email confirmations of payments.

## Features

- **Payment Processing**: Allows users to process payments via Stripe.
- **Kafka Integration**: Uses Kafka to manage confirmation messages of payments.
- **Email Notification**: Sends automatic emails to users confirming their payments using Amazon SES.
- **Monorepo Setup**: Both applications are set up within a single repository for ease of management and development.

## Technologies Used

- **Go**: Programming language used to develop both applications.
- **Apache Kafka**: Messaging system for processing and managing payment queues.
- **Amazon SES**: Email service for sending payment notifications.
- **Docker**: Used for containerization and easy distribution of applications.
- **Stripe**: Payments API to process financial transactions.

## Getting Started

### Prerequisites

- Go (version 1.x)
- Docker
- An instance of Apache Kafka
- Configuration of a domain in Amazon SES

### Installation

1. **Clone the Repository**

```bash
git clone https://github.com/your-username/payment-api.git
cd payment-api
```

2. **Configure Environment Variables**

Create an `.env` file at the root of the project and adjust the variables according to your setup:

```plaintext
STRIPE_KEY=your_stripe_api_key
KAFKA_BROKERS=localhost:9092
AWS_REGION=your_aws_region
AWS_ACCESS_KEY_ID=your_aws_access_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
```

3. **Start Kafka**

Ensure that Kafka is running and the `payment-intents` topic is created.

```bash
# Example using Docker
docker-compose up -d kafka
```

4. **Run the Payment API**

```bash
cd payment-api
go run main.go
```

5. **Run the Kafka Worker**

```bash
cd worker
go run main.go
```
