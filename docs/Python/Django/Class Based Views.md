## Async Class Based Views

```python
import asyncio
from django.http import HttpResponse
from django.views import View


class AsyncView(View):
    async def get(self, request, *args, **kwargs):
        # Perform io-blocking view logic using await, sleep for example.
        await asyncio.sleep(1)
        return HttpResponse("Hello async world!")
```

Within a single view-class, all user-defined method handlers must be either synchronous, using def, or all asynchronous, using async def. An ImproperlyConfigured exception will be raised in as_view() if def and async def declarations are mixed.

## How it works

Because Django’s URL resolver expects to send the request and associated arguments to a callable function, not a class, class-based views have an as_view() class method which returns a function that can be called when a request arrives for a URL matching the associated pattern. The function creates an instance of the class, calls setup() to initialize its attributes, and then calls its dispatch() method. dispatch looks at the request to determine whether it is a GET, POST, etc, and relays the request to a matching method if one is defined, or raises HttpResponseNotAllowed if not:

## Dispatching

divides requests to be handled by different methods for example get request will be handled by get method.

## Inheriting Generics

Note also that you can only inherit from one generic view - that is, only one parent class may inherit from View and the rest (if any) should be mixins. Trying to inherit from more than one class that inherits from View - for example, trying to use a form at the top of a list and combining ProcessFormView and ListView - won’t work as expected.

## Decorating class-based views

```python
from django.contrib.auth.decorators import login_required, permission_required
from django.views.generic import TemplateView

from .views import VoteView

urlpatterns = [
    path("about/", login_required(TemplateView.as_view(template_name="secret.html"))),
    path("vote/", permission_required("polls.can_vote")(VoteView.as_view())),
]
```

```python
decorators = [never_cache, login_required]


@method_decorator(decorators, name="dispatch")
class ProtectedView(TemplateView):
    template_name = "secret.html"


@method_decorator(never_cache, name="dispatch")
@method_decorator(login_required, name="dispatch")
class ProtectedView(TemplateView):
    template_name = "secret.html"
```

In this example, every instance of ProtectedView will have login protection. These examples use login_required, however, the same behavior can be obtained by using LoginRequiredMixin.

The decorators will process a request in the order they are passed to the decorator. In the example, never_cache() will process the request before login_required().

## What is reversing 

In Django, you can use the reverse() function to generate URLs from view names. This function allows you to look up the URL by supplying the name of the view, essentially going from "view name" to "URL". Additionally, reverse_lazy() is useful in scenarios where you want to delay the URL resolution until it is needed, such as in class-based views. Remember to double-check your view names and URLs to ensure they match correctly.

## reverse() and reverse_lazy()

The reverse() and reverse_lazy() functions are both used for URL reversing in Django, but they have key differences. reverse() returns a string representing the URL immediately, making it suitable for use in functions. In contrast, reverse_lazy() returns a lazy object, which means it will evaluate to the URL string when needed, making it useful when the URLConf is not available at the time of calling, such as in class-based views.

