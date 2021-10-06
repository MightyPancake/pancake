package pancake

import(
  //Using SDL2 Go bindgings
  "github.com/veandco/go-sdl2/sdl"
  //"github.com/veandco/go-sdl2/img"
  "github.com/veandco/go-sdl2/mix"
  //Go Lua interpreter
  "github.com/Shopify/go-lua"
  //Pancake GL
  "pancake/pgl"
  //And some other, basic stuff
  "fmt"
  "runtime"
  "pancake/keycodes"
  "math"
  "io/ioutil"
  "os"
  "strconv"
  //"pgl"
)

//Constants
var winWidth, winHeight int
//Variables
var Window *sdl.Window
var Renderer *sdl.Renderer

//Loop
var Running bool
var FPS float32
var SinceStart int
var StartUpTime int

var KeyState []uint8
var State state
var LuaState *lua.State

//Collections
var Canvases map[string]Canvas
var Sounds map[string]*mix.Chunk
var Music map[string]*mix.Music
var Images map[string]Image
var Shaders map[string]Shader

//Graphical variables
var CurrentCanvas string
var TranslateX int
var TranslateY int
var ScaleX float32
var ScaleY float32
var CurrentColor color

//PGL variables
var DefaultShader Shader

//Declaring event function types, setting default functions and putting them in pancake.Event.EventName fashion...
type StartFunc func()// <--- Thos happens after pancake have been started!
type UpdateFunc func(float32)// <----- dt (Time)
type DrawFunc func()// <--- Nothing needs to be passed, yey!
type MousePressed func(int,int,string, int)// <---- X, Y, mouse button, clicks (single click, double click etc.)
type MouseReleased func(int,int,string, int)// <---- X, Y, mouse button, clicks (single click, double click etc.)
type KeyPressed func(string)// <---- Key
type KeyReleased func(string)// <---- Key
func emptyStart(){}
func emptyUpdate(dt float32){}
func emptyDraw(){}
func emptyMouse(x,y int, button string, clicks int){}
func emptyKeyFunc(key string){}

type events struct{
  Start StartFunc
  Update UpdateFunc
  Draw DrawFunc
  MousePressed MousePressed
  MouseReleased MouseReleased
  KeyPressed KeyPressed
  KeyReleased KeyReleased
}

var Event events

func init(){
  fmt.Println("\nPancake is running!\n")
  Event.Start = emptyStart
  Event.Update = emptyUpdate
  Event.Draw = emptyDraw
  Event.MousePressed = emptyMouse
  Event.MouseReleased = emptyMouse
  Event.KeyPressed = emptyKeyFunc
  Event.KeyReleased = emptyKeyFunc
  if error := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048); error != nil {
    fmt.Println(error)
  }
  Sounds = make(map[string]*mix.Chunk)
  Music = make(map[string]*mix.Music)
  Images = make(map[string]Image)
  Canvases = make(map[string]Canvas)
  Shaders = make(map[string]Shader)
}

func fixPancakePath(p string) string {
  newS := ""
  for _, rune := range p {
          //fmt.Printf("%d: %c\n", i, rune)
          if string(rune) == "\\" {
            newS += "/"
          }else{
            newS += string(rune)
          }

  }
  fmt.Println(newS)
  return newS
}

func addEnvVar(name, value string){
  path := os.Getenv(name)
  os.Setenv(name,path + ";" + value)
  fmt.Println("Added " + value + " to " + name + ".")
}

func clear(s *sdl.Surface){
  s.FillRect(nil, uint32(0))
}

type color struct{
  r,g,b,a uint8
}

