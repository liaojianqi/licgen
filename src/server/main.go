package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/smartwalle/pongo2gin"
    "utils"
    "encoding/json"
    "fmt"
	"os"
    "time"
)
type Pair struct {
    s string
    license string
}
func checkFileIsExist(filename string) (bool) {
    var exist = true;
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false;
    }
    return exist;
}
func do_generate (c *gin.Context) {
    client := c.PostForm("client")
    not_before := c.PostForm("not_before")
    not_after := c.DefaultPostForm("not_after", "anonymous")
    max_host := c.PostForm("max_host")
    l := utils.License{client, not_before, not_after, max_host}
    json_string, _ := json.Marshal(l)
    cc := utils.Encrypt(json_string)
    license := utils.Hex_byte_to_string(cc)
    //save to file
    filename := "license_log.txt"
    var f *os.File
    var e error
    if checkFileIsExist(filename) {
        f, e = os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0666)
        if e != nil {
            panic(e)
        }
    }else {
        f, e = os.Create(filename)
        if e != nil {
            panic(e)
        }
    }
    defer f.Close()
    t := time.Now()
    f.WriteString(client + " " + not_before + " " + not_after + " " + max_host + " " + t.Format("2006-01-02") + " " + license + "\n")
    //下载文件后重定向
    c.HTML(http.StatusOK, "generate.tmpl", map[string]interface{}{"License": Pair{t.Format("2006-01-02") + client, license}})
}
func show_license(c *gin.Context) {
    var license_lists []Pair
    file, err := os.Open("license_log.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    for {
        var v1, v2, v3, v4, v5, v6 string
        _, err := fmt.Fscanln(file, &v1, &v2, &v3, &v4, &v5, &v6)
        if err != nil {
            break
        }
        license_lists = append(license_lists, Pair{v5 + v1, v6})
    }
    fmt.Println("len", len(license_lists))
    c.HTML(http.StatusOK, "show_archive.tmpl", map[string]interface{}{"license_lists": license_lists})
}
func main(){
    router := gin.Default()
    router.HTMLRender = pongo2gin.NewGinRender("./templates")
    router.GET("/index", func(c *gin.Context) {
       c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{"title": "客户License信息"})
    })
    router.GET("/show_license", show_license)
    router.POST("/generate", do_generate)
    router.Run(":8888")
}