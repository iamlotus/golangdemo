package main

type Handler interface {
	handle()
}

type ConcreateHandler int

func (c ConcreateHandler) handle() {

}

type HandlerInvoker interface {
	invoke(Handler)
}

type InvokerOnConcreateHandler int

func (self InvokerOnConcreateHandler) invoke(c ConcreateHandler) {
	c.handle()
}

func InvokeHandler(h HandlerInvoker) {}

func interfaceMain() {
	//InvokeHandler(InvokerOnConcreateHandler(1))
}
