# pancake.changeAnimation()

## Description

Function used to apply previously added [animation](/documentation/topics/animations) to an [object](/documentation/topics/objects).

!> **NOTE:** As stated above, these [animations](/documentation/topics/animations) **have to be added before using this function**. To add them, simply use [pancake.addAnimation()](/documentation/functions/pancake.addAnimation())

## Inputs

* [`object`](/documentation/topics/objects) <- Target [object](/documentation/topics/objects). This object will have its [animation](/documentation/topics/animations) changed!

## Outputs

Nothing.

## Example

```lua
 pancake.changeAnimation(player, "run")
```

This will change the [animation](/documentation/topics/animations) of `player` [object](/documentation/topics/objects) to "run".
