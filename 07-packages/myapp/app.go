package main

import(
          "fmt"
          "https://github.com/shaukatmakandar786/GolangTutorials/tree/main/07-packages/number"
          "https://github.com/shaukatmakandar786/GolangTutorials/tree/main/07-packages/strings/greeting"
          "https://github.com/shaukatmakandar786/GolangTutorials/tree/main/07-packages/strings"
          //str "strings" //Package Alias
          
          
)

func main(){
 
  if number.IsPrime(19){
    
    fmt.Println("Given number is prime")
  }else{
   
    fmt.Println("give number is not a prime")
  }
  
  fmt.Println(greeting.WelcomeText)
  
  fmt.Println(strings.Reverse("ShaukatMakandar"))
    
  //fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))


  
}
