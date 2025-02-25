# Querying

## Manager Class

The Manager is the main source of QuerySets for a model. For example, Blog.objects.all() returns a QuerySet that contains all Blog objects in the database.

## Retrieve Specific Objects

- filter(**kwargs)
- exclude(**kwargs)

These return QuerySets

## Querysets Are Lazy

QuerySets are lazy – the act of creating a QuerySet doesn’t involve any database activity. You can stack filters together all day long, and Django won’t actually run the query until the QuerySet is evaluated. Take a look at this example:

```python
q = Entry.objects.filter(headline__startswith="What")
q = q.filter(pub_date__lte=datetime.date.today())
q = q.exclude(body_text__icontains="food")
print(q)
```

Though this looks like three database hits, in fact it hits the database only once, at the last line (print(q)). In general, the results of a QuerySet aren’t fetched from the database until you “ask” for them. When you do, the QuerySet is evaluated by accessing the database.

## OFFSET and LIMIT

```python
>>> Entry.objects.all()[5:10]
```

offset 5 limit 5

## exact

```python
>>> Entry.objects.get(headline__exact="Cat bites dog")
```
Would generate SQL along these lines:
```sql
SELECT ... WHERE headline = 'Cat bites dog';
```

## contains

```python
Entry.objects.get(headline__contains="Lennon")
```

```sql
SELECT * FROM Entry WHERE headline LIKE "%Lennon%"
```

## JOINs

Django automatically joins

Django offers a powerful and intuitive way to “follow” relationships in lookups, taking care of the SQL JOINs for you automatically, behind the scenes. To span a relationship, use the field name of related fields across models, separated by double underscores, until you get to the field you want.

```python
Blog.objects.filter(entry__headline__contains="Lennon")
```

