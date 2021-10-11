# Physics

## What are physics?

In pancake physics are set of rules that can be applied to certain [[object](/documentation/topics/objects)s](/documentation/topics/[object](/documentation/topics/objects)s) in order to simulate our world. They include:
- Velocity and acceleration (with the use of [forces](/documentation/topics/forces))
- Gravity
- Friction

!> **NOTE:** An [object](/documentation/topics/objects) has to have physics applied in order to move (excluding use manually changing value of coordinate)!

## General use

Generally, you just apply physics to an [object](/documentation/topics/objects) using [pancake.applyPhysics()](/documentation/functions/pancake.applyPhysics()). Then you can apply [forces](/documentation/topics/forces) to [objects](/documentation/topics/objects))!

## Physic attributes for objects

After applying physics to an [object](/documentation/topics/objects) you can access its phycics related attributes. These include:
- `velocityX`/`velocityY` <- Axis representing velocity of an [object](/documentation/topics/objects).
- `mass`(pancake.physics.defaultMass) <- The mass of the [object](/documentation/topics/objects).
- `friction` <- This parameter is responsible for how strongly friction should affect the [object](/documentation/topics/objects). Set this to 0 if the [object](/documentation/topics/objects) should be frictionless (like an ice, so [objects](/documentation/topics/objects)) can slide on it!).
- `maxVelocity`/`maxVelocityX`/`maxVelocityY`(pancake.physics.maxVelocity) <- This controls how fast the given [object](/documentation/topics/objects) can move.
- [forces](/documentation/topics/forces) <- Table containg all [forces](/documentation/topics/forces) that were added using [pancake.addForce()](/documentation/functions/pancake.addForce()) and are still being applied.
- [force](/documentation/topics/forces)
  * `x` <- `X` parameter of resultant [force](/documentation/topics/forces) that is working on [object](/documentation/topics/objects).
  * `y` <- `Y` parameter of resultant [force](/documentation/topics/forces) that is working on [object](/documentation/topics/objects).

!> **NOTE:** To get any of the above parameter, use [pancake.getStat()](/documentation/functions/pancake.getStat())!

## Pancake physics attributes

There are also few attributes that change overall rules of how pancake physics behave. All of them can be accessed in a table; `pancake.physics`:

- `defaultMass`(10) <- Default mass for [objects](/documentation/topics/objects).
- `gravityX`(0)/`gravityY`(12*pancake.meter) <- Axis of gravity vector.
- `defaultMaxVelocity`/`defaultMaxVelocityX`/`defaultMaxVelocityY`(15*pancake.meter for all 3 values) <- Defines default value for max velocity. You can set different values for each axis.
- `defaultFriction`(0.75) <- Defines what's the default friction parameter for [objects](/documentation/topics/objects)
