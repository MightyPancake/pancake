# pancake.addSound()

## Description

Function that adds given [sound](/documentation/topics/sounds) file to `pancake.sounds`, so it can be easily used many times!

!> **NOTE:** All [sounds](/documentation/topics/sounds) in pancake **HAVE** to be `.wav` files!

?> **TIP:** All [sounds](/documentation/topics/sounds) added that way are stored in `pancake.sounds[sound_name]`, where `sound_name` is the name that was called when adding it, thus, filename **without** `.wav` extension.

## Inputs

- `name` <- Name of the [sound](/documentation/topics/sounds). This is the name of the files, aswell as the name of where the [sound](/documentation/topics/sounds) will be stored in `pancake.sounds`. This doesn't include `.wav`!
- `path` <- This is the path to the [sound](/documentation/topics/sounds) **FOLDER**.

## Outputs

Nothing.

## Example

```lua
pancake.addSound("clap", "sounds")
pancake.playSound("clap")
```

The example above will add a [sound](/documentation/topics/sounds) named "clap" from "sounds" folder of the main game directory and then play it once.