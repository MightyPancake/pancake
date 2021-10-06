package pgl

var DefaultVertexShader string = `
#version 330 core
#define M_PI 3.1415926535897932384626433832795

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

uniform float Rotation; //Rotation in radians
uniform vec4 origin; //Point of rotation
uniform float screenWidth; //Width of the screen
uniform float screenHeight; //Hieght of the screen
uniform sampler2D TextureUsed; //Texture used

uniform float Time;

out vec2 TexCoord; //Coordinates of texture (in %, from 0.0 to 1.0)
out vec2 ScreenCoord; //Coordinates on screen (pixels)

void main(){
  float XdevY = screenWidth/screenHeight;
  float angle = Rotation;
  mat4 rot = mat4(cos(angle), -sin(angle)/XdevY, 0.0, 0.0,
  sin(angle)*XdevY, cos(angle), 0.0, 0.0,
  0.0, 0.0, 0.0, 0.0,
  0.0, 0.0, 0.0, 1.0);
  vec4 pos = vec4(aPos, 1.0);
  mat4 offset = mat4(
  1.0, 0.0, 0.0, -origin.x,
  0.0, 1.0, 0.0, -origin.y,
  0.0, 0.0, 1.0, 0.0,
  0.0, 0.0, 0.0, 1.0);
  mat4 offsetBack = mat4(
  1.0, 0.0, 0.0, origin.x,
  0.0, 1.0, 0.0, origin.y,
  0.0, 0.0, 1.0, 0.0,
  0.0, 0.0, 0.0, 1.0);
  pos = pos*offset;
  pos = pos*rot;
  pos = pos*offsetBack;

  gl_Position = pos;

  TexCoord = aTexCoord;
  ScreenCoord = vec2((aPos.x+1.0)*screenWidth,(aPos.y+1.0)*screenHeight);
}

`
