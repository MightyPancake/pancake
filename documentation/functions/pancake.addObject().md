# pancake.addObject()

## Description

Creates and saves an [object](/documentation/topics/objects) with given attributes with other [objects](/documentation/topics/objects).

## Inputs

* `object_data` <- Table containing all attributes that the [object](/documentation/topics/objects) should have

## Outputs

* [`object`](/documentation/topics/objects) <- Pancake [object](/documentation/topics/objects)

## Example

```lua
player = pancake.addObject({name = "Bob", x = 0, y = 0, width = 10, height = 10, colliding = true, image = "bob})
```

Creates an [object](/documentation/topics/objects) with attribute name set to "Bob", x to 0, etc.
