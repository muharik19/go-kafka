# Learning Kafka and Connection Kafka To Golang Application

## Contents

- [Learning Kafka and Connection Kafka To Golang Application](#learning-kafka-and-connection-kafka-to-golang-application)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
  - [Kafka CLI](#kafka-cli)

## Installation

To install Kafka package, you need to install Go and set your Go workspace first.

1.  You first need [Go](https://golang.org/) installed (**version 1.18+ is required**), then you can use the below Go command to install Kafka.

```sh
$ go get -u github.com/confluentinc/confluent-kafka-go/kafka
```

2. Import it in your code:

```go
import "github.com/confluentinc/confluent-kafka-go/kafka"
```

3. You first need Makefile

```sh
# install windows
C:\> choco install make
```

```sh
# install linux
$ sudo apt-get install make
```

```sh
# install MacOS with Homebrew
$ brew install make
```

#### Verify Version Makefile

```sh
# verify that Makefile is installed correctly by checking the version
$ make --version
```

4. You first need [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed.

5. You first need Docker Compose.

#### Windows and macOS

- **Docker Compose is included in Docker Desktop for Windows and macOS.**

#### Linux

- **For Ubuntu and Debian, run:**

```sh
$ sudo apt-get update
$ sudo apt-get install docker-compose-plugin
```

- **For RPM-based distros, run:**

```sh
$ sudo yum update
$ sudo yum install docker-compose-plugin
```

#### Verify Version Docker Compose

```sh
# verify that Docker Compose is installed correctly by checking the version
$ docker compose version
```

6. You first need Kafka CLI or Web UI

```sh
# install kafka cli, make sure all container running
$ make compose
```

```sh
# install kafka web ui, make sure all container running, wait 1 minutes, then open to browser http://localhost:9021
$ make compose-ui
```

## Quick Start

- Producer

```sh
# application producer, run:
$ make producer
```

- Consumer

```sh
# application consumer1, run:
$ make consumer1
```

```sh
# application consumer2, run:
$ make consumer2
```

## Kafka CLI

**Note that you need to set the corresponding binding tag on Operating System. For example: `winpty`, `sudo` or etc**

```sh
# entry into docker container kafka, run:
$ make exec
```

#### Create Topic

```sh
$ kafka-topics --bootstrap-server localhost:9092 --create --topic YOURTOPIC
```

#### Create Topic with specific configuration

```sh
$ kafka-topics --bootstrap-server localhost:9092 --create --topic YOURTOPIC --replication-factor 1 --partitions 3 --config "cleanup.policy=compact" --config "delete.retention.ms=100"  --config "segment.bytes=204800" --config "min.cleanable.dirty.ratio=0.01"
```

#### List Topic

```sh
$ kafka-topics --bootstrap-server localhost:9092 --list
```

#### Delete Topic

```sh
$ kafka-topics --bootstrap-server localhost:9092 --delete --topic YOURTOPIC
```

#### Alter Topic

```sh
$ kafka-topics --bootstrap-server localhost:9092 --alter --topic YOURTOPIC --partitions 3
```

#### List Detail Topic

```sh
$ kafka-topics --bootstrap-server localhost:9092 --describe --topic YOURTOPIC
```

#### Send Message Producer

```sh
$ kafka-console-producer --broker-list localhost:9092 --topic YOURTOPIC
```

#### Received Message Consumer

```sh
$ kafka-console-consumer --bootstrap-server localhost:9092 --topic YOURTOPIC --from-beginning
```

#### Received Message Consumer Group

```sh
$ kafka-console-consumer --bootstrap-server localhost:9092 --topic YOURTOPIC --group YOURGROUP --from-beginning
```

#### List Consumer Group Offset

```sh
$ kafka-consumer-groups --bootstrap-server localhost:9092 --all-groups --all-topics --describe
```

#### List Consumer Group Offset 1

```sh
$ kafka-consumer-groups --bootstrap-server localhost:9092 --group YOURGROUP --describe
```
