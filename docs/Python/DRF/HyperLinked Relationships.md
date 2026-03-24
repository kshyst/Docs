# Hyperlinking our API

Dealing with relationships between entities is one of the more challenging aspects of Web API design. There are a number of different ways that we might choose to represent a relationship:

- Using primary keys.
- Using hyperlinking between entities.
- Using a unique identifying slug field on the related entity.
- Using the default string representation of the related entity.
- Nesting the related entity inside the parent representation.
- Some other custom representation.

REST framework supports all of these styles, and can apply them across forward or reverse relationships, or apply them across custom managers such as generic foreign keys.

In this case we'd like to use a hyperlinked style between entities. In order to do so, we'll modify our serializers to extend HyperlinkedModelSerializer instead of the existing ModelSerializer.

The **`HyperlinkedModelSerializer`** has the following differences from **`ModelSerializer`**:

- It does not include the id field by default.
- It includes a url field, using HyperlinkedIdentityField.
- Relationships use HyperlinkedRelatedField, instead of PrimaryKeyRelatedField.

