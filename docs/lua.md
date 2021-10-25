![lua_gif](https://www.lua.org/images/luaa.gif)

# Pancake.lua

Pancake.lua is a module that contains most logic based things in pancake. It is also responsible for binding output and input related functions so that they can be used in lua!

If you're new to Lua, please learn about it more [here](https://www.lua.org/)!

## Event callbacks

An event callback is a special type of function that runs whenever something happens. For example, `pancake.event.keypressed` will be called whenever user presses a button on his keyboard. Here's a list of all pancake callbacks:
  - pancake.event.draw() **NOT YET DOCUMENTED**
  - pancake.event.start() **NOT YET DOCUMENTED**
  - pancake.event.update() **NOT YET DOCUMENTED**
  - pancake.event.keypressed() **NOT YET DOCUMENTED**
  - pancake.event.keyreleased() **NOT YET DOCUMENTED**
  - pancake.event.mousepressed() **NOT YET DOCUMENTED**
  - pancake.event.mousereleased() **NOT YET DOCUMENTED**


## Lua callbacks
  - [pancake.init()](/documentation/functions/pancake.init()) **OUTDATED**
  - [pancake.draw()](/documentation/functions/pancake.draw()) **OUTDATED**

## Functions
 Below, you can find documentation for all functions used in pancake. Each includes a brief explenation, a list of inputs and outputs with descriptions and default values and a short example of how to use it:



### Animations
  - [pancake.addAnimation()](/documentation/functions/pancake.addAnimation()) **OUTDATED**
  - [pancake.changeAnimation()](/documentation/functions/pancake.changeAnimation()) **OUTDATED**

### Buttons
  - [pancake.addButton()](/documentation/functions/pancake.addButton()) **OUTDATED**
  - [pancake.isButtonClicked()](/documentation/functions/pancake.isButtonClicked()) **OUTDATED**

### Drawing
  - [pancake.drawImage()](/documentation/functions/pancake.drawImage())
  - [pancake.drawRectangle()](/documentation/functions/pancake.drawRectangle())
  - [pancake.drawCircle()](/documentation/functions/pancake.drawCircle())
  - [pancake.print()](/documentation/functions/pancake.print()) **OUTDATED**

### Images
  - [pancake.addImage()](/documentation/functions/pancake.addImage()) **OUTDATED**

### Files
  - [pancake.addAssets()](/documentation/functions/pancake.addAssets()) **OUTDATED**
  - [pancake.addFolder()](/documentation/functions/pancake.addFolder()) **OUTDATED**
  - [pancake.load()](/documentation/functions/pancake.load()) **OUTDATED**
  - [pancake.loadState()](/documentation/functions/pancake.loadState()) **OUTDATED**
  - pancake.getState() **NOT YET DOCUMENTED**
  - [pancake.save()](/documentation/functions/pancake.save()) **OUTDATED**
  - [pancake.saveState()](/documentation/functions/pancake.saveState()) **OUTDATED**

### Objects
**Basic**
  - [pancake.addObject()](/documentation/functions/pancake.addObject()) **OUTDATED**

**Coordinates**
  - [pancake.facing()](/documentation/functions/pancake.facing()) **OUTDATED**
  - [pancake.getFacingObjects()](/documentation/functions/pancake.getFacingObjects()) **OUTDATED**

**Physics**
  - [pancake.addForce()](/documentation/functions/pancake.addForce()) **OUTDATED**
  - [pancake.applyForce()](/documentation/functions/pancake.applyForce()) **OUTDATED**
  - [pancake.applyPhysics()](/documentation/functions/pancake.applyPhysics()) **OUTDATED**
  - [pancake.collisionCheck()](/documentation/functions/pancake.collisionCheck()) **OUTDATED**
  - [pancake.getSurfaceContact()](/documentation/functions/pancake.getSurfaceContact()) **OUTDATED**

### Sounds
  - [pancake.addSound()](/documentation/functions/pancake.addSound()) **OUTDATED**
  - [pancake.muteSounds()](/documentation/functions/pancake.playSound()) **OUTDATED**
  - [pancake.playSound()](/documentation/functions/pancake.playSound()) **OUTDATED**

### Timers
  - [pancake.addTimer()](/documentation/functions/pancake.addTimer()) **OUTDATED**

### Utility/Logic
  - [pancake.andCheck()](/documentation/functions/pancake.andCheck()) **OUTDATED**
  - [pancake.boolConversion()](/documentation/functions/pancake.boolConversion()) **OUTDATED**
  - [pancake.find()](/documentation/functions/pancake.find()) **OUTDATED**
  - [pancake.getDirectionName()](/documentation/functions/pancake.getDirectionName()) **OUTDATED**
  - [pancake.opposite()](/documentation/functions/pancake.opposite()) **OUTDATED**
  - [pancake.renderedObjects()](/documentation/functions/pancake.renderedObjects()) **OUTDATED**
  - [pancake.round()](/documentation/functions/pancake.round()) **OUTDATED**
  - [pancake.shakeScreen()](/documentation/functions/pancake.shakeScreen()) **OUTDATED**
  - [pancake.smartDelete()](/documentation/functions/pancake.smartDelete()) **OUTDATED**
  - [pancake.sumTables()](/documentation/functions/pancake.sumTables()) **OUTDATED**
  - [pancake.trash()](/documentation/functions/pancake.trash()) **OUTDATED**