func Start(w,h int, name string)(){
  err := sdl.Init(sdl.INIT_EVERYTHING)
  if err != nil {
    panic(err)
  }
  defer sdl.Quit()

  //Setting up variables
  winWidth = w
  winHeight = h
  name = StringIf(name=="","Pancake window", name)
  //Creating Window
  Window, err = sdl.CreateWindow(name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED ,int32(winWidth), int32(winHeight), sdl.WINDOW_OPENGL)
  if err != nil {
		return
	}
  defer Window.Destroy()
  //OpenGL related stuff
  SDL_INIT_GL(Window,3,3)
  pgl.Init(w,h)
  sh, err := pgl.LoadDefaultShaderProgram()
  if err != nil {
    panic(err)
  }
  DefaultShader = Shader(*sh)
  if err != nil {
    panic(err)
  }
  Origin()
  //Getting the window canvas
  if err != nil {
		return
	}
  //Getting renderer
  Renderer, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_TARGETTEXTURE)
  Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
  //Getting key state
  KeyState  = sdl.GetKeyboardState()
  //Creating entry state
  State = NewState(w,h, name)
  CurrentCanvas = name
  CurrentColor = color{0,0,0,255}
  //LOAD LUA FILES
  loadLua()
  //Here are some vars to keep frames capped
  FPS_Cap := 60.0
  var start uint64
  var end uint64
  var elapsedMS float32
  UniformFloat("screenWidth", float32(pgl.MaxW))
  UniformFloat("screenHeight", float32(pgl.MaxH))
  UniformFloat("Corners", 0.0)
  UniformFloat("Time", 0.0)
  //Trigger GO Start event
  Event.Start()
  Origin()
  //Trigger Lua Start event
  LuaStart()
  Origin()

  defer Window.Destroy()
  StartUpTime = int(sdl.GetTicks())
  Running = true
  for Running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
      handleEvent(event)
		}

    //Get frame start
    start = sdl.GetPerformanceCounter()
    //Update stuff
    update(float32(1.0/FPS_Cap))
    SinceStart = int(sdl.GetTicks()) - StartUpTime
    //Reset graphic state: color, translation, scale and shader!
    Origin()
    SetColor(0,0,0,255)
    ClearCanvas()
    //Actually draw stuff
    draw()
    //pgl.DrawQuad(0,0,300,300)
    Window.GLSwap()
    //Check if current shader has been edited
    CheckShaderUpdate()
    //Window.UpdateSurface() <--- Deprcated due to use of OpenGL
    //Renderer.Present()<-- Deprcated due to use of OpenGL
    //Calc time for sdl to wait...
    end = sdl.GetPerformanceCounter()
    elapsedMS = float32(end - start)/float32(sdl.GetPerformanceFrequency())*1000.0
    FPS = (1000/elapsedMS)
    if FPS > float32(FPS_Cap){
      FPS = float32(FPS_Cap)
    }
    //Wait the remaining time
    if elapsedMS < 16.666{
      sdl.Delay(uint32(16.666 - elapsedMS))
    }
	}
	return
}

func DrawRoundedRectangle(x,y,w,h,ox,oy int, rotation, corners float32){
  pgl.UnbindTexture()
  var fx float32 = (float32(x+ox)*2.0/float32(pgl.MaxW)) - 1
  var fy float32 = 1 - (float32(y+oy)*2.0/float32(pgl.MaxH))
  pgl.UniformVec4("origin", fx, fy, 0.0, 0.0)
  pgl.UniformFloat("Corners", corners)
  pgl.SetRotation(DegreesToRadians(rotation))
  pgl.DrawQuad(x,y,w,h)
}

func drawRoundedRectangle(state *lua.State) int {
  x, ok := state.ToNumber(1)
  y, ok := state.ToNumber(2)
  w, ok := state.ToNumber(3)
  h, ok := state.ToNumber(4)
  ox, ok := state.ToNumber(5)
  oy, ok := state.ToNumber(6)
  rot, ok := state.ToNumber(7)
  cor, ok := state.ToNumber(8)
  if !ok{
    fmt.Println("Error whle drawing a rounded rectangle!")
  }
  DrawRoundedRectangle(int(x),int(y),int(w),int(h),int(ox),int(oy), float32(rot), float32(cor))
  return 8
}


//Time
func TimeSinceStart() int{
  return SinceStart
}

//Pancake GL macros
type Shader pgl.Shader
type Texture pgl.TextureID
type Canvas *pgl.Canvas

type Image struct {
  Texture Texture
  Width int
  Height int
}

func SDL_INIT_GL(window *sdl.Window, major, minor int){
  sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
  sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, major)
  sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, minor)
  window.GLCreateContext()
}

func DrawRoundedQuad(x,y,w,h,ox,oy int, rotation, corners float32){
  var fx float32 = (float32(x+ox)*2.0/float32(pgl.MaxW)) - 1
  var fy float32 = 1 - (float32(y+oy)*2.0/float32(pgl.MaxH))
  pgl.UniformVec4("origin", fx, fy, 0.0, 0.0)
  pgl.UniformFloat("Corners", corners)
  pgl.SetRotation(DegreesToRadians(rotation))
  pgl.DrawQuad(x,y,w,h)
}

func LoadTexture(path string) (Texture, int, int) {
  t, w, h := pgl.LoadTextureAlpha(path)
  return Texture(t), w, h
}

func BindTexture(t Texture){
  pgl.BindTexture(pgl.TextureID(t))
}

func UnbindTexture(){
  pgl.UnbindTexture()
}

func CheckShaderUpdate(){
  pgl.UpdateShaderChanges()
}

