package pgl

import(
  "os"
  "fmt"
  "time"
  "github.com/go-gl/gl/v3.3-core/gl"
)

var CurrentShader *Shader

type Shader struct {
  ID ProgramID
  vertexPath string
  fragmentPath string
  vertexModified time.Time
  fragmentModified time.Time
}
//var LoadedShaders = make(map[ProgramID]*Shader)

func NewShader(vertexPath, fragmentPath string) (Shader, error) {
  id, err := NewProgram(vertexPath, fragmentPath)
  if err != nil {
    return Shader{}, err
  }

  sh := Shader{ID:id, vertexPath:vertexPath, fragmentPath:fragmentPath}
  if vertexPath != ""{
    sh.vertexModified = getModifiedTime(vertexPath)
  }
  if fragmentPath != ""{
    sh.fragmentModified = getModifiedTime(fragmentPath)
  }
  //LoadedShaders[id] = &sh
  return sh, nil
}

func UseShader(shader *Shader){
  UseProgram(shader.ID)
  CurrentShader = shader
}

func UniformFloat(name string, f float32) {
  shader := CurrentShader
  name_cstr := gl.Str(name + "\x00")
  location := gl.GetUniformLocation(uint32(shader.ID), name_cstr)
  gl.Uniform1f(location, f)
}

func UniformVec4(name string, f1,f2,f3,f4 float32) {
  shader := CurrentShader
  name_cstr := gl.Str(name + "\x00")
  location := gl.GetUniformLocation(uint32(shader.ID), name_cstr)
  gl.Uniform4f(location, f1, f2, f3, f4)
}

func getModifiedTime(filePath string) time.Time {
  file, err := os.Stat(filePath)
  if err != nil{
    panic(err)
  }
  return file.ModTime()
}

func UpdateShaderChanges(){
  shader := CurrentShader
  if shader.fragmentPath != "" {
    fragmentModTime := getModifiedTime(shader.fragmentPath)
    if !fragmentModTime.Equal(shader.fragmentModified) {
      id, err := NewProgram(shader.vertexPath, shader.fragmentPath)
      if err != nil {
        fmt.Println(err)
      }else{
        gl.DeleteProgram(uint32(shader.ID))
        shader.ID = id
      }
    }
  }
  if shader.vertexPath != "" {
    vertexModTime := getModifiedTime(shader.vertexPath)
    if !vertexModTime.Equal(shader.vertexModified) {
      id, err := NewProgram(shader.vertexPath, shader.fragmentPath)
      if err != nil {
        fmt.Println(err)
      }else{
        gl.DeleteProgram(uint32(shader.ID))
        shader.ID = id
      }
    }
  }
}
