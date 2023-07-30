# Ports Service

Ports is a service that reads large json files with a known structure and writes the data to an in memory database. The application uses the "hexagonal architecture" pattern.

## Assumptions

I assumed that the large json file is already on the same machine on which the application is located, and a completely different service (e.g. FTP) is responsible for placing the file on the server. To simulate these assumptions, I added a folder "json_files" to the repository, this folder is copied to the docker container. The json file must be in this folder before the docker container is created.

## How to run?

The computer requires [docker](https://www.docker.com/) and [golang](https://go.dev/) software to be installed.
To run the application you need to:

1. build a docker image
2. start the container with the application

All actions can be performed with the `make` command:

- `make image`: creates docker image
- `make run`: runs container
- `make test`: runs tests

### Flags

The application has 3 configuration parameters (flags):

- folder_path string: the path to folder containing .json files with ports data - e.g. json_files/ (default "json_files/")
- graceful-timeout duration: the duration for which the server gracefully wait for existing connections to finish - e.g. 15s (default 15s)
- port string: the path to folder containing .json files with ports data - e.g. 8080 (default "8080")

Without changing the Dockerfile and by using the make command, the application will be launched with the default parameters.

### REST API

Once launched, the application will listen on port 8080 by default. REST API has 2 endpoints:

- GET: It is used to check the contents of the database, after loading the file you can ask for specific port IDs.
  - sample request: `curl localhost:8080/port/CNLZH` where "CNLZH" is a port ID.
- POST: It is used to indicate the name of the file that the application should read and write to the database. This file must be located in the "json_files" folder (endpoint should not be vulnerable to path traversal attack).
  - sample request: `curl localhost:8080/file -d "file_name=ports.json"`

## Repo structure

The repository has 3 main folders:

- bin: contains the compiled application in the form of a binary file
- internal: contains application code
- json_files: contains large json files

### App Structure

The application uses a hexagonal architecture pattern, an internal folder structure:

- `app`: contains an application service and ports (ports known from hexagonal architecture, not marine ports). There are only 2 ports(interfaces) so I didn't take them out to a separate package, but as if there would be more of them, it would be worth consider moving them to the 'app/ports' package.
- `domain`: It includes models and domain services, in this case only the "Port" model.
- `infrastructure`: Includes adapters
