# -* - coding: UTF-8 -* -
import ConfigParser
#生成config对象
conf = ConfigParser.ConfigParser()
#用config对象读取配置文件
conf.read("test.cfg")
#以列表形式返回所有的section
sections = conf.sections()
print 'sections:', sections         #sections: ['sec_b', 'sec_a']
#得到指定section的所有option
#options = conf.options("sec_a")
#print 'options:', options           #options: ['a_key1', 'a_key2']
