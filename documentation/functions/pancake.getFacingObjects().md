# pancake.getFacingObjects()

## Description

Function that returns all **collidable** [objects](/documentation/topics/objects) that are next to target [object](/documentation/topics/objects) given as argument.

!> **NOTE:** This will only work on **collidable** [objects](/documentation/topics/objects)! Keep that in mind!

## Inputs

* [`object`](/documentation/topics/objects) <- Which [object](/documentation/topics/objects) should be inspected.

## Outputs

* `directions` <- A table containing all **collidable** [objects](/documentation/topics/objects) from left, right, up and down.
  - `left` <- Table containing all **collidable** [objects](/documentation/topics/objects) that have contact with target [object](/documentation/topics/objects) from its left side.
  - `right` <- Table containing all **collidable** [objects](/documentation/topics/objects) that have contact with target [object](/documentation/topics/objects) from its right side.
  - `up` <- Table containing all **collidable** [objects](/documentation/topics/objects) that have contact with target [object](/documentation/topics/objects) from its upper side.
  - `down` <- Table containing all **collidable** [objects](/documentation/topics/objects) that have contact with target [object](/documentation/topics/objects) from its down side.

## Example

```lua
local facingDown = pancake.getFacingObjects(player).down
if #facingDown > 0 then
  --Do things
  for i = 1, #facingDown do
    facingDown[i].name = "ground"
    facingDown[i].image = "ground"
  end
end
```

Above example will run when `player` [object](/documentation/topics/objects) is facing more then 0 [objects](/documentation/topics/objects) from its down edge. If so, all [object](/documentation/topics/objects) that it's facing from below will change its `name` and `image` to "ground". A more elegant way to check this condition is to use [pancake.facing()](/documentation/functions/pancake.facing()).