func NewShader(vertexPath, fragmentPath string) (string, error) {
  shader, err := pgl.NewShader(vertexPath, fragmentPath)
  var name string = "Shader_" + strconv.Itoa(len(Shaders))
  sh := Shader(shader)
  Shaders[name] = sh
  return name, err
}

func newShader(l *lua.State) int {
  vertex, ok := l.ToString(1)
  fragment, ok := l.ToString(2)
  if !ok{
    fmt.Println("Error while compiling shader!")
  }
  name, err := NewShader(vertex, fragment)
  if err != nil {
    panic(err)
  }
  l.PushString(name)
  return 1
}

func UseShader(name string){
  var sh pgl.Shader
  if name == ""{
    sh = pgl.Shader(DefaultShader)
  }else{
    sh = pgl.Shader(Shaders[name])
  }
  pgl.UseShader(&sh)
  SendBaseUniforms()
}

func useShader(state *lua.State) int {
  shader, ok := state.ToString(1)
  if !ok{
    fmt.Println("Error while using shader!")
  }
  UseShader(shader)
  return 0
}

func SendBaseUniforms(){
  UniformFloat("screenWidth", float32(pgl.MaxW))
  UniformFloat("screenHeight", float32(pgl.MaxH))
  UniformFloat("Corners", 0.0)
  UniformFloat("Time", float32(SinceStart))
  pgl.SetColor(float32(CurrentColor.r)/255, float32(CurrentColor.g)/255, float32(CurrentColor.b)/255, float32(CurrentColor.a)/255)
}

func GL_Clear(r,g,b,a float32){
  pgl.Clear(r,g,b,a)
}

func UniformVec4(name string, x,y,z,w float32){
  pgl.UniformVec4(name,x,y,z,w)
}

func UniformFloat(name string, x float32){
  pgl.UniformFloat(name,x)
}

func SetRotation(r float32){
  pgl.SetRotation(r)
}

func DrawQuad(x,y,w,h,ox,oy int, rotation, corners float32){
  var fx float32 = (float32(x+ox)*2.0/float32(pgl.MaxW)) - 1
  var fy float32 = 1 - (float32(y+oy)*2.0/float32(pgl.MaxH))
  pgl.UniformVec4("origin", fx, fy, 0.0, 0.0)
  pgl.UniformFloat("Corners", corners)
  pgl.SetRotation(DegreesToRadians(rotation))
  pgl.DrawQuad(x,y,w,h)
}

func drawImage(state *lua.State) int {
  imgName, ok := state.ToString(1)
  x, ok := state.ToNumber(2)
  y, ok := state.ToNumber(3)
  rotation, ok := state.ToNumber(4)
  sx, ok := state.ToNumber(5)
  sy, ok := state.ToNumber(6)
  if !ok{
    fmt.Println("Error whle drawing an image!")
  }
  DrawImage(imgName, int(x), int(y), float32(rotation), float32(sx), float32(sy))
  return 6
}

func DrawImage(name string, x, y int, rotation float32, sx, sy float32){
  if _, ok := Images[name]; ok {
    img := Images[name]
    w := img.Width
    h := img.Height
    fx := float32(x)*ScaleX + float32(TranslateX)
    fy := float32(y)*ScaleY + float32(TranslateY)
    //origSX := sx*ScaleX
    //rigSY := sy*ScaleY
    finalSX := float32(sx*ScaleX)
    finalSY := float32(sy*ScaleY)
    finalW := int(finalSX*float32(w))
    finalH := int(finalSY*float32(h))
    t := img.Texture
    pgl.BindTexture(pgl.TextureID(t))
    DrawQuad(int(fx),int(fy),finalW,finalH,0,0, rotation, 0.0)
    //fmt.Println("Drawing image", name, "...")
  }else{
    panic("Image " + name +" does not exist!")
  }

}

/*func DrawImage(name string, x, y float64, rotation float32, sx, sy float32){
  fx := float32(x)*ScaleX + float32(TranslateX)
  fy := float32(y)*ScaleY + float32(TranslateY)
  tex := Images[name]
  _, _, width, height, err := tex.Query()
  if err != nil{
    fmt.Println("Error while drawing image ",name ,err)
  }
  finalSX := float32(math.Abs(float64(sx)))*ScaleX
  finalSY := float32(math.Abs(float64(sy)))*ScaleY
  finalW := int32(finalSX*float32(width))
  finalH := int32(finalSY*float32(height))

  rotPoint := sdl.Point{X:finalW/2, Y:finalH/2}//Ignore this, smh
  //rotPoint = sdl.Point{X:0, Y:0}

  finalRotation := rotation
  var flip sdl.RendererFlip = sdl.FLIP_NONE
  if (sx<0 && sy<0){
    //finalRotation += 180
    flip = sdl.FLIP_HORIZONTAL | sdl.FLIP_VERTICAL
    fx -= float32(finalW)
    fy -= float32(finalH)
  }else if sx < 0 {
    flip = sdl.FLIP_HORIZONTAL
    fx -= float32(finalW)
  }else if sy < 0 {
    flip = sdl.FLIP_VERTICAL
    fy -= float32(finalH)
  }
  tex.SetAlphaMod(CurrentColor.a)
  tex.SetColorMod(CurrentColor.r, CurrentColor.g, CurrentColor.b)
  Renderer.CopyEx(tex, nil, &sdl.Rect{X:int32(fx),Y:int32(fy),W:int32(float32(width)*finalSX), H:int32(float32(height)*finalSY)}, float64(finalRotation), &rotPoint, flip)
  tex.SetAlphaMod(255)
  tex.SetColorMod(255,255,255)
}*/

