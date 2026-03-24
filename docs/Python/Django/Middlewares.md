# Middleware

[see built-in middlewares](https://docs.djangoproject.com/en/4.2/ref/middleware/)

A middleware factory is a callable that takes a get_response callable and returns a middleware. A middleware is a callable that takes a request and returns a response, just like a view.

```python
def simple_middleware(get_response):
    # One-time configuration and initialization.

    def middleware(request):
        # Code to be executed for each request before
        # the view (and later middleware) are called.

        response = get_response(request)

        # Code to be executed for each request/response after
        # the view is called.

        return response

    return middleware
```

Or with class based :

```python
class SimpleMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response
        # One-time configuration and initialization.

    def __call__(self, request):
        # Code to be executed for each request before
        # the view (and later middleware) are called.

        response = self.get_response(request)

        # Code to be executed for each request/response after
        # the view is called.

        return response
```

- Django initializes your middleware with only the get_response argument, so you can’t define __init__() as requiring any other arguments.

- Unlike the __call__() method which is called once per request, __init__() is called only once, when the web server starts.

## Activating middleware

To activate a middleware component, add it to the MIDDLEWARE list in your Django settings.

```python
MIDDLEWARE = [
    "django.middleware.security.SecurityMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]
```

The order in MIDDLEWARE matters because a middleware can depend on other middleware. For instance, AuthenticationMiddleware stores the authenticated user in the session; therefore, it must run after SessionMiddleware

You can think of it like an onion: each middleware class is a “layer” that wraps the view, which is in the core of the onion. If the request passes through all the layers of the onion (each one calls get_response to pass the request in to the next layer), all the way to the view at the core, the response will then pass through every layer (in reverse order) on the way back out.

If one of the layers decides to short-circuit and return a response without ever calling its get_response, none of the layers of the onion inside that layer (including the view) will see the request or the response. The response will only return through the same layers that the request passed in through.

## Middleware Hooks

we talked about __init__() and __call__()

### process_view()

process_view(request, view_func, view_args, view_kwargs)

request is an HttpRequest object. view_func is the Python function that Django is about to use. (It’s the actual function object, not the name of the function as a string.) view_args is a list of positional arguments that will be passed to the view, and view_kwargs is a dictionary of keyword arguments that will be passed to the view. Neither view_args nor view_kwargs include the first view argument (request).

**process_view() is called just before Django calls the view.**

It should return either None or an HttpResponse object. If it returns None, Django will continue processing this request, executing any other process_view() middleware and, then, the appropriate view. If it returns an HttpResponse object, Django won’t bother calling the appropriate view; it’ll apply response middleware to that HttpResponse and return the result.

### process_exception()

process_exception(request, exception)¶
request is an HttpRequest object. exception is an Exception object raised by the view function.

Django calls process_exception() when a view raises an exception. process_exception() should return either None or an HttpResponse object. If it returns an HttpResponse object, the template response and response middleware will be applied and the resulting response returned to the browser. Otherwise, default exception handling kicks in.

Again, middleware are run in reverse order during the response phase, which includes process_exception. If an exception middleware returns a response, the process_exception methods of the middleware classes above that middleware won’t be called at all.