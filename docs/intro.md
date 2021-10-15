# Introduction


## Dual nature

Pancake consists of two seperate parts that work together:
#### Pancake.go engine
It is the main part of the engine responsible for outputing image, sound etc. and taking in stuff like keybaord and mouse presses.

It's written in Go and you can use it by just importing it to your Go project like so:
```Go
go get github.com/MightyPancake/pancake
```

#### pancake.lua library

This is essentially all of pancake's logic outside of inputing/outputing. It uses the engine and it won't run without it, **however it is possible to use it in [LÖVE 2D](https://love2d.org/) framework,** and you can learn more about it [here](/documentation/topics/settings)!


## Travel deep into the rabbit's hole

As you read articles here, some keywords are actually web links to an article that is related, so if you ever find yourself not knowing a related subject, do not despair, you can click this word to be redirected to the according article!

## Feedback

You've noticed something out of place? Or maybe you think something could have been written better? Feel free to e-mail me at **amightypancake@gmail.com** or propose a merge of the website [here](https://github.com/MightyPancake/pancake/tree/website)!
