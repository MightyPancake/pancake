# pancake.addForce()

## Description

This function adds a [force](/documentation/topics/forces) to an [object](/documentation/topics/objects). This will result in applying it every frame! For one time applying, check [pancake.addForce()](/documentation/functions/pancake.addForce())!

## Inputs

* [`object`](/documentation/topics/objects) <- Which [object](/documentation/topics/objects) should have a force added to it.

* [`force`](/documentation/topics/forces) <- What force should be applied? **This should be a valid force.** You can [read more about forces here](/documentation/topics/forces).

## Outputs

Nothing.

## Example

```lua
pancake.addForce(bullet, {x = -pancake.physics.gravityX, y = -pancake.physics.gravityY, relativeToMass = true, time = "infinite"})
```

The example above will make bullet [object](/documentation/topics/objects) look like it defies gravity by adding force that is completely negating it! After this line, resultant force will be equal to 0 and bulet won't chnage it's velocity at all!