func DrawRectangle(x,y,w,h float64){
  fx := float32(x)*ScaleX + float32(TranslateX)
  fy := float32(y)*ScaleY + float32(TranslateY)
  fw := float32(w)*ScaleX
  fh := float32(h)*ScaleY
  rect(int(fx),int(fy),int(fw),int(fh))
}

func rect(x,y,w,h int){
  //rect := sdl.Rect{X:int32(x),Y:int32(y),W:int32(w),H:int32(h)}
  //Renderer.FillRect(&rect)
  //Deprcated due to use of OpenGL
  UnbindTexture()
  DrawQuad(x,y,w,h,0,0, 0.0,0.0)
}

func setColor(l *lua.State) int {
  r, ok := l.ToNumber(1)
  g, ok := l.ToNumber(2)
  b, ok := l.ToNumber(3)
  a, ok := l.ToNumber(4)
  if !ok{
    fmt.Println("Error while setting color!")
  }
  SetColor(int(r), int(g), int(b), int(a))
  return 4
}

func SetColor(r,g,b,a int){
  nr := r
  ng := g
  nb := b
  na := a
  if r < 0{nr=0}
  if g < 0{ng=0}
  if b < 0{nb=0}
  if a < 0{na=0}
  if r > 255{nr=255}
  if g > 255{ng=255}
  if b > 255{nb=255}
  if a > 255{na=255}
  CurrentColor = color{uint8(nr),uint8(ng),uint8(nb),uint8(na)}
  pgl.SetColor(float32(r)/255, float32(g)/255, float32(b)/255, float32(a)/255)
  //Renderer.SetDrawColor(uint8(nr),uint8(ng),uint8(nb),uint8(na)) <--- Deprcated due to use of OpenGL
}

func getColor(l *lua.State) int {
  l.NewTable()// {}
  l.PushNumber(float64(CurrentColor.r))// {}, color.r
  l.SetField(-2, "r")// {r = color.r}
  l.PushNumber(float64(CurrentColor.g))// {r = color.r}, color.g
  l.SetField(-2, "g")// {r = color.r, g = color.g}
  l.PushNumber(float64(CurrentColor.b))// {r = color.r, g = color.g}, color.b
  l.SetField(-2, "b")// {r = color.r, g = color.g, b = color.b}
  l.PushNumber(float64(CurrentColor.a))// {r = color.r, g = color.g, b = color.b}, color.a
  l.SetField(-2, "a")// {r = color.r, g = color.g, b = color.b, a = color.a}
  return 1
}

func drawRectangle(l *lua.State) int {
  x, ok := l.ToNumber(1)
  y, ok := l.ToNumber(2)
  w, ok := l.ToNumber(3)
  h, ok := l.ToNumber(4)
  if !ok{
    fmt.Print("Error while drawing a rectangle!")
  }
  if CurrentColor.a > 0{
    DrawRectangle(x,y,w,h)
  }
  return 4
}

func handleEvent(e sdl.Event){
  p := "pressed"
  switch event := e.(type) {
  case *sdl.QuitEvent:
    Running = false
  case *sdl.MouseButtonEvent:
    if event.Type == sdl.MOUSEBUTTONUP {p = "released"}
    if p == "pressed"{
      mousePressed(int(event.X), int(event.Y), MouseButtonString(event.Button), int(event.Clicks))
    }else if p == "released"{
      mouseReleased(int(event.X), int(event.Y), MouseButtonString(event.Button), int(event.Clicks))
    }
  case *sdl.KeyboardEvent:
    if event.State == sdl.PRESSED {p = "pressed"}else if event.State == sdl.RELEASED {p = "released"}
    if int(event.Repeat) == 0{
      key := getKeyFromCode(int(event.Keysym.Scancode))
      if p == "pressed"{
        keyPressed(key)
      }else if p == "released"{
        keyReleased(key)
      }
    }
  }
}

