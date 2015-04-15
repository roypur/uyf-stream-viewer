package main

import "encoding/hex"
import "os"
import "strings"
import "os/exec"
var url string;

func main(){
    var hexString string = "";

    if len(os.Args) > 1 { 
        hexString = string(os.Args[1]);
    };  

    hexString = strings.Trim(hexString, "/");
    
	hexString = strings.TrimLeft(hexString, "uyf://");
	
    binUrl,_ := hex.DecodeString(hexString);
	
    url = string(binUrl);
	
    cmd := exec.Command("C:/apps/uyf/ffplay.exe",url);
	cmd.Start();
}
