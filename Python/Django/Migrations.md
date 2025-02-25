# Migrations

Think of migrations as a version control system for database schema:

- migrate, which is responsible for applying and unapplying migrations.

- makemigrations, which is responsible for creating new migrations based on the changes you have made to your models.

- sqlmigrate, which displays the SQL statements for a migration.

- showmigrations, which lists a project’s migrations and their status.

makemigrations is responsible for packaging up your model changes into individual migration files - analogous to commits - and migrate is responsible for applying those to your database.

Django will make migrations for any change to your models or fields - even options that don’t affect the database - as the only way it can reconstruct a field correctly is to have all the changes in the history, and you might need those options in some data migrations later on (for example, if you’ve set custom validators).

```shell

$ python manage.py makemigrations
Migrations for 'books':
  books/migrations/0003_auto.py:
    - Alter field author on book
```

Your models will be scanned and compared to the versions currently contained in your migration files, and then a new set of migrations will be written out. Make sure to read the output to see what makemigrations thinks you have changed 

## Non Atomic Transaction

On databases that support DDL transactions (SQLite and PostgreSQL), all migration operations will run inside a single transaction by default. In contrast, if a database doesn’t support DDL transactions (e.g. MySQL, Oracle) then all operations will run without a transaction.

You can prevent a migration from running in a transaction by setting the atomic attribute to False. For example:

```python
from django.db import migrations


class Migration(migrations.Migration):
    atomic = False
```

## Dependencies

While migrations are per-app, the tables and relationships implied by your models are too complex to be created for one app at a time. When you make a migration that requires something else to run - for example, you add a ForeignKey in your books app to your authors app - the resulting migration will contain a dependency on a migration in authors.

This means that when you run the migrations, the authors migration runs first and creates the table the ForeignKey references, and then the migration that makes the ForeignKey column runs afterward and creates the constraint. If this didn’t happen, the migration would try to create the ForeignKey column without the table it’s referencing existing and your database would throw an error.

## Migration Files

Migrations are stored as an on-disk format, referred to here as “migration files”. These files are actually normal Python files with an agreed-upon object layout, written in a declarative style.

```python
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [("migrations", "0001_initial")]

    operations = [
        migrations.DeleteModel("Tribble"),
        migrations.AddField("Author", "rating", models.IntegerField(default=0)),
    ]
```

What Django looks for when it loads a migration file (as a Python module) is a subclass of django.db.migrations.Migration called Migration. It then inspects this object for four attributes, only two of which are used most of the time:

- dependencies, a list of migrations this one depends on.

- operations, a list of Operation classes that define what this migration does.

The operations are the key; they are a set of declarative instructions which tell Django what schema changes need to be made. Django scans them and builds an in-memory representation of all of the schema changes to all apps, and uses this to generate the SQL which makes the schema changes.

That in-memory structure is also used to work out what the differences are between your models and the current state of your migrations; Django runs through all the changes, in order, on an in-memory set of models to come up with the state of your models last time you ran makemigrations. It then uses these models to compare against the ones in your models.py files to work out what you have changed


## Adding migrations to apps

New apps come preconfigured to accept migrations, and so you can add migrations by running makemigrations once you’ve made some changes.

If your app already has models and database tables, and doesn’t have migrations yet (for example, you created it against a previous Django version), you’ll need to convert it to use migrations by running:

```shell

python manage.py makemigrations your_app_label
```
This will make a new initial migration for your app. Now, run python manage.py migrate --fake-initial, and Django will detect that you have an initial migration and that the tables it wants to create already exist, and will mark the migration as already applied. (Without the migrate --fake-initial flag, the command would error out because the tables it wants to create already exist.)

## Reverse Migration

Migrations can be reversed with migrate by passing the number of the previous migration. For example, to reverse migration books.0003:

```shell

python manage.py migrate books 0002
```

If you want to reverse all migrations applied for an app, use the name zero:

## Data Migrations

As well as changing the database schema, you can also use migrations to change the data in the database itself, in conjunction with the schema if you want.

Migrations that alter data are usually called “data migrations”; they’re best written as separate migrations, sitting alongside your schema migrations.

Django can’t automatically generate data migrations for you, as it does with schema migrations, but it’s not very hard to write them. Migration files in Django are made up of Operations, and the main operation you use for data migrations is RunPython.

To start, make an empty migration file you can work from (Django will put the file in the right place, suggest a name, and add dependencies for you):

```shell

python manage.py makemigrations --empty yourappname
```

Now, all you need to do is create a new function and have RunPython use it. RunPython expects a callable as its argument which takes two arguments - the first is an app registry that has the historical versions of all your models loaded into it to match where in your history the migration sits, and the second is a SchemaEditor, which you can use to manually effect database schema changes (but beware, doing this can confuse the migration autodetector!)