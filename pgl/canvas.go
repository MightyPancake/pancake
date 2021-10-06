package pgl

import(
  "github.com/go-gl/gl/v3.3-core/gl"
)

type FramebufferID uint32

var CanvasUsed bool = false

type Canvas struct {
  Framebuffer FramebufferID
  Texture TextureID
  Width int
  Height int
}

func SetCanvas(canvas *Canvas){
  BindFramebuffer(canvas.Framebuffer)
  w := canvas.Width
  h := canvas.Height
  gl.Viewport(0,0,int32(w),int32(h))
  MaxW = w
  MaxH = h
  CanvasUsed = true
}

func SetDefaultCanvas(){
  UnbindFramebuffer()
  w := winWidth
  h := winHeight
  gl.Viewport(0,0,int32(w),int32(h))
  MaxW = w
  MaxH = h
  CanvasUsed = false
}

func BindCanvas(canvas *Canvas){
  BindTexture(canvas.Texture)
}

func NewCanvas(w,h int) (*Canvas){
  frameBuffer := GenBindFramebuffer(gl.FRAMEBUFFER)
  tex := GenBindTexture()
  pixels := make([]byte, w*h*4)
  for i:=0;i<w*h*4;i++{
    pixels[i] = 0
  }
  //Do stuff to the texture
  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
  // Poor filtering. Needed !
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
  //Config framebuffer
  gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, uint32(tex), 0)
  //Set everything up in a canvas object
  canvas := Canvas{Framebuffer: frameBuffer, Texture: tex, Width:w, Height:h}
  buffer := uint32(frameBuffer)
  gl.DrawBuffers(1, &buffer)
  if (gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE){
    panic("Framebuffer error!")
  }
  UnbindFramebuffer()
  UnbindTexture()
  return &canvas
}

func GenBindFramebuffer(target uint32) FramebufferID {//Example: GL_FRAMEBUFFER
  var frameBuffer uint32
  gl.GenFramebuffers(1, &frameBuffer)
  gl.BindFramebuffer(target, frameBuffer)
  return FramebufferID(frameBuffer)
}

func BindFramebuffer(framebuffer FramebufferID){
  gl.BindFramebuffer(gl.FRAMEBUFFER, uint32(framebuffer))
}

func UnbindFramebuffer(){
  gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
