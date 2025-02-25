# Aggregations

## aggregate() vs annotate()

aggregate() is to generate summary values over an entire QuerySet. For example, say you wanted to calculate the average price of all books available for sale. Django’s query syntax provides a means for describing the set of all books:

```python
Book.objects.aggregate(Avg("price"))
```

annotate()  is to generate an independent summary for each object in a QuerySet. For example, if you are retrieving a list of books, you may want to know how many authors contributed to each book. Each Book has a many-to-many relationship with the Author; we want to summarize this relationship for each book in the QuerySet.

```python
# Build an annotated queryset
>>> from django.db.models import Count
>>> q = Book.objects.annotate(Count("authors"))
# Interrogate the first object in the queryset
>>> q[0]
<Book: The Definitive Guide to Django>
>>> q[0].authors__count
2
# Interrogate the second object in the queryset
>>> q[1]
<Book: Practical Django Projects>
>>> q[1].authors__count
1
```

Unlike aggregate(), annotate() is not a terminal clause. The output of the annotate() clause is a QuerySet; this QuerySet can be modified using any other QuerySet operation, including filter(), order_by(), or even additional calls to annotate().