func IsKeyDown(key string) bool {
  k := keycodes.Code[key]
  if KeyState[k] != 0{
    return true
  }else{
    return false
  }
}

func getKeyFromCode(code int) string {
  for key, keyCode := range keycodes.Code {
    if keyCode == code{
      return key
      break
    }
  }
  return ""
}

func Origin(){
  SetColor(255,255,255,255)
  TranslateX = 0
  TranslateY = 0
  ScaleX = 1
  ScaleY = 1
  SetCanvas("")
  UseShader("")
}

func origin (l *lua.State) int {
  Origin()
  return 0
}

type state struct{
  Name string
  Canvas Canvas
  Objects []*Object
}

func NewState(w,h int, name string)(state){
  s := state{}
  n := NewCanvas(name, w, h)
  s.Canvas = Canvases[n]
  s.Name = name
  return s
}

func AddImage(name, path string)(Image, int, int){
  if path != ""{
    path = path + "/"
  }
  path = path + name + ".png"
  tex, w, h := pgl.LoadTextureAlpha(path)
  newImg := Image{Texture: Texture(tex), Width: w, Height: h}
  /*
  surf, err := img.Load(path + name + ".png")
  newImg, err := Renderer.CreateTextureFromSurface(surf)
  newImg.SetBlendMode(sdl.BLENDMODE_BLEND)
  if err != nil {
    fmt.Println("Cannot add image", name + ".png!",err)
  }*/
  Images[name] = newImg
  return newImg, w, h
}

func addImage(l *lua.State) int {
  name, ok := l.ToString(1)
  path, ok := l.ToString(2)
  if ! ok{
    fmt.Println("Error while adding image!")
  }
  _, w, h := AddImage(name, path)
  l.PushString(name)
  l.PushNumber(float64(w))
  l.PushNumber(float64(h))
  return 3
}

func AddObject(o Object)(*Object){
  State.Objects = append(State.Objects,& o)
  return &o
}

type Object struct{
  Name string
  X, Y, Scale float32
  Image string
}

func ClearCanvas(){
  c := CurrentColor
  Clear(float32(c.r/255), float32(c.g/255), float32(c.b/255), float32(c.a/255))
  //Renderer.Clear() <--- Deprcated due to use of OpenGL
}

func Clear(r,g,b,a float32){
  pgl.Clear(r,g,b,a)
}

//Utils
func StringIf(statement bool, t,f string)(string){
  if statement{
    return t
  }else{
    return f
  }
}
func IntIf(statement bool, t,f int)(int){
  if statement{
    return t
  }else{
    return f
  }
}
func Float32If(statement bool, t,f float32)(float32){
  if statement{
    return t
  }else{
    return f
  }
}

//Lua Related

func loadLua(){
  LuaState = lua.NewState()
  lua.OpenLibraries(LuaState)

  if err:= lua.DoFile(LuaState, "pancake.lua"); err != nil{
    fmt.Println(err)
  }
   LuaState.Global("initPancake")
   err := LuaState.ProtectedCall(0, 0, 0)
  if err != nil {
    fmt.Println(err)
  }

  //Adding lua-go bindings...
  //Export lua functions
  //Change this, so that these funcs are inside pancake table!
  LuaState.NewTable()
  addLuaFunc("drawImage",drawImage)
  addLuaFunc("addImage",addImage)
  addLuaFunc("getWindowHeight", getWindowHeight)
  addLuaFunc("getWindowWidth", getWindowWidth)
  addLuaFunc("newCanvas", newCanvas)
  addLuaFunc("setCanvas", setCanvas)
  addLuaFunc("getCanvas", getCanvas)
  addLuaFunc("setColor", setColor)
  addLuaFunc("getColor", getColor)
  addLuaFunc("clearCanvas", clearCanvas)
  addLuaFunc("getOS", getOS)
  addLuaFunc("isKeyDown", isKeyDown)
  addLuaFunc("drawRectangle", drawRectangle)
  addLuaFunc("translate", translate)
  addLuaFunc("origin", origin)
  addLuaFunc("setScale", setScale)
  addLuaFunc("getScale", getScale)
  addLuaFunc("getScaleX", getScaleX)
  addLuaFunc("getScaleY", getScaleY)
  addLuaFunc("getDirectoryItems", getDirectoryItems)
  addLuaFunc("getPathType", getPathType)
  addLuaFunc("log", log)
  addLuaFunc("getMouseX", getMouseX)
  addLuaFunc("getMouseY", getMouseY)
  addLuaFunc("isMouseDown", isMouseDown)
  addLuaFunc("drawCircle", drawCircle)
  addLuaFunc("drawLine", drawLine)
  addLuaFunc("getFPS", getFPS)
  addLuaFunc("addSound", addSound)
  addLuaFunc("playSound", playSound)
  addLuaFunc("addMusic", addMusic)
  addLuaFunc("playMusic", playMusic)
  addLuaFunc("newShader", newShader)
  addLuaFunc("useShader", useShader)
  addLuaFunc("drawRoundedRectangle", drawRoundedRectangle)
  LuaState.SetGlobal("pancake_go")//stack is empty
  LuaState.Pop(1)
  LuaState.Global("prepPancake")
  err = LuaState.ProtectedCall(0, 0, 0)
   if err != nil {
     fmt.Println(err)
   }
  //Loading main.lua ...
  if err:= lua.DoFile(LuaState, "main.lua"); err != nil{
    fmt.Println(err)
  }

}

