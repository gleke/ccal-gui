// 由res2go自动生成，不要编辑。
package main

import (
	"github.com/ying32/govcl/vcl"
)

type TForm1 struct {
	*vcl.TForm
	Label1  *vcl.TLabel
	Label2  *vcl.TLabel
	Label3  *vcl.TLabel
	Label4  *vcl.TLabel
	Label5  *vcl.TLabel
	Label6  *vcl.TLabel
	Edit1   *vcl.TEdit   //年
	Edit2   *vcl.TEdit   //月
	Edit3   *vcl.TEdit   //日
	Edit4   *vcl.TEdit   //时辰
	Edit5   *vcl.TEdit   //生肖
	Edit6   *vcl.TEdit   //闰月
	Button1 *vcl.TButton //基础信息
	Button2 *vcl.TButton //２４节气
	Button3 *vcl.TButton //月历表
	Button4 *vcl.TButton //择吉
	Button5 *vcl.TButton //联系作者
	Button6 *vcl.TButton //地母经

	//::private::
	TForm1Fields
}

var Form1 *TForm1

// 以字节形式加载
// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1) {
	vcl.CreateResForm(owner, &root)
	return
}

var form1Bytes = []byte("\x54\x50\x46\x30\x06\x54\x46\x6F\x72\x6D\x31\x05\x46\x6F\x72\x6D\x31\x04\x4C\x65\x66\x74\x03\x5C\x01\x06\x48\x65\x69\x67\x68\x74\x03\x9B\x01\x03\x54\x6F\x70\x03\x75\x01\x05\x57\x69\x64\x74\x68\x03\x5D\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x8E\x86\xE6\xB3\x95\xE4\xBF\xA1\xE6\x81\xAF\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x9B\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x5D\x02\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x32\x2E\x30\x2E\x38\x2E\x30\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x2F\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x10\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE5\xB9\xB4\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x03\x83\x00\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x10\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE6\x9C\x88\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x03\xE0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x10\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE6\x97\xA5\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x03\x30\x01\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x10\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE6\x97\xB6\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x03\x90\x01\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x20\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\x94\x9F\xE8\x82\x96\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x36\x04\x4C\x65\x66\x74\x03\xE8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x20\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x97\xB0\xE6\x9C\x88\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x02\x18\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x32\x04\x4C\x65\x66\x74\x02\x70\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x0B\x45\x64\x69\x74\x32\x43\x68\x61\x6E\x67\x65\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x33\x04\x4C\x65\x66\x74\x03\xC8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x34\x04\x4C\x65\x66\x74\x03\x20\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x35\x04\x4C\x65\x66\x74\x03\x80\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x36\x04\x4C\x65\x66\x74\x03\xD8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x52\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xBA\xAA\xE5\xB9\xB4\xE4\xBF\xA1\xE6\x81\xAF\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\xE8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x5B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x32\x34\xE8\x8A\x82\xE6\xB0\x94\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x32\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x07\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x33\x04\x4C\x65\x66\x74\x03\xA0\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x5A\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE6\x9C\x88\xE5\x8E\x86\xE8\xA1\xA8\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x08\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x34\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x29\x03\x54\x6F\x70\x03\x10\x01\x05\x57\x69\x64\x74\x68\x02\x50\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x8B\xA9\xE5\x90\x89\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x09\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x35\x04\x4C\x65\x66\x74\x03\xA8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x22\x03\x54\x6F\x70\x03\x10\x01\x05\x57\x69\x64\x74\x68\x02\x52\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\x81\x94\xE7\xB3\xBB\xE4\xBD\x9C\xE8\x80\x85\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x0A\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x36\x04\x4C\x65\x66\x74\x03\xF2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x24\x03\x54\x6F\x70\x03\x10\x01\x05\x57\x69\x64\x74\x68\x02\x51\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE5\x9C\xB0\xE6\xAF\x8D\xE7\xBB\x8F\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x0B\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
