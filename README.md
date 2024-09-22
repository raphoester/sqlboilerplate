# SQLBoilerplate: SQlBoiler setup made easy

Setting up [sqlboiler](github.com/volatiletech/sqlboiler) can be a pain in the ass.

Running a live DB and re-syncing migrations with boil code every time you update them is a super annoying.

SQLBoilerplate fixes that.

It simplifies model generation with a single make target, making your workflow easier.

## Getting Started

**Requirements**
- Docker and Docker compose 
- [sqlboiler](github.com/volatiletech/sqlboiler) command line utility
- [go migrate](https://github.com/golang-migrate/migrate) for creating new migrations

Clone the repository :
```shell
git clone git@github.com:raphoester/sqlboilerplate.git
cd sqlboilerplate
```

It's a boilerplate, so this is essentially your project structure. 

In order to generate the boiler code, make sure you're at the root of the project and run :
```shell
make boil 
```

This command does a couple of things: 
- Creates a fresh docker instance that runs PostgreSQL
- Executes the migration script (in `migrations/comments`)
- Runs sqlboiler against that DB instance and puts the generated code in the `generated/sqlboiler` directory.

Magic, magic : run it and the boil code will spawn in the `generated/sqlboiler/comments` directory.

(`Comments` is the name of the default schema I set up for this example).

## Modular monolith baby

### Generating your own migrations and models

The project's structure has been designed for handling multiple bounded contexts within the same monolith 
(Each bounded context having its own schema and associated boiler code).

`Comments` is the default one I set up for the example.
Feel free to delete it.

Of course, you can create your own migrations and models.

For that, there is a make target.

For example, let's imagine that you want to create a new BC called `shopping` with a first migration called `products`:

```shell
make migration CTX=shopping NAME=products
````

### Switching between BCs

Modular monolith = multiple BCs in the same runtime/project.

That means you can switch between BCs and generate boiler code for each of them.

In order to switch DB schemas and associated boiler code, simply use the CTX variable:
```shell
make boil CTX=something
```

