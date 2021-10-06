package pgl

var DefaultFragmentShader string = `
#version 330 core

out vec4 FragColor;

in vec2 TexCoord; //Coordinates of texture (in %, from 0.0 to 1.0)
in vec2 ScreenCoord; //Coordinates on screen (pixels)

//Pancake related uniforms
uniform sampler2D TextureUsed; //Texture used
uniform float screenWidth; //Width of the screen
uniform float screenHeight; //Hieght of the screen
uniform vec4 Color; //Color picked to draw with
uniform float Corners; //Defines how much the corners should be rounded

//My uniforms
uniform float Time;

vec4 TexAt(sampler2D t, vec2 t_coords){
  vec2 coords = vec2(t_coords.x, -t_coords.y);
  vec4 ret = texture(t, coords);
  return ret;
}

vec4 getColor(){
  vec4 ret = TexAt(TextureUsed, TexCoord);
  float x = abs(TexCoord.x*2.0-1.0);
  float y = abs(TexCoord.y*2.0-1.0);
  float r = Corners;
  float c = 1.0 - r;
  if (x > c && y > c){
    x = x-c;
    y = y-c;
    if (x*x+y*y>r*r){
      ret = vec4(0.0, 0.0, 0.0, 0.0);
    }
  }
  return ret;
}

void main(){
  vec4 TexColor = getColor();
  FragColor = TexColor*Color;
}

`
