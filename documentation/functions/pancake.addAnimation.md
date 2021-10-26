# pancake.addAnimation()

## Description

This function is used to create [animation](/documentation/topics/animations) data!

## Inputs

* `object_name` <- This is the name of the [object](/documentation/topics/objects) that this [animation](/documentation/topics/animations) will be used on. For more information head to [this article](/documentation/topics/animations?id=how-to-create-animation)
* `animation_name` <- Name of the action this [animation](/documentation/topics/animations) is suppose to represent!
* `folder` <- String containing path to the [animation](/documentation/topics/animations).
* `speed` (150) <- Time between frames in milliseconds.

## Outputs

* [`animation`](/documentation/topics/animations) <- Table containing all data like frames, time between them, object name etc.

## Example

```lua
  pancake.addAnimation("player", "run", "images/animations", 100)
```

Creates an [animation](/documentation/topics/animations) called "run" that you may use on all [objects](/documentation/topics/objects) named "player" and it will change its frame every 100 milliseconds.
