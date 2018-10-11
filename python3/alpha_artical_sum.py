#!/usr/bin/env python
# coding=utf-8
#统计一篇文章每个字母出现的次数?

a = """Please complete the following form and we will immediately email you download instructions and your personal download account information.

VanDyke Software products include strong cryptography. U.S. export regulations for encryption software require an eligibility screen before downloading is allowed. This software is subject to export control and may be transmitted, exported, or re-exported only under applicable export laws and restrictions and regulations of the United States Bureau of Industry and Security or foreign agencies or authorities.

Please respond to the following statement and fill in the information fields below.

* Fields marked with an asterisk must be completed in order to submit this request.

*    I acknowledge affirmatively that I understand that the requested software is subject to export controls under the Export Administration Act and that I may only export or re-export the software under the laws, restrictions and regulations of the U.S. Bureau of Industry and Security or foreign agencies or authorities."""

dic = dict()
for i in a:
    if i >= 'a' and i <= 'z' or i >= 'A' and i <= 'Z':
        if i not in dic.keys():
            dic[i] = 0
        dic[i] +=1
print(dic)
for i in dic.keys():
    print(dic[i])
