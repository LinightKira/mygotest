package pkges

//插件式开发

import (
	"github.com/dullgiulio/pingo"
	"log"
)

//Create an object be exported
type MyPlugin struct {
}

//Exported method with a RPC signature
func (p *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "Hello," + name
	return nil
}

//插件部分
func PingoRun() {
	plugin := &MyPlugin{}

	//Register the objects to be exported
	pingo.Register(plugin)
	//Run the events handler
	pingo.Run()
}

//插件主函数
func PingoMain(){
	//从创建的可执行文件中创建一个新插件，通过TCP连接到它
	p:=pingo.NewPlugin("tcp","helloworldpingo")
	//启动插件
	p.Start()
	//使用完插件后停止它
	defer p.Stop()
	var resp string
	//从先前创建的对象调用函数
	if err :=p.Call("MyPlugin.SayHello","GoGoGo developer",&resp);err !=nil{
		log.Print(err)
	}else {
		log.Println(resp)
	}
}