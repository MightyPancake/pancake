# pancake.facing()

## Description

Function that returns if an [object](/documentation/topics/objects) is facing any **collidable** [object](/documentation/topics/objects) from a side.

!> **NOTE:** This will only work on **collidable** [objects](/documentation/topics/objects)! Keep that in mind!

## Inputs

* [`object`](/documentation/topics/objects) <- Which [object](/documentation/topics/objects) should be inspected.

## Outputs

* `directions` <- A table containing booleans for each side.
  - `left` <- Boolean that is true, if [object](/documentation/topics/objects) is facing something from left.
  - `right` <- Boolean that is true, if [object](/documentation/topics/objects) is facing something from right.
  - `up` <- Boolean that is true, if [object](/documentation/topics/objects) is facing something from left.up.
  - `down` <- Boolean that is true, if [object](/documentation/topics/objects) is facing something from down.

## Example

```lua
if pancake.facing(player).down then
  pancake.changeAnimation(player, "idle")
end
```

This will check if player [object](/documentation/topics/objects) is facing anything from below. If so, it changes its [animation](/documentation/topics/animations) to "idle".
