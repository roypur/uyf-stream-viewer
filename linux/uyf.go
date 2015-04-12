package main

import "os/exec"
import "encoding/hex"
import "os"
import "fmt"
import "strings"
import "io"

var url string;

func main(){
    var hexString string = "";

    if len(os.Args) > 1 { 
        hexString = string(os.Args[1]);
    };
    
    
    hexString = strings.TrimLeft(hexString, "uyf://");
    
    binUrl,_ := hex.DecodeString(hexString);
    
    url = string(binUrl);

    if(url == "install"){
        install();
    }else{
        cmd := exec.Command("/usr/bin/ffplay",url);
        cmd.Run();
    }    
    
}

func install(){

    var file string = "";
    
    cmd := exec.Command("/usr/bin/xdg-mime","--mode","system","default","uyf.desktop","x-scheme-handler/uyf");
    cmd.Run();
    
    f, err := os.Create("/usr/share/applications/uyf.desktop");

    if err != nil {
        fmt.Println(err);
    }
    
    file += "[Desktop Entry]\n";
    file += "Encoding=UTF-8\n";
    file += "Version=1.0\n";
    file += "Type=Application\n";
    file += "Terminal=false\n";
    file += "Exec=/usr/bin/uyf %u\n";
    file += "Name=UYF stream viewer\n";
    file += "Categories=Application;Network;\n";
    file += "MimeType=x-scheme-handler/uyf;\n";
    
    n, err := io.WriteString(f, file);

    if err != nil {
        fmt.Println(n, err);
    }

    f.Close();
}


