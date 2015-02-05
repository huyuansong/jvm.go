package class

//import cf "jvmgo/classfile"

type ClassMember struct {
    AccessFlags
    name        string
    descriptor  string
    class       *Class
}

func (self *ClassMember) Name() (string) {
    return self.name
}
func (self *ClassMember) Descriptor() (string) {
    return self.descriptor
}
func (self *ClassMember) Class() (*Class) {
    return self.class
}
