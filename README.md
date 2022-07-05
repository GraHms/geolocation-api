Introduction
-
**Geolocation Service**

The Geolocation Service allows to import geolocation data from a csv and expose it via an API.

Quickstart
----------



First run the following commands to bootstrap your environment with ``make``: ::

    git clone https://github.com/GraHms/geolocation-api.git
    cd super-agent-sale
    make dep

Then create ``config.env`` file (or rename and modify ``.env.example``) in project root and set environment variables for application: ::

    touch config.env
    echo DB_DRIVER=postgres
    echo DB_SOURCE=postgresql://postgres:west04@localhost:5432/geoloationsdb?sslmode=disable
    echo SERVER_ADDR=0.0.0.0:8080
    echo DEBUG=True
    echo POSTGRESQL_DB_URL=localhost >> config.env


To run the web application in debug use::

    make build

If you run into the following error in your docker container:

sqlalchemy.exc.OperationalError: (psycopg2.OperationalError) could not connect to server: No such file or directory
Is the server running locally and accepting
connections on Unix domain socket "/tmp/.s.PGSQL.5432"?

Ensure the DB_CONNECTION variable is set correctly in the `config.env` file.
It is most likely caused by POSTGRES_HOST not pointing to its localhost.

DB_CONNECTION=postgresql://postgres:postgres@0.0.0.0:5432/



Run tests
---------


To run all the tests of a project, simply run the ``make test`` command: ::



----------------------

You must have ``docker`` and ``docker-compose`` tools installed to work with material in this section.
First, create ``config.env`` file like in `Quickstart` section or modify ``.env.example``.
``POSTGRES_HOST`` must be specified as `db` or modified in ``docker-compose.yml`` also.
Then just run::

    docker-compose up -d db


Application will be available on ``localhost`` in your browser.

Web routes
----------

All metrics of compromised rows are available on  ``/metrics``.


Project structure
-----------------


Files related to the Helm Chart are in ``.k8s`` directory
Application parts are:

