``` 
- ST0258, Roller Tempo

- Juan David Valencia Torres, jdvalencit@eafit.edu.co
- Julian Giraldo Perez, jgiraldop@eafit.edu.co
- Kevin Mauricio Loaiza Arango, kmloaizaa@eafit.edu.co
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Daniel Arango Hoyos, darangoh@eafit.edu.co
```

---

## Table of contents[](#table-of-contents)
- [Introduction](#introduction)
- [Motivation](#motivation)
- [Composition](#composition)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Cloning the repository](#cloning-the-repository)
  - [Setting up the database](#setting-up-the-database)
  - [Running the webapi](#running-the-webapi)
  - [Running the webapp](#running-the-webapp)
- [Roadmap](#roadmap)
- [License](#license)
- [Call to action](#call-to-action)
- [Contact](#contact)

---

## Introduction[](#introduction)

**Roller Tempo** is a complete implementation of a turn management service for amusement parks. It was developed using Go + [Echo (Framework)](https://echo.labstack.com/) for the API management and Python + [Django (Framework)](https://www.djangoproject.com/) for the desktop and mobile applications. It was developed following coding best practices as well as SOLID principles. 

---

## Motivation[](#motivation)

This project was developed as part of a course work in conjunction with the company Globant (who acted as advisors, supervisors and evaluators). It was carried out with the intention of putting into practice concepts oriented towards fullstack development.

---

## Composition[](#composition)

Generally speaking, **Roller Tempo** has as main components:

- For the server:
    - **Programming Language:** Golang.
    - **Framework:** Echo.
    - **Database:** PostgreSQL.
    - **ORM:** Gorm.

- For the client:
    - **Programming Language:** Python.
    - **Framework:** Django.
    - **Rendering:** Django Templates.

We seek to manage a concise structure for the API design, with the following components:
- _Repositories_, where to perform the management of CRUD operations with the system models (this was done using Gorm ORM).
- _Services_, where the business logic is handled (Required functions to be offered to whoever consumes the api).
- _Controllers_, responsible for the correct handling of requests and their responses.

---

## Getting Started[](#getting-started)

### Prerequisites[](#prerequisites)

* **Golang:** Roller Tempo requires that you have Go language installed on your system to build and run the application.

* **Python:** It is also necessary that you have Python 3 installed on your system so the client can run properly.

* **Docker-compose:** You need to have docker-compose installed since the database is run through containers.

### Cloning the repository[](#cloning-the-repository)

First, you need to clone the repository from GitHub. To do this, you can open a terminal and run the following command:

```
git clone https://github.com/kevinloaiza12/roller-tempo.git
```

### Setting up the database[](#setting-up-the-database)

```
cd docker
./service.sh up
```

### Running the webapi[](#running-the-webapi)

```
cd app/server
go run .
```

### Running the webapp[](#running-the-webapp)

```
cd app/client
Python manage.py runserver
```

---

## Roadmap[](#roadmap)

The objective of Roller Tempo is to put development concepts into practice as well as to apply design patterns, architecture principles, and other relevant concepts.

Here are some of the aspects that could be considered in the future for the implementation of the project:
- Authentication + User validation.
- Session management.
- Turn notification system.

These features are open to public intention to contribute to open-source development.

---

## License[](#license)

Roller Tempo is under free open source [_MIT License_](https://github.com/kevinloaiza12/roller-tempo/blob/master/LICENSE.txt).

---

## Call to action[](#call-to-action)

Join this project and provide assistance by:
* Checking out the list of [open issues](https://github.com/kevinloaiza12/roller-tempo/issues) where Roller Tempo could need help.
* If you need new features, please open a [new issue](https://github.com/kevinloaiza12/roller-tempo/issues) or start a [discussion](https://github.com/kevinloaiza12/roller-tempo/discussions).

Feel free to contact me if you have any questions or contributions to make to the project.
