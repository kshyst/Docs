# Serializers

Serializers allow complex data such as querysets and model instances to be converted to native Python datatypes that can then be easily rendered into JSON, XML or other content types. Serializers also provide deserialization, allowing parsed data to be converted back into complex types, after first validating the incoming data.

The serializers in REST framework work very similarly to Django's Form and ModelForm classes. We provide a Serializer class which gives you a powerful, generic way to control the output of your responses, as well as a ModelSerializer class which provides a useful shortcut for creating serializers that deal with model instances and querysets.

### Serializations

```python
serializer = SnippetSerializer(snippet)
serializer.data
# {'id': 2, 'title': '', 'code': 'print("hello, world")\n', 'linenos': False, 'language': 'python', 'style': 'friendly'}
```

At this point we've translated the model instance into Python native datatypes. To finalize the serialization process we render the data into json.

```python
content = JSONRenderer().render(serializer.data)
content
# b'{"id":2,"title":"","code":"print(\\"hello, world\\")\\n","linenos":false,"language":"python","style":"friendly"}'
```

### Deserialization

```python
import io

stream = io.BytesIO(content)
data = JSONParser().parse(stream)
```

In this code we parsed a stream into Python native datatypes. We can then restore these native datatypes into a fully populated object instance.

JsonRenderer creates bytes and JSONParser parses bytes into Python native datatypes.

```python
serializer = SnippetSerializer(data=data)
serializer.is_valid()
# True
serializer.validated_data
# {'title': '', 'code': 'print("hello, world")', 'linenos': False, 'language': 'python', 'style': 'friendly'}
serializer.save()
# <Snippet: Snippet object>
```

### **Use many=true to serialize multiple objects**




