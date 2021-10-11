# pancake.playSound()

## Description

Plays a [sound](/documentation/topics/sounds) that was previously added using [`pancake.addSound()`](/documentation/functions/pancake.addSound()).

## Inputs

- `name` <- Name of the [sound](/documentation/topics/sounds).
- `overlap`(false) <- If this is false and this [sound](/documentation/topics/sounds) already plays, it will stop and and play again from the start.

## Outputs

Nothing.

## Example

```lua
pancake.addSound("clap", "sounds")
pancake.playSound("clap")
```

The example above will add [sound](/documentation/topics/sounds) name "clap" from "sounds" folder of the main game directory and then play it once.
