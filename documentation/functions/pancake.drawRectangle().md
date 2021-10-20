# pancake.drawRectangle()

## Description

Draws a rectangle with given attributes.

## Inputs

- `x` <- X coodrinate of the left side of the rectangle
- `y` <- Y coordinate of the upper side of the rectangle
- `w` <- Width of the rectangle
- `h` <- Height of the rectangle
- `ox (0)` <- X coordinate of the rotation origin point
- `oy (0)` <- Y coordinate of the rotation origin point
- `rotation (0)` <- The angle of rotation **in degrees**
- `cornerRounding (0)` <- Number that defines how much the corners should be rounded (from 0.0 to 1.0)

## Outputs

Nothing.

## Example

```lua
pancake.drawRectangle(0,0,200, 100)
```

Draws a 200 x 100 pixels rectangle.