func addLuaFunc(name string, f lua.Function){
  LuaState.PushGoFunction(f)
	LuaState.SetField(-2, name)
}

func log(l *lua.State) int {
  m, ok := l.ToString(1)
  fmt.Println(m)
  if !ok{
    fmt.Println("Error while logging a message!")
  }
  return 1
}

func getFPS(l *lua.State) int {
  l.PushNumber(float64(FPS))
  return 1
}

func getDirectoryItems(l *lua.State) int {
  path, ok := l.ToString(1)
  if !ok{
    fmt.Println("Error while getting directory items!")
  }
  l.NewTable()//<--t
  files, err := ioutil.ReadDir(path)
  if err != nil {
      fmt.Println(err)
  }
  for i, f := range files {
    //This will iterate over all files!
    name := f.Name()
    l.PushString(name)//{...}, name
    l.RawSetInt(-2, i+1)//{..., i+1 = name}
  }//{..., i+1 = name}
  if !ok{
    fmt.Println("Error while getting directory items!")
  }
  return 1
}

func GetPathType(p string) string {
  path := p
  fileinfo, err := os.Stat(path)
  if os.IsNotExist(err) {
		return ""
	}else{
    if fileinfo.IsDir(){
      return "directory"
    }else{
      return "file"
    }
  }
}

func getPathType(l *lua.State) int {
  p, ok := l.ToString(1)
  path := p
  if !ok{
    fmt.Println("Error while getting path type!")
  }
  t := GetPathType(path)
  if (t == ""){
    l.PushNil()
  }else{
    l.PushString(t)
  }
  return 1
}

func clearCanvas(l *lua.State) int {
  SetCanvas(CurrentCanvas)
  ClearCanvas()
  return 0
}

func getWindowHeight(l *lua.State) int{
  l.PushNumber(float64(winHeight))
  return 1
}

func getWindowWidth(l *lua.State) int{
  l.PushNumber(float64(winWidth))
  return 1
}

func NewCanvas(name string, w,h int) string {
  /*
  tex, err := Renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_TARGET, int32(w), int32(h))
  tex.SetBlendMode(sdl.BLENDMODE_BLEND)
  if err != nil {
    fmt.Println(err)
  }
  */
  canvas := pgl.NewCanvas(w,h)
  Images[name] = Image{Texture: Texture(canvas.Texture), Width: w, Height: h}
  Canvases[name] = canvas
  pgl.SetCanvas(canvas)
  return name
}

func newCanvas(l *lua.State) int {
  w, ok := l.ToNumber(1)
  h, ok := l.ToNumber(2)
  name, ok := l.ToString(3)
  canvas := NewCanvas(name, int(w), int(h))
  if !ok{
    fmt.Println("Error while creating new canvas!")
  }
  l.PushString(canvas)
  return 1
}

func setCanvas(l *lua.State) int {
  name, ok := l.ToString(1)
  SetCanvas(name)
  if !ok{
    fmt.Println("Error while setting canvas!")
  }
  return 1
}

func getCanvas(l *lua.State) int {
  l.PushString(CurrentCanvas)
  return 1
}

func SetCanvas(c string){
  CurrentCanvas = c
  if c == ""{
    //Renderer.SetRenderTarget(nil)
    pgl.SetDefaultCanvas()
  }else if _, ok := Canvases[c]; ok {
    //Renderer.SetRenderTarget(Canvases[c])
    canvas := pgl.Canvas(*Canvases[c])
    pgl.SetCanvas(&canvas)
  }else{
    fmt.Println("Canvas", c, "does not exist!")
  }
}

func getOS(l *lua.State) int {
  l.PushString(runtime.GOOS)
  return 1
}

