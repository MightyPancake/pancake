# pancake.applyForce()

## Description

This function applies a [force](/documentation/topics/forces) to an [object](/documentation/topics/objects). Doing so will change the [object](/documentation/topics/objects)'s velocity.

?> **TIP:** The formula is velocity_change = (force_value/mass) × time.

## Inputs

* [`object`](/documentation/topics/objects) <- Which [object](/documentation/topics/objects) should have a force applied to it.
* [`force`](/documentation/topics/forces) <- What force should be applied? **This should be a valid force.** You can [read more about forces here](/documentation/topics/forces).
* `time`(pancake.lastdt) <- Time that the force is being applied
* `unsaved`(false) <- This defines if the [force](/documentation/topics/forces) should be added to resultant [force](/documentation/topics/forces) of the [object](/documentation/topics/object).

## Outputs

Nothing.

## Example

```lua
pancake.applyForce(player, {x = 0, y = -50, relativeToMass = true}, 1)
```

The code above will apply a [force](/documentation/topics/forces) that will make this [object](/documentation/topics/object) jump with fixed time and thus; height.
