# pancake.isButtonClicked()

## Description

Checks whether the given [button](/documentation/topics/buttons) is clicked.

## Inputs

- [`button`](/documentation/topics/buttons) <- Table containg all data for [button](/documentation/topics/objects).

## Outputs

- `isClicked` <- Boolean that is `true` when [button](/documentation/topics/buttons) is pressed. Otherwise, it's `false`.

## Example

```lua
if pancake.isButtonClicked(right_button) then
  move()
end
```

The example will execute `move()` only when `righ_button` [button](/documentation/topics/buttons) is clicked.