func translate(l *lua.State) int {
  x, ok := l.ToNumber(1)
  y, ok := l.ToNumber(2)
  if !ok{
    fmt.Println("Error while translating graphic system!")
  }
  TranslateX += int(x*float64(ScaleX))
  TranslateY += int(y*float64(ScaleY))
  return 2
}

func setScale(l *lua.State) int {
  x, ok := l.ToNumber(1)
  y, ok := l.ToNumber(2)
  ScaleX *= float32(x)
  ScaleY *= float32(y)
  if !ok {
    fmt.Println("Error while scaling the graphic system!")
  }
  return 2
}

func getScale(l *lua.State) int {
  l.PushNumber(float64(ScaleX))
  l.PushNumber(float64(ScaleY))
  return 2
}

func getScaleY(l *lua.State) int {
  l.PushNumber(float64(ScaleY))
  return 1
}

func getScaleX(l *lua.State) int {
  l.PushNumber(float64(ScaleX))
  return 1
}

func drawCircle(l *lua.State) int {
  x, ok := l.ToNumber(1)
  y, ok := l.ToNumber(2)
  r, ok := l.ToNumber(3)
  if !ok{
    fmt.Print("Error while drawing circle!")
  }
  DrawCircle(x,y,r)
  return 3
}

func DrawCircle(x,y,r float64){
  fx := x*float64(ScaleX) + float64(TranslateX)
  fy := y*float64(ScaleY) + float64(TranslateY)
  fr := r*float64(ScaleX)
  drwCircle(int(fx),int(fy),int(fr))
}

func drwCircle(x,y,r int){
  /*for cx := -r; cx < r; cx++ {
    for cy := -r; cy < r; cy++ {
      if cx*cx + cy*cy <= r*r {
        Renderer.DrawPoint(int32(x+cx), int32(y+cy))
      }
    }
  }
  Deprcated due to use of OpenGL
  */
  UnbindTexture()
  DrawQuad(x-r, y-r, r*2, r*2, 0,0, 0.0, 1.0)
}

func DrawLine(x1,y1,x2,y2 float64){
  fx1 := x1*float64(ScaleX) + float64(TranslateX)
  fy1 := y1*float64(ScaleY) + float64(TranslateY)
  fx2 := x2*float64(ScaleX) + float64(TranslateX)
  fy2 := y2*float64(ScaleY) + float64(TranslateY)
  Renderer.DrawLine(int32(fx1), int32(fy1), int32(fx2), int32(fy2))
}

func drawLine(l *lua.State) int {
  x1, ok := l.ToNumber(1)
  y1, ok := l.ToNumber(2)
  x2, ok := l.ToNumber(3)
  y2, ok := l.ToNumber(4)
  if !ok{
    fmt.Println("Error while drawing line!")
  }
  DrawLine(x1,y1,x2,y2)
  return 4
}

func isKeyDown(l *lua.State) int {
  key, ok := l.ToString(1)
  if !ok{
    fmt.Println("Error while checking if ", key, " is down")
  }
  l.PushBoolean(IsKeyDown(key))
  return 1
}

func GetMouseX() int {
  x,_,_ := sdl.GetMouseState()
  return int(x)
}

func GetMouseY() int {
  _,y,_ := sdl.GetMouseState()
  return int(y)
}

func getMouseX(l *lua.State) int {
  l.PushNumber(float64(GetMouseX()))
  return 1
}

func getMouseY(l *lua.State) int {
  l.PushNumber(float64(GetMouseY()))
  return 1
}

func IsMouseDown(button string) bool {
  ret := false
  _,_,s := sdl.GetMouseState()
  if button =="left"{
    ret = (s == sdl.ButtonLMask())
  }else if button =="right"{
    ret = (s == sdl.ButtonRMask())
  }else if button =="middle"{
    ret = (s == sdl.ButtonMMask())
  }
  return ret
}

func MouseButtonString(button uint8) string {
  ret := "unknown"
  switch button {
  case sdl.BUTTON_LEFT:
    ret = "left"
  case sdl.BUTTON_RIGHT:
    ret = "right"
  case sdl.BUTTON_MIDDLE:
    ret = "middle"
  case sdl.BUTTON_X1:
    ret = "x1"
  case sdl.BUTTON_X2:
    ret = "x2"
  }
  return ret
}

func isMouseDown(l *lua.State) int {
  button, ok := l.ToString(1)
  if !ok{
    fmt.Println("Error while getting if mouse is down!")
  }
  l.PushBoolean(IsMouseDown(button))
  return 1
}

