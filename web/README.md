# Unicorn History Server Components

**Disclaimer:** This project is currently in development and is in the pre-alpha phase. We warmly welcome all input, contributions, and suggestions.

Unicorn History Server Components provides a remote components that are consumed by the [YuniKorn Web](https://github.com/G-Research/yunikorn-web/)

## Development environment setup

### Dependencies

- [Node.js](https://nodejs.org/en/)
- [Angular CLI](https://github.com/angular/angular-cli)

For managing node packages you can use `npm`, `yarn` or `pnpm`. Run `npm install` to set up all necessary dependencies.

### Development Server

#### Setting up the environment

To establish a connection between the web UI and UHS, you need to configure the `src/assets/config/envconfig.json` file with the following values:

```json
{
  "localUhsComponentsWebAddress": "http://localhost:3100",
  "externalLogsURL": "https://logs.example.com?token=abc123&applicationId=",
  "yunikornApiURL": "http://localhost:30001",
  "uhsApiURL": "http://localhost:8989"
}
```

#### Running the development server

Follow these steps to run the development server:

1. Follow the instructions in the UHS to run the server.
2. Run `pnpm start` for a dev server. Remote components will be served from this path `http://localhost:3100/remoteEntry.js`. The application will automatically reload if you change any of the source files.
3. Follow the instructions in the [YuniKorn Web](https://github.com/G-Research/yunikorn-web/) repository to set up the development environment. This is required to run the web UI.

### JSON Server

To run a mock server for local development, follow these steps:

**Start the JSON Server**:

- **Using Makefile**: you can start the server by running:

  ```sh
  make mock-server
  ```

- **Using pnpm**: If you are in the `./web` directory, you can run the JSON Server directly with pnpm by using:
  ```sh
  pnpm run start:json-server
  ```

This will start the JSON Server and serve mock data. You can access it at `http://localhost:3000`.

Some endpoints that can be tested with ID's are:

- Get partitions: `GET http://localhost:3000/api/v1/partitions`
- Get queues for partition ID `01JEE8TVV09AYGJT40Z2ZBN972`: `GET http://localhost:3000/api/v1/partition/01JEE8TVV09AYGJT40Z2ZBN972/queues`
- Get application for partition ID `01JEE8TVV09AYGJT40Z2ZBN972` and queue ID `01JEE8TVV05C707SVK0XG8EPVQ`: `GET http://localhost:3000/api/v1/partition/01JEE8TVV09AYGJT40Z2ZBN972/queue/01JEE8TVV05C707SVK0XG8EPVQ/applications`

### Build

Run `make web-build` from the project root or `pnpm run build`. Build output is set to `/assets` folder in project root as it will be served from the UHS server.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).

## Code scaffolding

Run `ng generate component component-name` to generate a new component.

You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.
