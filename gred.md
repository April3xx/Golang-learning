1) bool = true , false only != 0,1 and comparison is == no === cause golang has so strong types var declaration on golang can omit the type if assigned a value to it like : var booI,booJ,trK, = true,false,"STRING" --> is a type bool,bool,string accordingly
2) Tips: int is int no matter what size ... eg : srs int ='♥' fmt.Printf(%c,srs) --> ♥ default of bool,num,string =    false,0,""(empty string)
3) for i:=0 ; i<; i++ we call each a init statement;condition expression;post statement
4) init and post are optional using like when omit those--> // declare sum for ; sum<100 ;{//dosomething }
5) For is go while 
ุ6) forever loop is for without condition
7) if also has init use like --> if init; condition { //something here }
8) if a function is called and print in main the function will return results before fmt.print begins --> main { fmt.println(func()}
9) clean way to write multiple if else is switch with no condition it is switch true

