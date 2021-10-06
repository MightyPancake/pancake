package pgl

import(
  "github.com/go-gl/gl/v3.3-core/gl"
  //"github.com/veandco/go-sdl2/sdl"
  "strings"
  "fmt"
  "io/ioutil"
  "os"
  "image/png"
  "errors"
)

type ShaderID uint32
type ProgramID uint32
type BufferID uint32
type TextureID uint32

var winWidth int
var winHeight int

var Started bool = false
var MaxW int
var MaxH int

var WhiteTexture TextureID

func Init(w, h int){
  SetResolution(w,h)
  //Setting up version
  //Actually starting stuff
  gl.Init()
  fmt.Println("OpenGL Version", GetVersion())
  WhiteTexture = loadWhiteTex()
  UnbindTexture()
  //Setting up...
  GenBindBuffer(gl.ARRAY_BUFFER)
  VAO := GenBindVertexArray() //Returns Vertex array!
  //Add data to buffer
  GenBindBuffer(gl.ELEMENT_ARRAY_BUFFER)
  //Specify the start and endpoint of data, its type etc.
  gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
  gl.EnableVertexAttribArray(0)
  gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
  gl.EnableVertexAttribArray(1)
  UnbindVertexArray()
  //Blending
  gl.Enable(gl.BLEND)
  gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
  gl.BindVertexArray(uint32(VAO))

  Started = true
}

func UnbindTexture(){
  BindTexture(WhiteTexture)
}

func LoadTextureAlpha(filename string)(TextureID, int, int){
  infile, err := os.Open(filename)
  if err != nil{
    panic(err)
  }
  defer infile.Close()

  img, err := png.Decode(infile)
  if err != nil{
    panic(err)
  }

  w := img.Bounds().Max.X
  h := img.Bounds().Max.Y
  pixels := make([]byte, w*h*4)
  bIndex := 0
  for y := 0; y < h; y++ {
    for x := 0; x < w; x++ {
      r,g,b,a := img.At(x,y).RGBA()
      pixels[bIndex] = byte(r/256)
      bIndex++
      pixels[bIndex] = byte(g/256)
      bIndex++
      pixels[bIndex] = byte(b/256)
      bIndex++
      pixels[bIndex] = byte(a/256)
      bIndex++
    }
  }
  texture := GenBindTexture()
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
  //Nearest(?)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

  gl.GenerateMipmap(gl.TEXTURE_2D)

  return texture, w, h
}

func loadWhiteTex() (TextureID){
  w := 64
  h := 64
  pixels := make([]byte, w*h*4)
  for i := 0; i<w*h*4; i++ {
    pixels[i] = byte(255)
  }
  texture := GenBindTexture()
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
  //Nearest(?)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

  gl.GenerateMipmap(gl.TEXTURE_2D)

  return texture
}

func SetResolution(w,h int){
  winWidth = w
  winHeight = h
  MaxW = w
  MaxH = h
}

func GetVersion() string {
  return gl.GoStr(gl.GetString(gl.VERSION))
}

func GL_DrawQuad(x,y,w,h float32){
  vertices := []float32 {
    x+w, y, 0.0, 1.0, 1.0, //top right 0
    x+w, y-h, 0.0, 1.0, 0.0, //bottom right 1
    x, y-h, 0.0, 0.0, 0.0, //bottom left 2
    x, y, 0.0, 0.0, 1.0} //top left 1st RECTANGLE DONE 3
  indices := []uint32 {
    0,1,3,//triangle 1
    1,2,3}//triangle 2
  //Actually draw
  //Add data to buffer
  BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
  BufferDataInt(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

  gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, gl.PtrOffset(0))
}

func DrawQuad(x,y,w,h int){
  var fx float32 = ((float32(x*2)/float32(MaxW)) - 1)
  var fy float32 = (1 - (float32(y*2)/float32(MaxH)))
  var fw float32 = (float32(w)/float32(MaxW))*2
  var fh float32 = (float32(h)/float32(MaxH))*2
  if CanvasUsed{
    fy = -fy
    fh = -fh
  }
  GL_DrawQuad(fx, fy, fw, fh)
}

func GenBindTexture() TextureID {
  var texID uint32
  gl.GenTextures(1, &texID)
  gl.BindTexture(gl.TEXTURE_2D, texID)
  return TextureID(texID)
}

