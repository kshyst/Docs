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

## F Expression

Django supports the use of addition, subtraction, multiplication, division, modulo, and power arithmetic with F() objects, both with constants and with other F() objects. To find all the blog entries with more than twice as many comments as pingbacks, we modify the query:

## Q Expression

The ‘Q’ object is essentially a container for a single query expression, such as a filter condition, which can then be combined with other query expressions using logical operators such as AND, OR, and NOT. The ‘Q’ object allows you to build complex queries that are not easily achievable using standard query expressions.

you can save this Q(age__gt=18) objects and use them later


## Aggregation vs Annotation

Aggregation is used to calculate summary statistics over a set of objects. It involves summarizing the values of a field of a model. For example, you might want to calculate the average price of all books in your database.

```python
from django.db.models import Avg

Book.objects.all().aggregate(Avg('price'))
```

Annotation is used to add extra fields to each object in the QuerySet. These extra fields can be used to sort objects or filter them. For example, you might want to add a field to each book in your database that indicates whether the book is expensive.

```python
from django.db.models import F

Book.objects.all().annotate(is_expensive=F('price') > 100)
```

Annotate returns queryset but aggregate returns a value

## select_related() and prefetch_related()

#### select_related()

select_related() returns a QuerySet that will “follow” foreign-key relationships, selecting additional related-object data when it executes its query. This is a performance booster which results in a single more complex query but means later use of foreign-key relationships won’t require database queries.

```python
# Hits the database.
e = Entry.objects.get(id=5)

# Hits the database again to get the related Blog object.
b = e.blog
```

```python
# Hits the database.
e = Entry.objects.select_related("blog").get(id=5)

# Doesn't hit the database, because e.blog has been prepopulated
# in the previous query.
b = e.blog
```

#### prefetch_related()

Returns a QuerySet that will automatically retrieve, in a single batch, related objects for each of the specified lookups.

This has a similar purpose to select_related, in that both are designed to stop the deluge of database queries that is caused by accessing related objects, but the strategy is quite different.

select_related works by creating an SQL join and including the fields of the related object in the SELECT statement. For this reason, select_related gets the related objects in the same database query. However, to avoid the much larger result set that would result from joining across a ‘many’ relationship, select_related is limited to single-valued relationships - foreign key and one-to-one.

prefetch_related, on the other hand, does a separate lookup for each relationship, and does the ‘joining’ in Python. This allows it to prefetch many-to-many, many-to-one, and GenericRelation objects which cannot be done using select_related, in addition to the foreign key and one-to-one relationships that are supported by select_related. It also supports prefetching of GenericForeignKey, however, the queryset for each ContentType must be provided in the querysets parameter of GenericPrefetch.

```python
class Restaurant(models.Model):
    pizzas = models.ManyToManyField(Pizza, related_name="restaurants")
    best_pizza = models.ForeignKey(
        Pizza, related_name="championed_by", on_delete=models.CASCADE
    )
```

```python
Restaurant.objects.prefetch_related("pizzas__toppings")
```

This will prefetch all pizzas belonging to restaurants, and all toppings belonging to those pizzas. This will result in a total of 3 database queries - one for the restaurants, one for the pizzas, and one for the toppings.

```python
Restaurant.objects.prefetch_related("best_pizza__toppings")
```

This will fetch the best pizza and all the toppings for the best pizza for each restaurant. This will be done in 3 database queries - one for the restaurants, one for the ‘best pizzas’, and one for the toppings.

The best_pizza relationship could also be fetched using select_related to reduce the query count to 2:

```python
Restaurant.objects.select_related("best_pizza").prefetch_related("best_pizza__toppings")
```

Since the prefetch is executed after the main query (which includes the joins needed by select_related), it is able to detect that the best_pizza objects have already been fetched, and it will skip fetching them again.

