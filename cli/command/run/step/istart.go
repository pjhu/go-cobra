package step

type IStartStepTemplateMethod interface {
	step1()
	step2()
	step3()
}

type IStartStep struct {
	Template        IStartStepTemplateMethod
}

func (o *IStartStep) Start() {
	o.Template.step1()
	o.Template.step2()
	o.Template.step3()
}