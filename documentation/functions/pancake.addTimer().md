# pancake.addTimer()

## Description

This function adds a [timer](/documentation/topics/timer) to `pancake.timers`!

## Inputs

- `time` <- Time attribute for this [timer](/documentation/topics/timer).
- `mode` <- This controls whether it should run once or be a loop: **"single"** or **"repetetive"**!
- `func` <- This is the function that will run when `time` hits 0.
- `arguments` <- These are the arguments that will be used to this function when called by this [timer](/documentation/topics/timer).

## Outputs

- [`timer`](/documentation/topics/timer) <- A pancake [timer](/documentation/topics/timer)!

## Example

```lua
  local bullet = pancake.applyPhysics(pancake.addObject({name = "bullet", x = 10, y = 10, width = 3, height = 1, image = "bullet", mass = 1}))
  bullet.velocityX = 30
  pancake.addTimer(2300, "single", delete, bullet)

--Somewhere else in the code
function delete(object)
  pancake.smartDelete(pancake.objects, object.ID, "ID")
end
```

The example above will add a bullet [object](/documentation/topics/objects) and then delete it after 2.3 seconds using a [timer](/documentation/topics/timer). Keep in mind, gravity will still work on this bullet. To defy it, read [this article](/documentation/topics/forces)!