func BindTexture(id TextureID) {
  gl.BindTexture(gl.TEXTURE_2D, uint32(id))
}

func LoadShader(path string, shaderType uint32) (ShaderID, error) {
  shaderFile, err := ioutil.ReadFile(path)
  if err != nil{
    panic(err)
  }
  return CreateShader(string(shaderFile), shaderType)
}

func CreateShader(shaderSource string, shaderType uint32) (ShaderID, error) {
  shaderID := gl.CreateShader(shaderType)
  shaderSource = shaderSource + "\x00"
  csource, free := gl.Strs(shaderSource)
  gl.ShaderSource(shaderID, 1, csource, nil)
  gl.CompileShader(shaderID)
  free()
  var status int32
  gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    var logLength int32
    gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)
    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(log))
    panic("Failed to compile shader: \n" + log)
    return 0, errors.New("Failed to compile shader: \n" + log)
  }
  return ShaderID(shaderID), nil
}

func NewProgram(vertPath, fragPath string) (ProgramID,error) {
  var err error
  var vert ShaderID
  var frag ShaderID
  if vertPath == ""{
    vert, err = CreateShader(DefaultVertexShader, gl.VERTEX_SHADER)
  }else{
    vert, err = LoadShader(vertPath, gl.VERTEX_SHADER)
  }
  if err != nil{
    return 0, err
  }
  if fragPath == ""{
    frag, err = CreateShader(DefaultFragmentShader, gl.FRAGMENT_SHADER)
  }else{
    frag, err = LoadShader(fragPath, gl.FRAGMENT_SHADER)
  }
  if err != nil{
    return 0, err
  }
  shaderProgram := gl.CreateProgram()
  gl.AttachShader(shaderProgram, uint32(vert))
  gl.AttachShader(shaderProgram, uint32(frag))
  gl.LinkProgram(shaderProgram)

  var success int32
  gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
  if success == gl.FALSE {
    var logLength int32
    gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
    return 0, errors.New("Failed to link program: \n" + log)
  }
  //Delete shaders, beause we don't need them anymore
  gl.DeleteShader(uint32(vert))
  gl.DeleteShader(uint32(frag))
  return ProgramID(shaderProgram), nil
}

func LoadDefaultShaderProgram() (*Shader, error) {
  vert, err := CreateShader(DefaultVertexShader, gl.VERTEX_SHADER)
  if err != nil{
    return &Shader{}, err
  }
  frag, err := CreateShader(DefaultFragmentShader, gl.FRAGMENT_SHADER)
  if err != nil{
    return &Shader{}, err
  }
  shaderProgram := gl.CreateProgram()
  gl.AttachShader(shaderProgram, uint32(vert))
  gl.AttachShader(shaderProgram, uint32(frag))
  gl.LinkProgram(shaderProgram)

  var success int32
  gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
  if success == gl.FALSE {
    var logLength int32
    gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
    return &Shader{}, errors.New("Failed to link program: \n" + log)
  }
  //Delete shaders, beause we don't need them anymore
  gl.DeleteShader(uint32(vert))
  gl.DeleteShader(uint32(frag))
  sh := Shader{ID:ProgramID(shaderProgram)}
  return &sh, nil
}

func GenBindBuffer(target uint32) BufferID {
  var buffer uint32
  gl.GenBuffers(1, &buffer)
  gl.BindBuffer(target, buffer)
  return BufferID(buffer)
}

func GenBindVertexArray() BufferID{
  var VAO uint32
  gl.GenVertexArrays(1, &VAO)
  gl.BindVertexArray(VAO)
  return BufferID(VAO)
}

func BindVertexArray(id BufferID){
  gl.BindVertexArray(uint32(id))
}

func BufferDataFloat(target uint32, data []float32, usage uint32){
  gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}

func BufferDataInt(target uint32, data []uint32, usage uint32){
  gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}

func UnbindVertexArray(){
  gl.BindVertexArray(0)
}

func UseProgram(id ProgramID){
  gl.UseProgram(uint32(id))
}

func SetColor(r,g,b,a float32){
  if Running(){
    UniformVec4("Color", r, g, b, a)
  }
}

func SetRotation(angle float32){
  UniformFloat("Rotation", angle)
}

func Clear(r,g,b,a float32){
  gl.ClearColor(r,g,b,a)
  gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Running() bool {
  return (CurrentShader != nil && Started)
}
