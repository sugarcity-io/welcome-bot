<!-- Create a Readme for this slack chat bot -->

# Sugarcity.io Slack Chat Bot

This is a slack chat bot that is used to help manage the slack chat for the sugarcity.io slack chat.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

* `go` version `1.19` or higher
* Slack account
* Slack workspace for the bot to be installed into for testing purposes.

### API and Bot Tokens

When registering the bot to your test workspace you will need to create an API and bot token for the bot to work.
Both of these tokens will need to be set as environment variables.

## Local Development

To get the bot running locally, first download the required go modules by running the following command in the root of the project.

```
go get -u ./...
```

Then set the two environment variables required for the slack bot:
    
```
SLACK_API_TOKEN
SLACK_BOT_TOKEN
```

Setting the environment variables can either be done in your terminal or by adding them to a .env file in the root of the project, and then passing them to the bot when it is started.

### Running the Bot Locally

To start the bot, run the following command in the root of the project:

```
go run ./cmd/sugarcitybot
```

### Building the Bot

The chat bot can be built two ways, either as a binary file or as a docker image.

#### Building the Binary File

To build the bot, run the following command in the root of the project:

```
go build ./cmd/chat-bot
```

This will create a binary file in the root of the project called `chat-bot`, which can then be run by running the following command:

```
./chat-bot
```

#### Building the Docker Image

To build the docker image, run the following command in the root of the project:

```
docker build -t chat-bot .
```

This will create a docker image called `chat-bot`, which can then be run by running the following command:

```
docker run --env-file <path-to-dot-env> -d chat-bot
```

## Deployment

Once you are satisfied with your changes, raise a pull request and get it reviewed and merged into the `main` branch.

Once merged, the bot will be automatically deployed to the production environment, and the changes will be live in the official Sugarcity.io slack workspace.

Deployment is managed by [Railway.app](https://railway.app/).

## To-Do

* Write tests
* Add more commands
* Add more features
