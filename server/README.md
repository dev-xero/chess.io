# Chess.io: Backend

This is chess.io's backend source code.

## Pre-commit Hooks

I use [pre-commit](https://pre-commit.com/) to run some simple pre-git commit hooks, like formatting, linting, and commit message spell checks. See the previous link for instructions on how to set it up.

Using Homebrew:

```shell
brew install pre-commit
pre-commit install
```

## Environment Variables

The project makes use of environment variables loaded from a `.env` file at the backend root directory. There is an example config file at `.env.example`. Rename this file to `.env` then use your own environment specific variables.

```shell
chmod +x ./scripts/load_env.sh # do this once
source ./scripts/load_env.sh
```

Verify that it works.

```shell
env | grep PORT
```

## Docker Compose

This project uses Docker and Docker compose to run the backend and other container dependencies, such as Postgres. Make sure you have Docker installed on your system, more information [here](https://docs.docker.com/engine/install/).

```shell
docker compose down -v && docker compose up -d
```

If the containers build successfully, you may access the API at `http://localhost:8080/api/v1/`. To view container logs, use:

```shell
docker compose logs -f [container-name]
```

Replace `[container-name]` with the actual container name from the Docker Compose yaml file.