//--------Audio----------
//Sounds
func AddSound(name, dir string) *mix.Chunk {
  path := dir
  if dir != ""{path = path + "/"}
  path = path + name + ".wav"
  chunk, err := mix.LoadWAV(path)
  if err != nil{
    fmt.Println(err)
  }
  Sounds[name] = chunk
  return chunk
}
func addSound(l *lua.State) int {
  name, ok := l.ToString(1)
  dir, ok := l.ToString(2)
  if !ok{
    fmt.Println("Lua:","Error while adding sound '", name, "'")
  }
  AddSound(name,dir)
  l.PushString(name)
  return 2
}

func PlaySound(name string, repetitions int) int {
  times := repetitions - 1
  ret := -1
  if chunk, ok := Sounds[name]; ok {
    channel, _ := chunk.Play(-1, times)
    ret = channel
  }else{
    fmt.Println("Sound '", name , "' doesn't exit!")
  }
  return ret
}
func playSound(l *lua.State) int {
  name, ok := l.ToString(1)
  reps, ok := l.ToNumber(2)
  if !ok{
    fmt.Println("Lua:","Error while playing sound'", name, "'")
  }
  channel := PlaySound(name, int(reps))
  l.PushNumber(float64(channel))
  return 2
}

//Music
func AddMusic(name, dir string) *mix.Music {
  path := dir
  if dir != ""{path = path + "/"}
  path = path + name + ".mp3"
  mus, err := mix.LoadMUS(path)
  if err != nil{
    fmt.Println(err)
  }
  Music[name] = mus
  return mus
}
func addMusic(l *lua.State) int {
  name, ok := l.ToString(1)
  dir, ok := l.ToString(2)
  if !ok{
    fmt.Println("Lua:","Error while adding music '", name, "'")
  }
  AddMusic(name,dir)
  l.PushString(name)
  return 2
}

func PlayMusic(name string, loops int) {
  if mus, ok := Music[name]; ok {
    err := mus.Play(loops - 1)
    if err != nil{
      fmt.Println(err)
    }
  }else{
    fmt.Println("Music '", name , "' doesn't exit!")
  }
}
func playMusic(l *lua.State) int {
  name, ok := l.ToString(1)
  reps, ok := l.ToNumber(2)
  if !ok{
    fmt.Println("Lua:","Error while playing '", name, "' music.")
  }
  PlayMusic(name, int(reps))
  return 2
}

//--------Start----------
func LuaStart(){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"start")
  err := LuaState.ProtectedCall(0, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}

//--------Update----------
func update(dt float32){
  Event.Update(dt)
  LuaUpdate(dt)
}
func LuaUpdate(dt float32){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"update")
  LuaState.PushNumber(float64(dt))
  err := LuaState.ProtectedCall(1, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}

//--------Draw----------
func draw(){
  Origin()
  Event.Draw()
  Origin()
  LuaDraw()
  Origin()
}
func LuaDraw(){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"draw")
  err := LuaState.ProtectedCall(0, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}

//--------Mouse----------
func mousePressed(x, y int, button string, clicks int){
  Event.MousePressed(x, y, button, clicks)
  LuaMousePressed(x, y, button, clicks)
}

func mouseReleased(x, y int, button string, clicks int){
  Event.MouseReleased(x, y, button, clicks)
  LuaMouseReleased(x, y, button, clicks)
}
func LuaMousePressed(x,y int, button string, clicks int){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"mousePressed")
  LuaState.PushNumber(float64(x))
  LuaState.PushNumber(float64(y))
  LuaState.PushString(button)
  LuaState.PushNumber(float64(clicks))
  err := LuaState.ProtectedCall(4, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}
func LuaMouseReleased(x,y int, button string, clicks int){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"mouseReleased")
  LuaState.PushNumber(float64(x))
  LuaState.PushNumber(float64(y))
  LuaState.PushString(button)
  LuaState.PushNumber(float64(clicks))
  err := LuaState.ProtectedCall(4, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}

//--------Key----------
func keyPressed(key string){
  Event.KeyPressed(key)
  LuaKeyPressed(key)
}

func keyReleased(key string){
  Event.KeyReleased(key)
  LuaKeyReleased(key)
}
func LuaKeyPressed(key string){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"keyPressed")
  LuaState.PushString(key)
  err := LuaState.ProtectedCall(1, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}
func LuaKeyReleased(key string){
  LuaState.Global("pancake")
  LuaState.Field(-1,"event")
  LuaState.Field(-1,"keyReleased")
  LuaState.PushString(key)
  err := LuaState.ProtectedCall(1, 0, 0)
  if err != nil{
    fmt.Println(err)
  }
}

//Utils
func DegreesToRadians(a float32) (float32) {
  return float32(a/180*math.Pi)
}

func Sigma(x int) int{
  var ret int = -1
  if x == 0 {
    ret = 0
  }else if x > 0 {
    ret = 1
  }
  return ret
}
